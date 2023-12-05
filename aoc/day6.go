package aoc

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInputDay6(input []string) {
}

func day6part1(input []string) int {
	parseInputDay6(input)

	result := 0

	return result
}

func day6part2(input []string) int {
	parseInputDay6(input)

	result := 0

	return result
}

func inputDay6(test bool) []string {
	if test {
		return []string{}

	}

	content, err := os.ReadFile("aoc/input/day6")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day6() {

	resultPart1 := day6part1(inputDay5(false))
	fmt.Println(resultPart1)
	resultPart2 := day6part2(inputDay5(false))
	fmt.Println(resultPart2)

}
