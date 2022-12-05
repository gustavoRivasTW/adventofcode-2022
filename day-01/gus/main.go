package main

// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

// Leer el archivo
// Crear Array bi-dimensional
// Obtener el con mayor calorias

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func readFile() string {
	bytes, err := ioutil.ReadFile("../input-day-01.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func max(values []int) int {
	maxVal := 0
	for _, v := range values {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}

func reverse(values []int) []int {
	reverse_values := []int{}
	for i := (len(values) - 1); i >= 0; i-- {
		reverse_values = append(reverse_values, values[i])
	}
	return reverse_values
}

func topThree(values []int) []int {
	sort.Ints(values[:]) // convert array to slice and sort the array
	reverse_values := reverse(values)
	return reverse_values[:3]
}

func ejercicio_uno() {
	data := readFile()
	lines := strings.Split(data, "\n")

	var total_calories_per_elf []int
	var calories_per_elf []int
	for _, line := range lines {
		if line == "" {
			total_calories_per_elf = append(total_calories_per_elf, sum(calories_per_elf))
			calories_per_elf = []int{}
		} else {
			calories, _ := strconv.Atoi(line)
			calories_per_elf = append(calories_per_elf, int(calories))
		}
	}
	value := max(total_calories_per_elf)
	fmt.Printf("Max Value : %d", value)
	sumTop := sum(topThree(total_calories_per_elf))
	fmt.Printf("\nSum Top 3 : %d", sumTop)

}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	ejercicio_uno()
}
