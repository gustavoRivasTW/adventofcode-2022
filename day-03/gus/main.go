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

func findRepeatedItem(partOne string, partTwo string) string {
	var itemRepetead string
	for _, v := range strings.Split(partOne, "") {
		if strings.Contains(partTwo, v) {
			itemRepetead = v
		}
	}
	return itemRepetead
}

func sumRepeatedItemPriority(data []string) int {
	sum := 0
	for _, row := range data {
		splitRow := splitStringInMiddle(row)
		repeteadItem := findRepeatedItem(splitRow[0], splitRow[1])
		sum += calculateItemTypePriority(repeteadItem)
	}

	return sum
}

func main() {
	data := utils2.ReadFileLines(fileName)
	total := sumRepeatedItemPriority(data)
	fmt.Printf("%d", total)
}
