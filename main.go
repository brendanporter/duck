package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"

	//"strconv"
	"encoding/json"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/brendanporter/quack"
	"golang.org/x/net/ipv4"
)

const CLR_0 = "\x1b[30;1m"
const CLR_R = "\x1b[31;1m"
const CLR_G = "\x1b[32;1m"
const CLR_Y = "\x1b[33;1m"
const CLR_B = "\x1b[34;1m"
const CLR_M = "\x1b[35;1m"
const CLR_C = "\x1b[36;1m"
const CLR_W = "\x1b[37;1m"
const CLR_N = "\x1b[0m"

var elog *log.Logger

var pingResults map[string][]quack.PingResult

var pingResultChan chan quack.PingResult
var resultRequestChan chan int
var resultResponseChan chan []byte

var ttlTraceResultChan chan []quack.PingResult
var unhealthyPingResultChan chan quack.PingResult

var pathResultsChan chan *PathStats
var pathResultsRequestChan chan int
var pathResultsResponseChan chan map[string]*PathStats

var hostResultsChan chan *quack.PingResult
var hostResultsRequestChan chan int
var hostResultsResponseChan chan map[string]*HostStats

func saveProcessor() {

	saveTicker := time.NewTicker(time.Minute)

	for {
		select {
		case <-saveTicker.C:
			err := saveDataToDisk("pings", pingResults)
			if err != nil {
				log.Print(err)
			}

			pathResultsRequestChan <- 1
			prr := <-pathResultsResponseChan

			err = saveDataToDisk("paths", prr)
			if err != nil {
				log.Print(err)
			}

			pathsResult := PathsResult{
				Paths:       prr,
				MessageType: "paths",
			}

			pathsJSON, err := json.Marshal(pathsResult)
			if err != nil {
				elog.Print(err)
			}

			wsHub.broadcast <- string(pathsJSON)

			hostResultsRequestChan <- 1
			hrr := <-hostResultsResponseChan

			err = saveDataToDisk("hosts", hrr)
			if err != nil {
				log.Print(err)
			}

			hostsResult := HostsResult{
				Hosts:       hrr,
				MessageType: "hosts",
			}

			hostsJSON, err := json.Marshal(hostsResult)
			if err != nil {
				elog.Print(err)
			}

			wsHub.broadcast <- string(hostsJSON)

		}
	}
}

