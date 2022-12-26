package main

import (
	"reflect"
	"testing"
)

func TestCalculateItemTypePriority(t *testing.T) {
	priorityTests := []struct {
		letter string
		value  int
	}{
		{letter: "a", value: 1},
		{letter: "Z", value: 52},
	}
	for _, tt := range priorityTests {
		t.Run("priority of "+tt.letter, func(t *testing.T) {
			actual := calculateItemTypePriority(tt.letter)
			expected := tt.value
			if actual != expected {
				t.Errorf("got %d want %d", actual, expected)
			}
		})
	}
}

func TestSplitStringInMiddle(t *testing.T) {
	given := "abcaBC"
	expected := []string{"abc", "aBC"}

	actual := splitStringInMiddle(given)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %s want %s given %s", actual, expected, given)
	}

}

func TestFindRepeatedItem(t *testing.T) {
	t.Run("Find repeated in two items", func(t *testing.T) {
		given := []string{"abc", "afg"}
		expected := "a"
		actual := findRepeatedItem(given)

		if actual != expected {
			t.Errorf("got %s want %s given %v", actual, expected, given)
		}
	})

	t.Run("Find repeated in three items", func(t *testing.T) {
		given := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}
		expected := "r"

		actual := findRepeatedItem(given)

		if actual != expected {
			t.Errorf("got %s want %s given %v", actual, expected, given)
		}
	})
}

func TestSumRepeatedItemPriority(t *testing.T) {
	given := []string{"abcafg", "aijalm"}
	expected := 1 + 1

	actual := sumRepeatedItemPriority(given)

	if actual != expected {
		t.Errorf("got %d want %d given %v", actual, expected, given)
	}
}

func TestSumRepeatedItemInGroup(t *testing.T) {
	given := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}
	expected := 18 + 52 //r,Z

	actual := sumRepeatedItemInGroup(given)

	if actual != expected {
		t.Errorf("got %d want %d given %v", actual, expected, given)
	}
}

func TestRemoveDuplicatedElements(t *testing.T) {
	given := "MMFFFp"
	expected := "MFp"

	actual := removeDuplicatedElements(given)

	if actual != expected {
		t.Errorf("got %s want %s given %s", actual, expected, given)
	}

}
