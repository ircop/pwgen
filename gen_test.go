package pwgen

import (
	"fmt"
	"strings"
	"testing"
)

func TestPwgen_Generate(t *testing.T) {
	p := New(Defaults())

	for i := 0; i < 5; i++ {
		pw := p.Generate()
		if pw == "" {
			t.Fatal("Empty password")
		}
		fmt.Printf("password: '%s'\n", pw)
	}
}


func TestOwnDict(t *testing.T) {
	p := NewWithDict(Defaults(), []string{"testword"})

	for i := 0; i < 5; i++ {
		pw := p.Generate()
		if !strings.Contains(strings.ToLower(pw), "testword") {
			t.Fatal("Wrong password: does not contain testword: " + pw)
		}
		fmt.Printf("password: '%s'\n", pw)
	}
}