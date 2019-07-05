package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var dashboard string = fmt.Sprintf(`<!doctype html>
<html>
<head>
<script>%s</script>
<style>%s</style>
<script>%s</script>

<style>%s</style>
<script>%s</script>
<script>%s</script>
<script>%s</script>

<title>Duck Latency Tester</title>
<script>
$(document).ready(function(){
	console.log("ready");
	init();
});


function roundToThree(num) {    
    return Math.round(num * 1000) / 1000;
}

var pingChart;

function init(){

	now = new Date;
	tzo = now.getTimezoneOffset() / 60;

	pingChart = Highcharts.chart('pingChart', {
	time: {
		timezoneOffset: 6 * 60,
	},
    chart: {
    	height: 300,
    },
    title: {
        text: 'Ping Latency Live'
    },
    subtitle: {
        text: '10-minute rolling window'
    },
    xAxis: {
        type: 'datetime'
    },
    yAxis: {
        title: {
            text: 'Latency (ms)'
        },
        plotBands: [
        	{
        		className: 'band-green',
        		color: '#114411',
        		from: 0,
        		to: 40
        	},
        	{
        		className: 'band-yellow',
        		color: '#A0792A',
        		from: 40,
        		to: 80,
        	},
        	{
        		className: 'band-red',
        		color: '#441111',
        		from: 80,
        		to: 999999,
        	},
        ],
        min: 0,
    },
    credits: {
		enabled: false
	},
    exporting: {
        enabled: false
    },
    tooltip: {
    	crosshairs: true,
    	valueSuffix: ' ms'
    },
    plotOptions: {
        series: {
            lineWidth: 1,
            animation: false,
        }
    },
    series: [
    	{name: 'Latency', type: 'line', lineWidth: 3},
    	{name: 'Alarming Latency', type: 'scatter', tooltip: {pointFormatter: function(i,e){
    		return  Highcharts.dateFormat('%%A, %%b %%e, %%H:%%M:%%S', new Date(this.x)) + '<br/><b>' + this.Target + ': ' + this.y + ' ms</b>';
    	}}}
    ],
});

pingHistoryChart = Highcharts.chart('pingHistoryChart', {
	time: {
		timezoneOffset: 6 * 60,
	},
    chart: {
        height: 200,
        zoomType: 'x',
    },
    title: {
        text: 'Ping Latency History'
    },
    subtitle: {
        text: '30-second averaged samples'
    },
    xAxis: {
        type: 'datetime'
    },
    yAxis: {
        title: {
            text: 'Latency (ms)'
        },
        plotBands: [
        	{
        		className: 'band-green',
        		color: '#114411',
        		from: 0,
        		to: 40
        	},
        	{
        		className: 'band-yellow',
        		color: '#A0792A',
        		from: 40,
        		to: 80,
        	},
        	{
        		className: 'band-red',
        		color: '#441111',
        		from: 80,
        		to: 999999,
        	},
        ],
        min: 0,
    },
    credits: {
		enabled: false
	},
    exporting: {
        enabled: false
    },
    tooltip: {
    	crosshairs: true,
    	shared: true,
    	valueSuffix: ' ms',
    },
    plotOptions: {
        series: {
            lineWidth: 1,
            animation: false,
        }
    },
    series: [
    	{name: 'Average', type: 'line', zIndex: 1, lineWidth: 3},
    	{name: 'Range', type: 'arearange', color: Highcharts.getOptions().colors[0], fillOpacity: 0.3, zIndex: 0, marker: {enabled: false}, lineWidth: 0, linkedTo: ':previous'},
    	{name: "Alarming Latency", type: 'scatter', tooltip: {pointFormatter: function(i,e){
    		return Highcharts.dateFormat('%%A, %%b %%e, %%H:%%M:%%S', new Date(this.x)) + '<b>' + this.Target + ': ' + this.y + ' ms</b>';
    	}}}],
});


	getResultHistory();

	connect();
}

function downRezChartData(chart, sampleSize) {
	console.log("Downrezing data")
	var newSeries = [];

	var newPoint = {};
	var newPointSamples = 0;
	var lastSampleXValue = chart.series[0].data[0].x;

	for(var i = 0; i < chart.series[0].data.length; i++){

		if(typeof chart.series[0].data[i].res != "undefined"){
			if(chart.series[0].data[i].res >= sampleSize){
				newSeries.push(chart.series[0].data[i]);
				continue;
			}
		}

		if(chart.series[0].data[i].x - lastSampleXValue > sampleSize * 1000){ // if samples are greater than sampleSize seconds apart, finish averaging batch, reset, skip
			lastSampleXValue = chart.series[0].data[i].x
			newPoint.x = newPoint.x / newPointSamples
			newPoint.y = newPoint.y / newPointSamples
			newSeries.push(newPoint);
			//chart.series[0].addPoint(newPoint, false)
			newPoint = {x: 0, y: 0, res: sampleSize};
			newPointSamples = 0;
			continue;
		}

		if(newPointSamples === sampleSize){
			// Average out the point data across sampleSize
			newPoint.x = newPoint.x / newPointSamples
			newPoint.y = newPoint.y / newPointSamples
			newSeries.push(newPoint);
			//chart.series[0].addPoint(newPoint, false)
			newPoint = {x: 0, y: 0, res: sampleSize};
			newPointSamples = 0;
		} else {
			//console.log("X:" + chart.series[0].data[i].x + " Y:" + chart.series[0].data[i].y)
			newPoint.x += chart.series[0].data[i].x
			newPoint.y += chart.series[0].data[i].y
			newPointSamples++;
			//chart.series[0].data[i].remove()
		}
	}

	chart.series[0].setData(newSeries);
	//chart.addSeries({data: newSeries});

	console.log("Finished downrezing data");
	//console.log(newSeries);
}

var webSocketFailed = false;
var webSocket;
var webSocketRetryTicker;

var chartDownRezCountdown = 30;
var historyX;
var avgdHistoryY;
var maxHistoryY;
var minHistoryY;

var tzo;

function getResultHistory(){

	var tenMinutesAgo = Date.now() - (tzo * 60 * 1000) - (10*60*1000)

	$.ajax({
		url: '/history',
		dataType: 'json',
		success: function(json){

			$.each(json['8.8.8.8'], function(i, jsonData){


				historyX += jsonData.Time * 1000;
			    avgdHistoryY += jsonData.Latency;

			    if(minHistoryY > jsonData.Latency){
			    	minHistoryY = jsonData.Latency;
			    }
			    if(maxHistoryY < jsonData.Latency){
			    	maxHistoryY = jsonData.Latency;
			    }


			    chartDownRezCountdown--;
			    if(chartDownRezCountdown === 0){
			    	console.log("Triggered downrez");
			    	chartDownRezCountdown = 30;
			    	pingHistoryChart.series[0].addPoint({x: historyX/chartDownRezCountdown, y: Math.round(avgdHistoryY/chartDownRezCountdown)}, false);
			    	pingHistoryChart.series[1].addPoint([historyX/chartDownRezCountdown, Math.round(minHistoryY), Math.round(maxHistoryY)], false);
			    	historyX = 0;
			    	maxHistoryY = 0;
			    	minHistoryY = 999999;
			    	avgdHistoryY = 0;
			    }

			    if(i < jsonData.length - 600){
			    	return;
			    } else {
			    	shift = false;
					if(pingChart.series[0].data[0] != null){
					    if(pingChart.series[0].data[0].x < tenMinutesAgo){
					    	shift = true;
					    }
				    }

				    if(jsonData.Time * 1000 > tenMinutesAgo){
				    	pingChart.series[0].addPoint({x: jsonData.Time * 1000, y: Math.round(jsonData.Latency)}, false, shift);
				    }
			    }


		    });

		    pingChart.redraw();
		    pingHistoryChart.redraw();
		}
	});
}

function handleWSMessage(event){
    if(webSocketFailed){
      webSocketFailed = false;
      console.log("Websocket connected")
      clearInterval(webSocketRetryTicker);
    }
    jsonData = JSON.parse(event.data);

    var historyChartChanged = false;

    if(typeof jsonData.mt != "undefined"){

    	if(jsonData.mt === "unhealthyPingResult"){
    		//console.log("Unhealthy ping result: " + jsonData)
    		if(jsonData.Time < 10000000){
    			return
    		}
    		pingChart.series[1].addPoint({x: jsonData.Time * 1000, y: Math.round(jsonData.Latency), Target: jsonData.Target}, false);
    		pingHistoryChart.series[2].addPoint({x: jsonData.Time * 1000, y: Math.round(jsonData.Latency), Target: jsonData.Target}, false);
    	}

    	if(jsonData.mt === "hops"){
    		var rows = "<tr><th>Hop</th><th>IP Address</th><th>Latency</th></tr>";
    		$.each(jsonData.Hops, function(i,hop){
    			rows += "<tr><td>"+(i+1)+"</td><td>"+hop.Target+"</td><td>"+Math.round(hop.Latency)+"ms</td></tr>";
			});

    		$('#hopTable tbody').html(rows);
    	}

    	if(jsonData.mt === "paths"){
    		var rows = "<tr><th>Path</th><th>Trips</th><th>Avg Latency</th><th>Avg Max Latency</th></tr>";
    		$.each(jsonData.Paths, function(i,path){
    			rows += "<tr><td>"+path.PathName+"</td><td>"+path.TripCount+"</td><td>"+Math.round(path.AvgLatency)+"ms</td><td>"+Math.round(path.MaxLatencyAvg)+"ms</td></tr>";
			});

    		$('#pathsTable tbody').html(rows);
    	}

    	if(jsonData.mt === "hosts"){
    		var rows = "<tr><th>Host</th><th>DNS Name</th><th>Trips</th><th>Min Latency</th><th>Avg Latency</th><th>Max Latency</th><th>&gt;100</th><th>&gt;400</th><th>&gt;700</th></tr>";
    		$.each(jsonData.Hosts, function(i,host){
    			rows += "<tr><td>"+host.HostName+"</td><td>"+host.HostDNSName+"</td><td>"+host.TripCount+"</td><td>"+Math.round(host.MinLatency)+"ms</td><td>"+Math.round(host.AvgLatency)+"ms</td><td>"+Math.round(host.MaxLatency)+"ms</td><td>"+host.HighLatency100+"</td><td>"+host.HighLatency400+"</td><td>"+host.HighLatency700+"</td></tr>";
			});

    		$('#hostsTable tbody').html(rows);
    	}

    	return
    }

    //console.log(jsonData);
    shift = false;
    if(pingChart.series[0].data.length > 600){
    	shift = true;
    }
    pingChart.series[0].addPoint({x: jsonData.Time * 1000, y: Math.round(jsonData.Latency)}, false, shift);
    historyX += jsonData.Time * 1000;
    avgdHistoryY += jsonData.Latency;

    if(minHistoryY > jsonData.Latency){
    	minHistoryY = jsonData.Latency;
    }
    if(maxHistoryY < jsonData.Latency){
    	maxHistoryY = jsonData.Latency;
    }

    if(pingChart.series[1].data.length > 0){
	    if(pingChart.series[1].data[0].x < pingChart.series[0].data[0].x){
	    	pingChart.series[1].data[0].remove(); // Remove Unhealthy Ping Trace Result that is older than the oldest Live Ping history point
	    }
    }


    chartDownRezCountdown--;
    if(chartDownRezCountdown === 0){
    	historyChartChanged = true;
    	console.log("Triggered downrez");
    	chartDownRezCountdown = 30;
    	pingHistoryChart.series[0].addPoint({x: historyX/chartDownRezCountdown, y: Math.round(avgdHistoryY/chartDownRezCountdown)}, false);
    	pingHistoryChart.series[1].addPoint([historyX/chartDownRezCountdown, Math.round(minHistoryY), Math.round(maxHistoryY)], false);
    	historyX = 0;
    	maxHistoryY = 0;
    	minHistoryY = 999999;
    	avgdHistoryY = 0;
    }

    pingChart.redraw(false);
    
    if(historyChartChanged){
    	pingHistoryChart.redraw(false);
    }
    
  }
  
  function handleWSClose(event){
    if(!webSocketFailed){
      console.log("Websocket disconnected")
      webSocketFailed = true;
      webSocketRetryTicker = setInterval(connect, 5000);
    }
  }

  function connect(){
    console.log("Attempting to connect")

    webSocket = new WebSocket("ws://localhost:14445/ws");
    webSocket.onmessage = function(event){handleWSMessage(event);}
    webSocket.onclose = function(event){handleWSClose(event);}
  }
</script>
</head>
<body style='background-color: #292923; color:white;'>
<div id='pingHistoryChart'></div>
<br/>
<div id='pingChart'></div>
<br/>
<div class='container'>
	<table class='table table-striped table-condensed table-bordered' id='hopTable'><tbody></tbody></table>
	<table class='table table-striped table-condensed table-bordered' id='pathsTable'><tbody></tbody></table>
	<table class='table table-striped table-condensed table-bordered' id='hostsTable'><tbody></tbody></table>
<div>
</body>
</html>
`, jquery, boostrapCSS, bootstrapJS, bootstrapSlateCSS, highchartsJS, highchartsMoreJS, highchartsDarkUnicaThemeJS)

