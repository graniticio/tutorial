package main

import (
	gy "github.com/graniticio/granitic-yaml/v2"
	"recordstore/bindings"
)

func main() {
	gy.StartGraniticWithYaml(bindings.Components())
}
