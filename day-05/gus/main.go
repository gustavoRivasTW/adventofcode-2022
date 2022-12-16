package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

type Movement struct {
	origin  string
	destiny string
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
	return Movement{words[3], words[5], toInt(words[1])}
}

func moveCratesFromStackToAnother(stacks map[string][]string, movement Movement) map[string][]string {
	//Append to Destiny
	for i := 1; i <= movement.amount; i++ {
		lastCrate := stacks[movement.origin][len(stacks[movement.origin])-i]
		stacks[movement.destiny] = append(stacks[movement.destiny], lastCrate)
	}
	// Remove from Origin
	stacks[movement.origin] = stacks[movement.origin][:len(stacks[movement.origin])-movement.amount]

	return stacks
}

func concatTopCrates(stacks map[string][]string) string {
	word := ""
	for i := 0; i < len(stacks); i++ {
		key := strconv.Itoa(i + 1)
		v := stacks[string(key)]
		word += v[len(v)-1]
	}
	return word
}

func toInt(s string) int {
	v, _ := strconv.Atoi(s)
	return int(v)
}

func main() {
	stackFile := utils.ReadFileLines("../stacks.txt")
	movementFile := utils.ReadFileLines("../input-data.txt")

	stack := readActualStack(stackFile)
	for _, v := range movementFile {
		movement := newMovement(v)
		stack = moveCratesFromStackToAnother(stack, movement)
	}
	word := concatTopCrates(stack)

	fmt.Print(word)

}
