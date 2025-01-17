package main

var confTemp=`port: 12306
servername: "test"
prometheus:
  switch_on: true
  port: 8080
limit:
  switch_on: true
  qps: 10
  size: 1000
  type: lpLimit
  timediff: -1
logs:
  chansize: 100
  loglevel: debug
register:
  switch_on: true
  addr: 
    - localhost:2379
  timeOut: 1s
  registerPath: "/test"
  heartBeat: 5
trace:
  switch_on: true
  report_addr: ""
  sample_type: const
  sample_rate: 1
`