func resultProcessor() {

	paths := make(map[string]*PathStats)
	hosts := make(map[string]*HostStats)

	err := readDataFromDisk("pings", &pingResults)
	if err != nil {
		elog.Print(err)
	}

	err = readDataFromDisk("paths", &paths)
	if err != nil {
		elog.Print(err)
	}

	err = readDataFromDisk("hosts", &hosts)
	if err != nil {
		elog.Print(err)
	}

	go wsHub.run()

	http.HandleFunc("/history", historyHandler)
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", baseHandler)
	server := http.Server{
		Addr: ":14445",
	}

	go server.ListenAndServe()

	go saveProcessor()

	for {
		select {
		case pr := <-pingResultChan:

			switch pr.ICMPMessage.Type {
			case ipv4.ICMPTypeEchoReply:

			case ipv4.ICMPTypeDestinationUnreachable:
				fmt.Printf("Destination network unreachable")
			case ipv4.ICMPTypeTimeExceeded:
			}

			var color string = CLR_W
			if pr.Latency < 40.0 {
				color = CLR_G
			} else if pr.Latency > 65.0 && pr.Latency < 150.0 {
				color = CLR_Y
			} else if pr.Latency > 150.0 {
				color = CLR_R
			}

			barCount := int((pr.Latency / 2000) * 80)

			fmt.Printf("%d bytes from %s: icmp_seq=%d ttl=%d time=%s%.3f ms |%s|%s\n", pr.Size, pr.Peer, pr.Sequence, pr.TTL, color, pr.Latency, strings.Repeat("-", barCount), CLR_W)

			pingResults[pr.Target] = append(pingResults[pr.Target], pr)
			jsonBytes, err := json.Marshal(pr)
			if err != nil {
				elog.Print(err)
			} else {
				wsHub.broadcast <- string(jsonBytes)
			}

		case traceResult := <-ttlTraceResultChan:

			tr := TraceResult{
				Hops:        traceResult,
				MessageType: "hops",
			}
			jsonBytes, err := json.Marshal(tr)
			if err != nil {
				elog.Print(err)
			} else {
				wsHub.broadcast <- string(jsonBytes)
			}

		case unhealthyPingResult := <-unhealthyPingResultChan:

			if unhealthyPingResult.Time < 100000 {
				break
			}
			upr := UnhealthyPingResult{
				Target:      unhealthyPingResult.Target,
				Latency:     unhealthyPingResult.Latency,
				Time:        unhealthyPingResult.Time,
				MessageType: "unhealthyPingResult",
			}
			jsonBytes, err := json.Marshal(upr)
			if err != nil {
				elog.Print(err)
			} else {
				wsHub.broadcast <- string(jsonBytes)
			}

		case <-resultRequestChan:
			resultsJSON, err := json.Marshal(pingResults)
			if err != nil {
				elog.Print(err)
			}

			resultResponseChan <- resultsJSON

		case nps := <-pathResultsChan:

			pathName := nps.PathName

			if _, ok := paths[pathName]; !ok {
				paths[pathName] = nps
			} else {
				paths[pathName].AvgLatency = (paths[pathName].AvgLatency + nps.AvgLatency) / 2
				paths[pathName].MaxLatencyAvg = (paths[pathName].MaxLatencyAvg + nps.MaxLatencyAvg) / 2
				paths[pathName].TripCount++
			}

		case pr := <-hostResultsChan:

			hostName := pr.Target

			if _, ok := hosts[hostName]; !ok {

				hostDNSNames, err := net.LookupAddr(hostName)
				if err != nil {
					elog.Print(err)
				}

				var hostDNSName string

				if len(hostDNSNames) > 0 {
					hostDNSName = hostDNSNames[0]
				}

				hosts[hostName] = &HostStats{}
				if pr.Latency > 0 {
					hosts[hostName].MinLatency = pr.Latency
				} else {
					hosts[hostName].MinLatency = 9999
				}
				hosts[hostName].AvgLatency = pr.Latency
				hosts[hostName].HostName = hostName
				hosts[hostName].HostDNSName = hostDNSName

			} else {

				if hosts[hostName].MinLatency > pr.Latency {
					hosts[hostName].MinLatency = pr.Latency
				}

				if hosts[hostName].MaxLatency < pr.Latency {
					hosts[hostName].MaxLatency = pr.Latency
				}

				hosts[hostName].AvgLatency = (hosts[hostName].AvgLatency + pr.Latency) / 2

			}

			hosts[hostName].TripCount++

			if pr.Latency > 700 {
				hosts[hostName].HighLatency700++
			} else if pr.Latency > 400 {
				hosts[hostName].HighLatency400++
			} else if pr.Latency > 100 {
				hosts[hostName].HighLatency100++
			}

		case <-pathResultsRequestChan:

			prr := make(map[string]*PathStats)

			for k, v := range paths {
				prr[k] = v
			}

			pathResultsResponseChan <- paths

		case <-hostResultsRequestChan:

			hrr := make(map[string]*HostStats)

			for k, v := range hosts {
				hrr[k] = v
			}

			hostResultsResponseChan <- hosts

		}

	}
}

var thirtySampleRollingLatency []float64

func init() {
	pingResults = make(map[string][]quack.PingResult)
	pingResultChan = make(chan quack.PingResult, 10)
	resultResponseChan = make(chan []byte, 10)
	resultRequestChan = make(chan int, 10)
	ttlTraceResultChan = make(chan []quack.PingResult, 10)
	unhealthyPingResultChan = make(chan quack.PingResult, 10)

	pathResultsChan = make(chan *PathStats, 10)
	pathResultsRequestChan = make(chan int, 10)
	pathResultsResponseChan = make(chan map[string]*PathStats, 10)

	hostResultsChan = make(chan *quack.PingResult, 10)
	hostResultsRequestChan = make(chan int, 10)
	hostResultsResponseChan = make(chan map[string]*HostStats, 10)

	elog = log.New(os.Stdout, "Error: ", log.LstdFlags|log.Lshortfile)

}

func echoResults(target string, packetsTx, packetsRx int, minLatency, avgLatency, maxLatency, stdDevLatency float64) {

	fmt.Print("\n")
	log.Printf("--- %s ping statistics ---", target)
	fmt.Printf("%d packets transmitted, %d packets received, %.1f%% packet loss\n", packetsTx, packetsRx, (float64(packetsTx-packetsRx) / float64(packetsTx) * 100))
	fmt.Printf("round-trip min/avg/max/stddev = %.3f/%.3f/%.3f/%.3f ms\n", minLatency, avgLatency, maxLatency, stdDevLatency)
	fmt.Printf("View charted results at: http://localhost:14445\n\n")
}

var resultFileNameSuffix string

