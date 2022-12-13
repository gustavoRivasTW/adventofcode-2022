package main

import (
	"fmt"
	"strconv"
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

func newMovement(move string) Movement {
	words := strings.Split(move, " ")
	return Movement{toInt(words[3]), toInt(words[5]), toInt(words[1])}
}

func moveCratesFromStackToAnother(stacks map[string][]string, movement Movement) map[string][]string {
	newStack := make(map[string][]string)

	return newStack
}

func toInt(s string) int {
	v, _ := strconv.Atoi(s)
	return int(v)
}
