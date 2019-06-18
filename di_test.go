package main

import "testing"
import "bytes"

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tony")
	got := buffer.String()
	expected := "Hello, Tony"
	if got != expected {
		t.Errorf("expected '%s' but got '%s'", expected, got)
	}
}
