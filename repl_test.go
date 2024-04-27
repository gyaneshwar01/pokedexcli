package main

import (
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	input := "hello, world!"
	got := cleanInput(input)

	want := []string{"hello,", "world!"}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
