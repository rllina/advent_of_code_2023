package main

import (
	"fmt"
	"strings"
)

var numberConversions = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getNumbersFromString(line string) int {
	//line[:4]
	
	//for key, value := range numberConversions {
	//	line = strings.ReplaceAll(line, key, value)
	//}
	//
	//numbers := strings.Split(line, "")
	//result, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
	//return result
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
	//file, _ := os.ReadFile("part2Input.txt")
	//input = string(file)
	//
	//getNumbersFromLines(input)
}
