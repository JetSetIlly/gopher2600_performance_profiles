# gopher2600_performance_profiles

main.go is the minimum code necessary to run the [gopher2600](https://github.com/JetSetIlly/Gopher2600)
project in `performance` measurment mode.

Build the executable with 'go build .'

Run the executable with no arguments. Each invocation of the tool will run for 1
minute. FPS is measured every second and recorded to a benchstat file with a
name suffices with the current go version. For example:

> benchstat_devel go1.22-ad943066f6 Thu Jul 20 21:39:57 2023 +0000.txt

