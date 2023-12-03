package aoc

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func day4part1(input []string) int {
	result := 0

	return result
}

func day4part2(input []string) int {
	result := 0

	return result
}

func input(test bool) []string {
	if test {
		return []string{
			"",
		}
	}

	content, err := os.ReadFile("aoc/input/day4")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day4() {

	resultPart1 := day4part1(input(true))
	fmt.Println(resultPart1)
	resultPart2 := day4part2(input(true))
	fmt.Println(resultPart2)

}
