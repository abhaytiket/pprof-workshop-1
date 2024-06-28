# app2

This golang application will be used for a demo on how to do cpu profiling using [benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks) feature in Golang. CPU profiling is performed to get an idea about what functions are running in CPU during the lifespan of a Go program. It is handy to find out which function(s) are responsible for slowing down your Go application. Although it doesn't give any idea if a Goroutine is sleeping or waiting for some other function to get executed.

To run benchmark and generate a cpu profile run following command
```
go test -run=XXX -cpuprofile cpu.prof -bench .
```

This will generate cpu profile in `cpu.prof` file. Then to visualize and analyze this profile data via pprof web interface we need to run following command on terminal
```
go tool pprof -http=0.0.0.0:8888 cpu.prof
```
