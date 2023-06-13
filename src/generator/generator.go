package generator

import (
	"lucx-lab/passGenerator/src/options"
	"math/rand"
	"time"
)

const (
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers   = "0123456789"
	Symbols   = "!@#$%^&*()_+{}[]:;?/.,"
)

type Generator struct {
	Options     options.Options
	CharOptions []string
}

func (g Generator) SetCharOptions() []string {
	var opt []string

	if g.Options.Lowercase {
		opt = append(opt, Lowercase)
	}
	if g.Options.Uppercase {
		opt = append(opt, Uppercase)
	}
	if g.Options.Numbers {
		opt = append(opt, Numbers)
	}
	if g.Options.Symbols {
		opt = append(opt, Symbols)
	}
	return opt
}

func (g Generator) Generate() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]byte, g.Options.Length)

	for i := 0; i < g.Options.Length; i++ {
		rType := r.Intn(len(g.CharOptions))
		res[i] = RandomElement(r, g.CharOptions[rType])
	}
	return string(res)
}

func NewGenerator(options options.Options) Generator {
	var g Generator

	g.Options = options
	g.CharOptions = g.SetCharOptions()
	return g
}

func RandomElement(r *rand.Rand, s string) byte {
	return s[r.Intn(len(s))]
}
