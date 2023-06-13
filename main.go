package main

import (
	"fmt"
	"github.com/LuLmth/passGenerator/src/generator"
	"github.com/LuLmth/passGenerator/src/options"
)

func main() {
	options := options.NewOptions()
	gen := generator.NewGenerator(options)

	fmt.Println(gen.Generate())
}
