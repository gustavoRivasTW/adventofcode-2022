package main

// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

// Leer el archivo
// Crear Array bi-dimensional
// Obtener el con mayor calorias

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func leer_archivo() string {
	datosComoBytes, err := ioutil.ReadFile("../input-day-01.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(datosComoBytes)
}

func maxValue(values []int) int {
	maxVal := 0
	for _, v := range values {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}

func ejercicio_uno() {
	datos := leer_archivo()
	lines := strings.Split(datos, "\n")

	var elf_calories []int
	var line_values []int
	for _, line := range lines {
		if line == "" {
			sum_calories_elf := sum(line_values)
			elf_calories = append(elf_calories, int(sum_calories_elf))
			line_values = []int{}
		} else {
			lineVal, _ := strconv.Atoi(line)
			line_values = append(line_values, int(lineVal))
		}
	}
	value := maxValue(elf_calories)
	fmt.Println(value)

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
