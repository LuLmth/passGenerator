package generator

import (
	"lucx-lab/passGenerator/src/options"
	"math/rand"
	"testing"
	"time"
)

const (
	LengthTesting = 8
)

func TestGenerate(t *testing.T) {
	o := options.Options{Length: LengthTesting, Lowercase: true, Uppercase: false, Numbers: false, Symbols: false}
	g := NewGenerator(o)

	res := g.Generate()
	if len(res) != LengthTesting {
		t.Errorf("Expected %v to be 8", len(res))
	}
	for _, c := range res {
		if c < 'a' || c > 'z' {
			t.Errorf("Expected %v a lowercase", c)
		}
	}
}

func TestSetCharOptions(t *testing.T) {
	o := options.Options{Length: LengthTesting, Lowercase: true, Uppercase: true, Numbers: true, Symbols: true}
	var g Generator

	g.Options = o
	g.CharOptions = g.SetCharOptions()
	if len(g.CharOptions) != 4 {
		t.Errorf("Expected all options (lowercase, uppercase, numbers and symbols) enabled")
	}
}

func TestRandomElement(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := "abc"
	elem := RandomElement(r, s)
	if elem != 'a' && elem != 'b' && elem != 'c' {
		t.Errorf("Expected %v to be one of %v", elem, s)
	}
}
