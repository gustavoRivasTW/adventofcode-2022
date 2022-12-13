package main

import (
	"adventofcode/utils"
	"reflect"
	"testing"
)

func TestReadActualStack(t *testing.T) {
	given := utils.ReadFileLines("../stacks.txt")

	expected := map[string][]string{
		"1": []string{"Z", "N"},
		"2": []string{"M", "C", "D"},
		"3": []string{"P"},
	}

	actual := readActualStack(given)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %s want %s given %s", actual, expected, given)
	}
}

func TestTranslateMovesToNumbers(t *testing.T) {
	given := "move 1 from 2 to 1"
	expected := Movement{
		origin:  2,
		destiny: 1,
		amount:  1,
	}
	actual := translateMovesToNumbers(given)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v given '%s'", actual, expected, given)
	}

}

func testMoveCratesFromStackToAnother(t *testing.T) {
	stacks := map[string][]string{
		"1": []string{"Z", "N"},
		"2": []string{"P"},
	}
	movement := Movement{
		origin:  1,
		destiny: 2,
		amount:  2,
	}

	expected := map[string][]string{
		"1": []string{},
		"3": []string{"P", "Z", "N"},
	}

	actual := moveCratesFromStackToAnother(stacks, movement)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %s want %s given %s and %v", actual, expected, stacks, movement)
	}

}
