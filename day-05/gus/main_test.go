package main

import (
	"reflect"
	"testing"
)

func TestReadActualStack(t *testing.T) {
	given := []string{"Z,N", "M,C,D", "P"}
	expected := map[string][]string{
		"1": {"Z", "N"},
		"2": {"M", "C", "D"},
		"3": {"P"},
	}

	actual := readActualStack(given)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %s want %s given %s", actual, expected, given)
	}
}

func TestNewMovement(t *testing.T) {
	given := "move 1 from 2 to 1"
	expected := Movement{
		origin:  "2",
		destiny: "1",
		amount:  1,
	}

	actual := newMovement(given)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v given '%s'", actual, expected, given)
	}
}

func TestMoveCratesFromStackToAnother(t *testing.T) {

	priorityTests := []struct {
		movement string
		expected map[string][]string
	}{
		{
			movement: "move 1 from 2 to 1",
			expected: map[string][]string{
				"1": {"Z", "N", "Q", "R", "T"},
				"2": {"P", "Y", "G"},
				"3": {"A", "C", "K"},
			},
		},
		{
			movement: "move 1 from 1 to 2",
			expected: map[string][]string{
				"1": {"Z", "N", "Q"},
				"2": {"P", "Y", "G", "T", "R"},
				"3": {"A", "C", "K"},
			},
		},
		{
			movement: "move 3 from 3 to 1",
			expected: map[string][]string{
				"1": {"Z", "N", "Q", "R", "A", "C", "K"},
				"2": {"P", "Y", "G", "T"},
				"3": {},
			},
		},
	}

	for _, tt := range priorityTests {
		t.Run("Movement "+tt.movement, func(t *testing.T) {
			given := map[string][]string{
				"1": {"Z", "N", "Q", "R"},
				"2": {"P", "Y", "G", "T"},
				"3": {"A", "C", "K"},
			}
			movement := newMovement(tt.movement)

			actual := moveCratesFromStackToAnother(given, movement)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got %s want %s given %s and %v", actual, tt.expected, given, movement)
			}
		})
	}

}

func TestConcatTopCrates(t *testing.T) {
	given := map[string][]string{
		"1": {"Z", "N", "Q", "R"},
		"2": {"P", "Y", "G", "T"},
		"4": {"N", "M", "I"},
		"3": {"A", "C", "K"},
	}
	expected := "RTKI"

	actual := concatTopCrates(given)

	if actual != expected {
		t.Errorf("got %s want %s given %s ", actual, expected, given)
	}
}
