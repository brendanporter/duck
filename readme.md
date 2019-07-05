# Duck long-term network latency tester

Goals
---

- In-place replacement for terminal `ping` command
- Capture high latency events in the background and make them clear on review


Features
---
- In-terminal colored ASCII latency charting
- Live web dashboard for visualizing latency over extended tests
  - Long-term chart with 30-second averaged samples
  - Short-term chart (last 10 minutes) with live results
  - Traceroute results
  - Traceroute Path tracking with latency averaging
  - Individual traceroute host tracking with latency averaging and incident counting
- Keeps history of ping performance older than 10 minutes as down-rezzed 30-second sample averages
- Traceroute performed every 30 seconds
  - Paths stored and path traversal tracked
  - Host DNS lookup
  - Average and Max latency tracked over time