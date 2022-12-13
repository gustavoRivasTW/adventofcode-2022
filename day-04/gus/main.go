package main

import (
	utils2 "adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func getArrayFromPair(pair string) []int {
	pairs := strings.Split(pair, "-")
	pairOne, _ := strconv.Atoi(pairs[0])
	var start = int(pairOne)
	pairTwo, _ := strconv.Atoi(pairs[1])
	var finish = int(pairTwo)

	return []int{start, finish}
}

func CheckPairContains(firstPair string, secondPair string) bool {
	firstArray := getArrayFromPair(firstPair)
	secondArray := getArrayFromPair(secondPair)
	if firstArray[0] <= secondArray[0] {
		if firstArray[1] >= secondArray[1] {
			return true
		}
	}

	if secondArray[0] <= firstArray[0] {
		if secondArray[1] >= firstArray[1] {
			return true
		}
	}

	return false
}

func countPairFullyContains(given []string) int {
	counter := 0
	for _, pairLine := range given {
		pairs := strings.Split(pairLine, ",")
		if CheckPairContains(pairs[0], pairs[1]) {
			counter++
		}
	}
	return counter
}

func main() {
	data := utils2.ReadFileLines("../input-data.txt")
	total := countPairFullyContains(data)
	fmt.Printf("%d", total)
}
