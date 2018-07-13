package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	expected := "Hello, World!"

	if got != expected 	{
		t.Errorf("got '%s' want '%s'", got, expected)
	}
	
}