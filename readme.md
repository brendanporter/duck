# Duck long-term network latency tester

How to use
--- 

- Find your binary and run from the command line as administrator/root (for ICMP listener)
  - Watch the pings roll in
  - Browse to http://localhost:14445 and watch the results
- Build from source with `go build`
  - Mac/Linux
    - Run duck with `sudo ./duck <target_ip>`
  - Windows
    - Run duck from administrator command prompt with `duck.exe <target_ip>`

Screenshots
---

![Image of Duck in Mac Terminal](https://brendanporter.github.io/duck-terminalui_v2.png)
![image of Duck in Browser](https://brendanporter.github.io/duck-webui.png)

Goals
---

- Be a better, in-place replacement for terminal `ping` command
- Capture high latency events in the background and make them clear on review


Features
---
- In-terminal colored ASCII latency charting
- Statistics printed every 30 seconds
- Live web dashboard for visualizing latency over extended tests
  - Long-term chart with 30-second averaged samples
  - Short-term chart (last 10 minutes) with live results
  - Traceroute results
  - Traceroute Path tracking with latency averaging
  - Individual traceroute host tracking with latency averaging and incident counting
- Keeps history of ping performance older than 10 minutes as derezzed 30-second sample averages
- Traceroute performed every 30 seconds
  - Paths stored and path traversal tracked
  - Host DNS lookup
  - Average and Max latency tracked over time
- Stores pings, paths, and hosts to JSON files every minute

Roadmap
---
- Derezzing of ping results after 10 minutes
  - Exception to derezzeing is any high-latency event, with 5 minutes of high resolution context on either side
- Variable control from WebUI
  - Payload size
  - Latency threshold adjustments
- Jitter visualization