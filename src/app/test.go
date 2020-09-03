package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {

	got := greeting("Code.education Rocks!")
	want := "<b>Code.education Rocks!</b>"

	if got != want {
		t.Errorf("Not returning: got %s want %s!", got, want)
	}
}