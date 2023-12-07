package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input string

// look into the flag library

func getNumbersFromString(input string) int {
	arr := strings.Split(input, "")
	var numbers []string

	for _, element := range arr {
		if _, err := strconv.Atoi(element); err == nil {
			numbers = append(numbers, element)
		}
	}

	result, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
	return result
}

func getNumbersFromLines(lines string) int {
	linesArr := strings.Split(lines, "\n")
	var result int

	for _, element := range linesArr {
		lineNumbers := getNumbersFromString(element)
		result += lineNumbers
	}

	fmt.Println("********* result *********")
	fmt.Println(result)
	return result
}

func main() {
	file, _ := os.ReadFile("part1Input.txt")
	input = string(file)

	getNumbersFromLines(input)
}
