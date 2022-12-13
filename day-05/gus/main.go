package main

import (
	"fmt"
	"strings"
)

type Movement struct {
	origin  int
	destiny int
	amount  int
}

func readActualStack(given []string) map[string][]string {
	actualStack := make(map[string][]string)
	for i, v := range given {
		key := fmt.Sprint(i + 1)
		crates := strings.Split(v, ",")
		actualStack[key] = crates
	}
	return actualStack
}

func translateMovesToNumbers(move string) Movement {
	return Movement{}
}

func moveCratesFromStackToAnother(stacks map[string][]string, movement Movement) map[string][]string {
	newStack := make(map[string][]string)

	return newStack
}
