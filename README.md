# gopher2600_performance_profiles

The `cpu.profile` files were created with go version `1.21.7` and `1.22.0` and
stored in the appropriately named folder

They are profiles of the `gopher2600` project running in `performance` mode. The
project can be found here https://github.com/JetSetIlly/Gopher2600

The specific revision is `d1482c139ab9c78b18902426e87367a58cf8b101`

The project was compiled with `go build`. The executables, compiled for
`linux/amd64`, are included for reference.

The invocation of the executable for both go versions is

`gopher2600 performance -profile=cpu -duration=5m Cosmcark.bin`

`Cosmcark.bin` is an example VCS binary file that can be found in the
root of this repository

### Results

After five minutes the profiling will end and will output a performance summary
to the terminal.

In the case of the specific profiles in this repository, these were the results
at the time of collection

For version 1.21.7

`146.28 fps (43883 frames in 300.00 seconds) 243.8%`

For version 1.22.0

`133.48 fps (40044 frames in 300.00 seconds) 222.5%`

In this instance, this is a *drop in performance of 8.75%*

`(146.28 - 133.48) / 146.28 * 100`

The precise figures will vary from execution to execution but the value is
typical of every execution comparision tried. 

### A note about compilation

The full `gopher2600` project includes components that make calls to C libraries
via cgo (SDL, OpenGL, etc.) So although the compilation process requires and
compiles these components, they are NOT used when the program is invoked in the
`performance` mode. As we can see through the profile files, the code executed
is pure Go code.
