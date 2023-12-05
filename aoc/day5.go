package aoc

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type FunctionSection struct {
	startSource      int
	startDestination int
	sectionRange     int
}

type Function struct {
	sections []FunctionSection
}

func parseInputDay5(input []string) ([]int, map[string]Function) {
	m := make(map[string]Function)
	seeds := []int{}
	var seedsRaw string

	seedsRaw, input = input[0], input[2:]
	re := regexp.MustCompile(`(?P<destination>\d+) (?P<source>\d+) (?P<range>\d+)`)

	for _, element := range strings.Split(seedsRaw, " ")[1:] {
		seed, _ := strconv.Atoi(element)
		seeds = append(seeds, seed)
	}

	var currentFunction Function
	var currentFunctionName string
	for _, line := range input {
		fmt.Println(line)
		if strings.Contains(line, "map") {
			currentFunctionName = line
			currentFunction = Function{[]FunctionSection{}}
			continue
		} else if line == "" {
			fmt.Println("add function to map")
			m[currentFunctionName] = currentFunction
			continue
		}

		matches := re.FindStringSubmatch(line)
		fmt.Println(matches, line)
		source, _ := strconv.Atoi(matches[re.SubexpIndex("source")])
		destination, _ := strconv.Atoi(matches[re.SubexpIndex("destination")])
		sectionRange, _ := strconv.Atoi(matches[re.SubexpIndex("range")])

		currentFunction.sections = append(currentFunction.sections, FunctionSection{source, destination, sectionRange})

	}

	return seeds, m
}

func day5part1(input []string) int {
	seeds, m := parseInputDay5(input)

	fmt.Println(seeds, m)

	//TODO: find which seed has the lowest location.
	//TODO: process the seed numbers using the output of the parse function

	result := 0

	return result
}

func day5part2(input []string) int {
	result := 0

	return result
}

func inputDay5(test bool) []string {
	if test {
		return []string{
			"seeds: 79 14 55 13",
			"",
			"seed-to-soil map:",
			"50 98 2",
			"52 50 48",
			"",
			"soil-to-fertilizer map:",
			"0 15 37",
			"37 52 2",
			"39 0 15",
			"",
			"fertilizer-to-water map:",
			"49 53 8",
			"0 11 42",
			"42 0 7",
			"57 7 4",
			"",
			"water-to-light map:",
			"88 18 7",
			"18 25 70",
			"",
			"light-to-temperature map:",
			"45 77 23",
			"81 45 19",
			"68 64 13",
			"",
			"temperature-to-humidity map:",
			"0 69 1",
			"1 0 69",
			"",
			"humidity-to-location map:",
			"60 56 37",
			"56 93 4",
		}
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
