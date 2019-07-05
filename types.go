package main

import "github.com/brendanporter/quack"

type UnhealthyPingResult struct {
	Target      string
	Latency     float64
	Time        int64
	MessageType string `json:"mt"`
}

type TraceResult struct {
	Hops        []quack.PingResult
	MessageType string `json:"mt"`
}

type PathsResult struct {
	Paths       map[string]*PathStats
	MessageType string `json:"mt"`
}

type HostsResult struct {
	Hosts       map[string]*HostStats
	MessageType string `json:"mt"`
}

type HostStats struct {
	HostName       string
	HostDNSName    string
	AvgLatency     float64
	MaxLatency     float64
	MinLatency     float64
	TripCount      int
	HighLatency100 int
	HighLatency400 int
	HighLatency700 int
}

type PathStats struct {
	PathName      string
	MaxLatencyAvg float64
	AvgLatency    float64
	TripCount     int
}
