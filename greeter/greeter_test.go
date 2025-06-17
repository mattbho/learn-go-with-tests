package main

import (
	"bytes"
	"testing"
)

func TestGreeter(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Matt")

	got := buffer.String()
	want := "Hello, Matt"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
