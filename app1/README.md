# app1

This application is a very basic web service developed using [Chi](https://github.com/go-chi/chi). We will be initially using [hey](https://github.com/rakyll/hey) for load testing while we generate a heap profile as we have also added [benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks) in this project.

To run a basic load test with 200 clients using hey for 10 minute duration we can run command after executing `main.go`
```
hey -z 10m http://127.0.0.1:8090/
```

Now open another terminal and as we mentioned that pprof can capture the profile data via http as well we will be using command
```
go tool pprof -http=0.0.0.0:8888 http://localhost:6060/debug/pprof/heap
```
to visualise and analyse the heap profile data using pprof.

But the correct way is to use benchmarks. Run following command to generate cpu profile
```
go test -run=XXX -cpuprofile cpu.prof -bench .
```

And then to analyse the cpu profile run
```
go tool pprof -http=0.0.0.0:8888 cpu.prof
```

Now the problem is that currently benchmarks are running in a serial order but if let's say we want to run benchmarks in parallel then we will have to switch to `B.RunParallel` to facilitate it. Check `app3` for an example.

