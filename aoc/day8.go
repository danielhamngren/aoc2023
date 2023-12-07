package aoc

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInputDay8(input []string) {
}

func day8part1(input []string) int {
	return 0
}

func day8part2(input []string) int {
	return 0
}

func inputDay8(test bool) []string {
	if test {
		return []string{}

	}

	content, err := os.ReadFile("aoc/input/day8")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day8() {
	resultPart1 := day8part1(inputDay8(true))
	fmt.Println(resultPart1)
	resultPart2 := day8part2(inputDay8(true))
	fmt.Println(resultPart2)
}
