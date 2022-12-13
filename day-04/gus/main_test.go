package main

import (
	"reflect"
	"testing"
)

func TestGetArrayFromPair(t *testing.T) {
	given := "2-8"
	expected := []int{2, 8}

	actual := getArrayFromPair(given)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %d want %d given %s", actual, expected, given)
	}
}
func TestCheckPairContains(t *testing.T) {
	priorityTests := []struct {
		pairOne       string
		pairTwo       string
		expectedValue bool
	}{
		{pairOne: "2-8", pairTwo: "3-7", expectedValue: true},
		{pairOne: "3-7", pairTwo: "2-8", expectedValue: true},
		{pairOne: "3-7", pairTwo: "5-9", expectedValue: false},
	}

	for _, tt := range priorityTests {
		t.Run("priority of "+tt.pairOne+","+tt.pairTwo, func(t *testing.T) {
			actual := CheckPairContains(tt.pairOne, tt.pairTwo)

			if actual != tt.expectedValue {
				t.Errorf("got %t want %t", actual, tt.expectedValue)
			}
		})
	}
}

func TestCountPairFullyContains(t *testing.T) {
	given := []string{"2-8,3-7", "3-7,2-8", "3-7,5-9"}
	expected := 2

	actual := countPairFullyContains(given)

	if actual != expected {
		t.Errorf("got %d want %d", actual, expected)
	}
}

func TestGenerateArrayWithPair(t *testing.T) {
	given := "2-8"
	expected := []int{2, 3, 4, 5, 6, 7, 8}

	actual := generateArrayWithPair(given)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %d want %d given %s", actual, expected, given)
	}
}

func TestCheckPairOverlap(t *testing.T) {
	priorityTests := []struct {
		pairOne       string
		pairTwo       string
		expectedValue bool
	}{
		{pairOne: "2-8", pairTwo: "3-7", expectedValue: true},
		{pairOne: "3-7", pairTwo: "2-8", expectedValue: true},
		{pairOne: "3-7", pairTwo: "5-9", expectedValue: true},
		{pairOne: "1-2", pairTwo: "2-4", expectedValue: true},
		{pairOne: "1-2", pairTwo: "8-9", expectedValue: false},
	}

	for _, tt := range priorityTests {
		t.Run("priority of "+tt.pairOne+","+tt.pairTwo, func(t *testing.T) {
			actual := CheckPairOverlap(tt.pairOne, tt.pairTwo)

			if actual != tt.expectedValue {
				t.Errorf("got %t want %t, given %s,%s", actual, tt.expectedValue, tt.pairOne, tt.pairTwo)
			}
		})
	}
}
