# go-plugin-example

A quick example project using Hashicorp's 
[go-plugin](https://github.com/hashicorp/go-plugin) system.

This project attempts to show how 
[go-plugin](https://github.com/hashicorp/go-plugin) can be used to satisfy 
implementations of plugin code called dynamically from a core 
[Go](https://golang.org/) application.

[go-plugin](https://github.com/hashicorp/go-plugin) provides a system for plugin
code to be written in distince processes (implementing agreed upon interfaces).
These processes speak to a core Go application via an RPC mechanism.

> Note: quick and dirty example from a few hours trying to understand how
[go-plugin](https://github.com/hashicorp/go-plugin) works and can be used in 
different design patterns.

## Goals

This exmaple project demonstrates a simple architectural pattern with the
following goals.

* Go developer wishes to allow a plugin sytem for extensibility
* Go developer wishes to decouple external plugins from main application build
* Plugin developers can write/ship custom proprietary plugins if needed 

## Example Architecture

This example creates a main application that loads and calls plugins from 
plugin implmentations that have been compiled into plugin binaries.  These
plugins implment agreed upon interfaces and ahere to a filename pattern of
`<type>-<name`.  For example, the plugin named `greeter-foo` would represent 
a plugin of `type: greeter` and `implementation: foo`.  Here we use the term 
`type` for a class of plugins that implmenet a similar interface definition.

This example shows how one could interface with two plugin types, each with 
different implementations.

* [main.go](./main.go) - Example of a core GO application that wishes to
interact with plugins.
* [./plugin](./plugin) - Go package wrapping up a Manager for loading plugin 
binaries from specific directories.  It also contains interfaces that need to 
be 
* [./plugin/greeter.go](./plugin/greeter.go) - Interfaces for plugins of type 
`greeter`
* [./plugin/clubber.go](./plugin/greeter.go) - Interfaces for plugins of type
`clubber`
* [./plugin/config.go](./plugin/config.go) - Common items needed for configuring
underlying plugin clients/servers 
(abstracted by [go-plugin](https://github.com/hashicorp/go-plugin))
* [./plugin/manager.go](./plugin/manager.go) - Convenient package for loading/
using plugins by the main application.
* [./plugins](./plugins) - Individual plugin implementations.  In the real world
these would probably be seperate Git repos for 3rd party plugins.

## Setup

1. Make sure you have Go installed and setup
2. Clone the project 
3. Get dependencies
```bash
go get github.com/hashicorp/go-plugin
```
4. Run `make` in order to build and place the plugins in the expected location
```bash
make clean && make
```
5. Run the main program
```bash
go run main.go
```
