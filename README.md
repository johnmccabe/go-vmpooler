# go-vmpooler

[![Go Report Card Badge]][Go Report Card]
[![GoDoc Badge]][GoDoc]
[![Release Badge]][Release]

![vmpooler](images/go-vmpooler.jpg)

go-vmpooler provides a dependency free Go client library for working with [vmpooler](https://github.com/puppetlabs/vmpooler).

Vmpooler provides configurable 'pools' of instantly-available virtual machines running on VMWare VSphere. 

## Usage

Refer to the example programs in `examples/token` and `examples\vm` for some ideas on how to use the client.

*Note that the `token` example requires the `golang.org/x/crypto/ssh/terminal` package be installed, its used to suppress echoing of the password to stdout when prompting the user to enter it.*

Refer to the [godocs](https://godoc.org/github.com/johnmccabe/go-vmpooler) for the following packages for more info.

- [token](https://godoc.org/github.com/johnmccabe/go-vmpooler/token)
- [vm](https://godoc.org/github.com/johnmccabe/go-vmpooler/token)


## Stability

There is currently only an early v0.0.x release available, there may be breaking changes introduced, please raise an issue or reach out to me with feedback/suggestions.

[Go Report Card]: https://goreportcard.com/report/github.com/johnmccabe/go-vmpooler
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/johnmccabe/go-vmpooler
[GoDoc]: https://godoc.org/github.com/johnmccabe/go-vmpooler
[GoDoc Badge]: https://godoc.org/github.com/johnmccabe/go-vmpooler?status.svg
[Release]: https://github.com/johnmccabe/go-vmpooler/releases/latest
[Release Badge]: https://img.shields.io/github/release/johnmccabe/go-vmpooler.svg?style=flat-square
