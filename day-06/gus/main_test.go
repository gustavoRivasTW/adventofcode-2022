package main

import (
	"testing"
)

func TestFindMarker(t *testing.T) {
	given := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	expected := 19

	actual := findMarker(given)

	if actual != expected {
		t.Errorf("got %d want %d given %s", actual, expected, given)
	}

}
