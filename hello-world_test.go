package main

import (
	"bytes"
	"fmt"
	"testing"
)

func helloWorld(out *bytes.Buffer) {
	fmt.Fprintln(out, "Hello, World!")
}

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!\n"
	buffer := bytes.Buffer{}
	helloWorld(&buffer)

	got := buffer.String()
	if got != expected {
		t.Errorf("Unexpected output: got %q, want %q", got, expected)
	}
}

