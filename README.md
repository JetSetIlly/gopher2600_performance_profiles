# gopher2600_performance_profiles

main.go is the minimum code necessary to run the [gopher2600](https://github.com/JetSetIlly/Gopher2600)
project in `performance` measurment mode.

Build the executable with 'go build .'

Run the executable with no arguments. Each invocation of the tool will run for 5
minutes. 

### Results

The two folders `1.21.7` and `1.22.0` contain the cpu profiles for an execution
of the tool when compiled with the respective go version.

The output to the terminal for each of these executions was as follows:

For version 1.21.7

> 147.04 fps (44112 frames in 300.00 seconds) 245.1%

For version 1.22.0

> 137.57 fps (41271 frames in 300.00 seconds) 229.3%

In this instance, this is a *drop in performance of 6.45%

`(147.04 - 137.57) / 147.04 * 100`

The precise figures will vary from execution to execution (and from machine to
machine) but the value is typical of every execution comparision tried on the
developer's PC.
