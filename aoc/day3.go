package aoc

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	row int
	col int
}

type Label struct {
	value             int
	adjacentPositions []Position
}

type Symbol struct {
	value    string
	position Position
}

func get_labels_and_symbols(input []string) ([]Label, []Symbol) {
	// Get labels and symbols
	reNumber := regexp.MustCompile(`\d+`)
	reSymbol := regexp.MustCompile(`[^\d\.]`)

	labels := []Label{}
	symbols := []Symbol{}

	for row, element := range input {
		matchesNumber := reNumber.FindAllStringIndex(element, -1)
		matchesSymbol := reSymbol.FindAllStringIndex(element, -1)

		// add matches to labels
		for _, match := range matchesNumber {
			col_start := match[0]
			col_end := match[1]
			value, _ := strconv.Atoi(element[col_start:col_end])

			adjacentPositions := []Position{}
			for r := row - 1; r <= row+1; r++ {
				for c := col_start - 1; c <= col_end; c++ {
					adjacentPositions = append(adjacentPositions, Position{r, c})
				}
			}

			labels = append(labels, Label{value, adjacentPositions})
		}

		for _, match := range matchesSymbol {
			col_start := match[0]
			col_end := match[1]
			value := element[col_start:col_end]
			symbols = append(symbols, Symbol{value, Position{row, col_start}})
		}

	}

	return labels, symbols

}

func part1(input []string) int {

	labels, symbols := get_labels_and_symbols(input)

	sum := 0

	for _, symbol := range symbols {
	label_check:
		for _, label := range labels {
			for _, adjacentPosition := range label.adjacentPositions {
				if reflect.DeepEqual(symbol.position, adjacentPosition) {
					sum += label.value
					continue label_check
				}
			}
		}
	}
	return sum
}

func part2(input []string) int {
	labels, symbols := get_labels_and_symbols(input)

	sum := 0

	for _, symbol := range symbols {
		if symbol.value != "*" {
			continue
		}

		ratios := []int{}

	label_check:
		for _, label := range labels {
			for _, adjacentPosition := range label.adjacentPositions {
				if reflect.DeepEqual(symbol.position, adjacentPosition) {
					ratios = append(ratios, label.value)
					continue label_check
				}
			}
		}

		if len(ratios) == 2 {
			sum += ratios[0] * ratios[1]
		}
	}
	return sum

}

func test_input() []string {
	test_input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	return test_input
}

func real_input() []string {
	//Read input from file
	content, err := os.ReadFile("aoc/input/day3")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input
}

func Day3() {

	resultPart1 := part1(real_input())
	fmt.Println(resultPart1)
	resultPart2 := part2(real_input())
	fmt.Println(resultPart2)

}
