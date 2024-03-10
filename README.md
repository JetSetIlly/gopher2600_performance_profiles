# gopher2600_performance_profiles

main.go is the minimum code necessary to run the [gopher2600](https://github.com/JetSetIlly/Gopher2600)
project in `performance` measurment mode.

Build the executable with 'go build .'

PGO files are supplied. `default.pgo` is suitable for Go v1.22.0 and `default_1.21.0.pgo` for earlier versions

Run the executable with no arguments. By default each invocation of the tool will run for 1
minute.

Benchstat and CPU profiles, etc. can be specified by command line argument

```
Usage of gopher2600_performance_profiles:
  -benchstat
    	record benchstats to file
  -duration string
    	duration of execution (default "1m")
  -profile string
    	run with profiling: CPU, MEM, TRACE, ALL (default "none")
  -uncapped
    	run emulation with no FPS cap (default true)
```

Output of `-benchstat` can be used with the [benchstat tool](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat). It can also be used with the program in the `smooth` directory. This program compares two benchstat files and tries to identify which of the two is "smoother". ie. has less jitter in the input data
