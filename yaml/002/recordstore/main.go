package main

import (
	gy "github.com/graniticio/granitic-yaml/v2"
)
import "recordstore/bindings"

func main() {
	gy.StartGraniticWithYaml(bindings.Components())
}
