package aoc

import (
	"fmt"
	"log"
	"math"
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
	for i, line := range input {
		fmt.Println(line)
		if strings.Contains(line, "map") {
			currentFunctionName = strings.Split(line, " ")[0]
			currentFunction = Function{[]FunctionSection{}}
			continue
		} else if line == "" || i == len(input)-1 {
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

func executeFunction(f Function, x int) int {
	for _, section := range f.sections {
		fmt.Println(section.startDestination, section.startSource, section.sectionRange, x)

		if x >= section.startSource && x <= section.startSource+section.sectionRange {
			fmt.Println("section valid")
			return section.startDestination + x - section.startSource
		}
	}

	return x
}

func executeMaps(functions map[string]Function, x int) int {
	fmt.Println("seed:", x)
	x = executeFunction(functions["seed-to-soil"], x)
	fmt.Println("soil:", x)

	x = executeFunction(functions["soil-to-fertilizer"], x)
	fmt.Println("fert:", x)
	x = executeFunction(functions["fertilizer-to-water"], x)
	fmt.Println("water:", x)
	x = executeFunction(functions["water-to-light"], x)
	fmt.Println("light:", x)
	x = executeFunction(functions["light-to-temperature"], x)
	fmt.Println("temp:", x)
	x = executeFunction(functions["temperature-to-humidity"], x)
	fmt.Println("humidity:", x)
	x = executeFunction(functions["humidity-to-location"], x)
	fmt.Println("location:", x)

	return x

}

func day5part1(input []string) int {
	seeds, m := parseInputDay5(input)

	fmt.Println(seeds, m)

	//TODO: find which seed has the lowest location.
	//TODO: process the seed numbers using the output of the parse function
	min_location := math.MaxInt64
	min_seed := seeds[0]

	for _, seed := range seeds {
		location := executeMaps(m, seed)
		if location < min_location {
			min_seed = seed
			min_location = location
		}
	}

	fmt.Println(min_seed, min_location)

	result := min_location

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

	resultPart1 := day5part1(inputDay5(false))
	fmt.Println(resultPart1)
	resultPart2 := day5part2(inputDay5(true))
	fmt.Println(resultPart2)

}
