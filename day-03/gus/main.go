package main

import (
	utils2 "adventofcode/utils"
	"fmt"
	"strings"
)

const fileName = "../input-data.txt"

var itemTypePriority map[string]int

func generateItemTypePriority() map[string]int {
	letters := "abcdefghijklmnopqrstuvwxyz"
	letters += strings.ToUpper(letters)
	data := strings.Split(letters, "")
	typePriority := make(map[string]int)
	for index, v := range data {
		typePriority[v] = index + 1
	}
	return typePriority
}

func getItemTypePriority() map[string]int {
	if itemTypePriority == nil {
		itemTypePriority = generateItemTypePriority()
	}
	return itemTypePriority

}

func calculateItemTypePriority(item string) int {
	itemTypePriority := getItemTypePriority()
	v, _ := itemTypePriority[item]
	return v
}

func splitStringInMiddle(data string) []string {
	splitData := strings.Split(data, "")
	splitDataLen := len(splitData)
	splitDataMiddle := splitDataLen / 2
	partOne := strings.Join(splitData[0:splitDataMiddle], "")
	partTwo := strings.Join(splitData[splitDataMiddle:], "")

	return []string{partOne, partTwo}
}

func findRepeatedItem(parts []string) string {
	var itemRepetead string
	letterCounter := make(map[string]int)
	firstPart := removeDuplicatedElements(parts[0])
	for _, l := range strings.Split(firstPart, "") {

		for _, p := range parts[1:] {
			if strings.Contains(p, l) {
				letterCounter[l] += 1
			}
		}
	}
	itemRepeteadCounter := 0
	for k, v := range letterCounter {
		if itemRepetead == "" {
			itemRepetead = k
			itemRepeteadCounter = v
			continue
		}

		if v > itemRepeteadCounter {
			itemRepetead = k
			itemRepeteadCounter = v
		}
	}
	return itemRepetead
}

func removeDuplicatedElements(data string) string {
	letters := strings.Split(data, "")
	s := ""
	for _, l := range letters {
		if !strings.Contains(s, l) {
			s += l
		}
	}
	return s
}

func sumRepeatedItemPriority(data []string) int {
	sum := 0
	for _, row := range data {
		splitRow := splitStringInMiddle(row)
		repeteadItem := findRepeatedItem(splitRow)
		sum += calculateItemTypePriority(repeteadItem)
	}

	return sum
}

func sumRepeatedItemInGroup(data []string) int {
	sum := 0
	for i := 0; (i + 3) <= len(data); {
		group := []string{data[i], data[i+1], data[i+2]}
		repeteadItem := findRepeatedItem(group)
		sum += calculateItemTypePriority(repeteadItem)
		i += 3
	}
	return sum
}

func partOne() {
	data := utils2.ReadFileLines(fileName)
	total := sumRepeatedItemPriority(data)
	fmt.Printf("%d", total)
}

func partTwo() {
	data := utils2.ReadFileLines(fileName)
	total := sumRepeatedItemInGroup(data)
	fmt.Printf("%d", total)
}

func main() {
	partTwo()
}
