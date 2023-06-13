package options

import (
	"flag"
)

type Options struct {
	Length    int
	Lowercase bool
	Uppercase bool
	Numbers   bool
	Symbols   bool
}

func NewOptions() Options {
	var flags Options
	length := flag.Int("length", 12, "length of the password")
	lowercase := flag.Bool("lowercase", true, "use lowercase characters")
	uppercase := flag.Bool("uppercase", true, "use uppercase characters")
	numbers := flag.Bool("numbers", true, "use numbers")
	symbols := flag.Bool("symbols", false, "use symbols")

	flag.Parse()
	flags.Length = *length
	flags.Lowercase = *lowercase
	flags.Uppercase = *uppercase
	flags.Numbers = *numbers
	flags.Symbols = *symbols
	return flags
}
