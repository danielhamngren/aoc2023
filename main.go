package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string) int {
	re := regexp.MustCompile(`[0-9]`)
	result := 0

	for i := range input {

		digits := re.FindAllString(input[i], -1)

		cal, err := strconv.Atoi(digits[0] + digits[len(digits)-1])

		if err != nil {
			panic(err)
		}

		result += cal

	}

	return result

}

func digitConv(x string) string {
	digits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}
	return digits[x]
}

func handleLine(x string) string {
	var first = make(map[string]int)
	var last = make(map[string]int)

	digits := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"0":     "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}

	for key, _ := range digits {
		first[key] = strings.Index(x, key)
		last[key] = strings.LastIndex(x, key)
	}

	var firstKey string
	firstIndex := 10000
	for key, val := range first {
		if val < firstIndex && val > -1 {
			firstIndex = val
			firstKey = key
		}
	}

	var lastKey string
	lastIndex := -1
	for key, val := range last {
		if val > lastIndex {
			lastIndex = val
			lastKey = key
		}
	}
	return digits[firstKey] + digits[lastKey]
}

func part2(input []string) int {
	re := regexp.MustCompile(`[0-9]`)
	result := 0

	// fmt.Println(input)

	reDigits := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|zero")

	for i := range input {
		input[i] = handleLine(input[i])
		input[i] = reDigits.ReplaceAllStringFunc(input[i], digitConv)
	}

	// fmt.Println(input)

	for i := range input {
		digits := re.FindAllString(input[i], -1)

		cal, err := strconv.Atoi(digits[0] + digits[len(digits)-1])

		if err != nil {
			panic(err)
		}

		result += cal

	}

	return result
}

func day1() {
	// input := []string{
	// 	"1abc2",
	// 	"pqr3stu8vwx",
	// 	"a1b2c3d4e5f",
	// 	"treb7uchet",
	// }

	// input := []string{
	// 	"two1nine",
	// 	"eightwothree",
	// 	"abcone2threexyz",
	// 	"xtwone3four",
	// 	"4nineeightseven2",
	// 	"zoneight234",
	// 	"7pqrstsixteen",
	// }

	// Read input from file
	content, err := os.ReadFile("input1-1")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	// result1 := part1(input)
	// fmt.Println(result1)

	result2 := part2(input)
	fmt.Println(result2)

}

func main() {
	day1()
}
