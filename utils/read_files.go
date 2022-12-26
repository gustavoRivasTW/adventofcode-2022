package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
func ReadFileLines(fileName string) []string {
	lines := strings.Split(ReadFile(fileName), "\n")
	return lines
}
