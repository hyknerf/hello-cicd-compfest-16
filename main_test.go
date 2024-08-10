package main

import "testing"

func TestSayHello(t *testing.T) {
	result := sayHello("world")
	expected := "Hello world"
	if result != expected {
		t.Errorf("should produce \"Hello world\" but got %s", result)
	}
}
