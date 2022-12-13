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

func generateArrayWithPair(pair string) []int {
	values := getArrayFromPair(pair)
	start := values[0]
	end := values[1]
	numbersInPair := []int{}
	for start <= end {
		numbersInPair = append(numbersInPair, start)
		start++
	}
	return numbersInPair
}
func CheckPairOverlap(firstPair string, secondPair string) bool {
	firstArray := generateArrayWithPair(firstPair)
	secondArray := generateArrayWithPair(secondPair)
	for _, v := range firstArray {
		for _, x := range secondArray {
			if v == x {
				return true
			}
		}
	}

	return false
}

func countPairOverlap(given []string) int {
	counter := 0
	for _, pairLine := range given {
		pairs := strings.Split(pairLine, ",")
		if CheckPairOverlap(pairs[0], pairs[1]) {
			counter++
		}
	}
	return counter
}

func main() {
	data := utils2.ReadFileLines("../input-data.txt")
	totalcontains := countPairFullyContains(data)
	totalOverlap := countPairOverlap(data)

	fmt.Printf("countPairFullyContains: %d, countPairOverlap: %d ", totalcontains, totalOverlap)
}
