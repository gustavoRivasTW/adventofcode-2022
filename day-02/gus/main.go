package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const fileName = "../input-day-02.txt"

func readFileLines() [][]string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")
	data := [][]string{}
	for _, line := range lines {
		match := strings.Split(line, " ")
		data = append(data, match)
	}

	return data
}

func getScoreForShape(shape string) int {
	switch shape {
	case "A", "X": //Rock
		return 1
	case "B", "Y": //Paper
		return 2
	case "C", "Z": //Scissors
		return 3
	default:
		return 0
	}
}

func getMatchOutcome(oponent, you string) int {
	lostPoint := 0
	drawPoint := 3
	winPoint := 6

	if oponent == "A" { //Rock
		switch you {
		case "X":
			return drawPoint
		case "Y":
			return winPoint
		case "Z":
			return lostPoint
		}
	}

	if oponent == "B" { //Paper
		switch you {
		case "X":
			return lostPoint
		case "Y":
			return drawPoint
		case "Z":
			return winPoint
		}
	}

	if oponent == "C" { //Scissors
		switch you {
		case "X":
			return winPoint
		case "Y":
			return lostPoint
		case "Z":
			return drawPoint
		}
	}

	return 0
}

func getScoreForSingleRound(oponent, you string) int {
	score := getScoreForShape(you) + getMatchOutcome(oponent, you)
	return score
}

func applyStrategy(oponent, action string) string {
	if action == "X" { //Loose
		if oponent == "A" {
			return "Z"
		}
		if oponent == "B" {
			return "X"
		}
		if oponent == "C" {
			return "Y"
		}
	} else if action == "Y" { //draw
		if oponent == "A" {
			return "X"
		}
		if oponent == "B" {
			return "Y"
		}
		if oponent == "C" {
			return "Z"
		}

	} else if action == "Z" { // win
		if oponent == "A" {
			return "Y"
		}
		if oponent == "B" {
			return "Z"
		}
		if oponent == "C" {
			return "X"
		}
	}
	return "N"
}

func sumScore() {
	lines := readFileLines()
	totalScore := 0
	for _, line := range lines {
		mymove := applyStrategy(line[0], line[1])
		totalScore += getScoreForSingleRound(line[0], mymove)
	}
	fmt.Printf("\nTotalScore: %d", totalScore)
}

func main() {
	sumScore()
}
