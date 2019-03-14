package main

import "github.com/graniticio/granitic/v2"
import "recordstore/bindings"

func main() {
	granitic.StartGranitic(bindings.Components())
}
