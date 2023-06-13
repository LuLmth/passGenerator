package options

import "testing"

func TestNewOptions(t *testing.T) {
	options := NewOptions()

	if options.Length != 12 {
		t.Errorf("Expected length to be 12, got %d", options.Length)
	}

	if options.Lowercase != true {
		t.Errorf("Expected lowercase to be true, got %t", options.Lowercase)
	}

	if options.Uppercase != true {
		t.Errorf("Expected uppercase to be true, got %t", options.Uppercase)
	}

	if options.Numbers != true {
		t.Errorf("Expected numbers to be true, got %t", options.Numbers)
	}

	if options.Symbols != false {
		t.Errorf("Expected symbols to be false, got %t", options.Symbols)
	}
}