func main() {

	go resultProcessor()

	var maxLatency float64
	var minLatency float64 = 99999.9
	var avgLatency float64
	var stdDevLatency float64
	var packetsTx int
	var packetsRx int

	var target string

	//flag.StringVar(&target, "target", "8.8.8.8", "IPv4 address of target")

	//target := "8.8.8.8"

	re := regexp.MustCompile(`(!?\d{1,3})\.(!?\d{1,3})\.(!?\d{1,3})\.(!?\d{1,3})`)

	if len(os.Args) > 1 && re.Match([]byte(os.Args[1])) {
		target = os.Args[1]
	}

	flag.StringVar(&resultFileNameSuffix, "suffix", fmt.Sprintf("_%s_%s", target, time.Now().Format(time.RFC3339)), "filename suffix for storing result data (results[-<target>-<timestamp>].json)")

	flag.Parse()

	fmt.Printf("View charted results at: http://localhost:14445\n")

	pingTicker := time.NewTicker(time.Second)
	traceTicker := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-pingTicker.C:
			packetsTx++
			latency, err := quack.SendPing(target, packetsTx, 0, pingResultChan)
			if err != nil {
				log.Print(err)
				fmt.Printf("Request timeout for icmp_seq %d\n", packetsTx)
				latency = 2000.0
			} else {
				packetsRx++
			}

			if avgLatency == 0 {
				avgLatency = latency
			} else {
				avgLatency = (avgLatency + latency) / 2
			}

			stdDevLatency += math.Pow(latency-avgLatency, 2)
			stdDevLatency = math.Sqrt(stdDevLatency / 2)

			if latency < minLatency {
				minLatency = latency
			}

			if latency > maxLatency {
				maxLatency = latency
			}

			if packetsTx%30 == 0 {
				echoResults(target, packetsTx, packetsRx, minLatency, avgLatency, maxLatency, stdDevLatency)
			}

			//time.Sleep(time.Duration(time.Second.Nanoseconds() - int64(latency*1000000)))

			if len(thirtySampleRollingLatency) == 30 {
				thirtySampleRollingLatency = thirtySampleRollingLatency[1:]
			}
			thirtySampleRollingLatency = append(thirtySampleRollingLatency, latency)

		case <-traceTicker.C:
			go func() {
				traceResults, err := quack.TTLTrace(target)
				if err != nil {
					elog.Print(err)
				}

				ttlTraceResultChan <- traceResults

				var lastLatency float64
				var highLatencyHosts []quack.PingResult
				var latencyHighWaterMark float64

				for i, traceResult := range traceResults {

					log.Printf("TTL %d result: %#v", i+1, traceResult)

					if lastLatency == 0 {
						lastLatency = traceResult.Latency
						continue
					}

					if math.Abs(traceResult.Latency-lastLatency) > 100 && traceResult.Latency != 0 && traceResult.Latency > latencyHighWaterMark {
						if i > 0 && traceResults[i-1].Latency > 100 {
							highLatencyHosts = append(highLatencyHosts, traceResults[i-1])
						}
						highLatencyHosts = append(highLatencyHosts, traceResult)
					}
					lastLatency = traceResult.Latency

					if traceResult.Latency > latencyHighWaterMark {
						latencyHighWaterMark = traceResult.Latency
					}

				}

				// If latency was high, perform additional ttlTraces to collect more path observations

				if latencyHighWaterMark > 100 {
					log.Printf("High latency of %.1fms detected. Performing additional traces.", latencyHighWaterMark)
					for x := 0; x < 4; x++ {
						traceResults, err := quack.TTLTrace(target)
						if err != nil {
							elog.Print(err)
						}

						_ = traceResults

						time.Sleep(time.Millisecond * 200)
					}
				}

				for _, highLatencyHost := range highLatencyHosts {
					log.Printf("Found potentially unhealthy host in trace: %#v", highLatencyHost)

					unhealthyPingResultChan <- highLatencyHost
				}

			}()
		}
	}

}

func digestTraceResults(traceResults []quack.PingResult) {

	var pathName string
	var targetNames []string
	var maxPathLatency float64
	var avgPathLatency float64

	for _, traceresult := range traceResults {

		targetNames = append(targetNames, traceresult.Target)

		tr := traceresult

		hostResultsChan <- &tr

		if avgPathLatency == 0 {
			avgPathLatency = traceresult.Latency
		} else {
			avgPathLatency = (avgPathLatency + traceresult.Latency) / 2
		}
		if maxPathLatency < traceresult.Latency {
			maxPathLatency = traceresult.Latency
		}

	}

	pathName = strings.Join(targetNames, "-")

	newPathStats := &PathStats{
		PathName:      pathName,
		AvgLatency:    avgPathLatency,
		MaxLatencyAvg: maxPathLatency,
		TripCount:     1,
	}

	pathResultsChan <- newPathStats

}
