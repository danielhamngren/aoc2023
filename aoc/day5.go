package aoc

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func day5part1(input []string) int {
	result := 0

	return result
}

func day5part2(input []string) int {
	result := 0

	return result
}

func inputDay5(test bool) []string {
	if test {
		return []string{}
	}

	content, err := os.ReadFile("aoc/input/day5")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day5() {

	resultPart1 := day5part1(inputDay5(true))
	fmt.Println(resultPart1)
	resultPart2 := day5part2(inputDay5(true))
	fmt.Println(resultPart2)

}