func wsHandler(w http.ResponseWriter, req *http.Request) {
	ws, err := wsUpgrader.Upgrade(w, req, nil)
	if err != nil {
		elog.Print(err)
		return
	}

	c := &wsClient{
		send: make(chan []byte, maxMessageSize),
		ws:   ws,
	}

	wsHub.register <- c

	go c.writePump()
	go c.readPump()
}

func historyHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	resultRequestChan <- 1
	historyJSON := <-resultResponseChan

	w.Write(historyJSON)
}

func baseHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(dashboard))
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type hub struct {
	clients    map[*wsClient]bool
	broadcast  chan string
	register   chan *wsClient
	unregister chan *wsClient

	content string
}

type wsClient struct {
	ws   *websocket.Conn
	send chan []byte
}

var wsHub = hub{
	broadcast:  make(chan string),
	register:   make(chan *wsClient),
	unregister: make(chan *wsClient),
	clients:    make(map[*wsClient]bool),
	content:    "",
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
			c.send <- []byte(h.content)
			break

		case c := <-h.unregister:
			_, ok := h.clients[c]
			if ok {
				delete(h.clients, c)
				close(c.send)
			}
			break

		case m := <-h.broadcast:
			h.content = m
			h.broadcastMessage()
			break
		}
	}
}

func (h *hub) broadcastMessage() {
	for c := range h.clients {
		select {
		case c.send <- []byte(h.content):
			break

		// We can't reach the client
		default:
			close(c.send)
			delete(h.clients, c)
		}
	}
}

func (c *wsClient) readPump() {
	defer func() {
		wsHub.unregister <- c
		c.ws.Close()
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		wsHub.broadcast <- string(message)
	}
}

func (c *wsClient) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *wsClient) write(mt int, message []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, message)
}
