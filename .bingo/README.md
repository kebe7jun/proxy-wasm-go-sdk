# Binary dependencies

This directory supports tools written in go, but used in `make`. Tools like `goimports` are built
on-demand, and their dependencies do not affect the main project [go.mod](../go.mod).

The contents here were generated by [bingo](https://github.com/bwplotka/bingo) but have no runtime
dependency on it in any way.

Integration in [Makefile](../Makefile) requires inclusion of the [Variables.mk](Variables.mk), then
reference as a variable like `$(GOIMPORTS)`.

This differs slightly from the defaults of the Bingo project:
* ENV variable support is not included as this project only uses `make`