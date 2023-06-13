package main

import (
	"fmt"
	"lucx-lab/passGenerator/src/generator"
	"lucx-lab/passGenerator/src/options"
)

func main() {
	options := options.NewOptions()
	gen := generator.NewGenerator(options)

	fmt.Println(gen.Generate())
}
