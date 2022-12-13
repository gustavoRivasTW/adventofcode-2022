package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadFileLines(fileName string) []string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(bytes), "\n")
	return lines
}
