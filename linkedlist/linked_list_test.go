package main

import "testing"

func TestHello(t *testing.T) {
	got := printAll(getExampleList())
	want := "0 10 20 30 40 50 60 70 80 90"

	if got != want {
		t.Errorf("you fucked up, moron! got '%s' want '%s'", got, want)
	}
}
