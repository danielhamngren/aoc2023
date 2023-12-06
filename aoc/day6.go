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

func parseInputDay6(input []string) ([]float64, []float64) {
	re := regexp.MustCompile(`\d+`)

	ts := []float64{}
	ds := []float64{}

	for _, t := range re.FindAllString(input[0], -1) {
		res, _ := strconv.ParseFloat(t, 64)
		ts = append(ts, res)
	}

	for _, d := range re.FindAllString(input[1], -1) {
		res, _ := strconv.ParseFloat(d, 64)
		ds = append(ds, res)
	}

	return ts, ds
}

func parseInputDay6Part2(input []string) (float64, float64) {
	re := regexp.MustCompile(`\d+`)

	t, _ := strconv.ParseFloat(strings.Join(re.FindAllString(input[0], -1), ""), 64)
	d, _ := strconv.ParseFloat(strings.Join(re.FindAllString(input[1], -1), ""), 64)

	return t, d
}

func day6part1(input []string) float64 {
	ts, ds := parseInputDay6(input)
	result := 1.0

	for i := range ts {
		t := ts[i]
		d := ds[i]

		diffFromMax := math.Sqrt(t*t-4*d) / 2
		lowerLimit := math.Ceil(t/2 - diffFromMax)
		upperLimit := math.Floor(t/2 + diffFromMax)
		diff := upperLimit - lowerLimit

		modifier := 1.0
		if math.Abs(diffFromMax-math.Round(diffFromMax)) < 0.0000000001 {
			modifier = -1.0
		}
		fmt.Println(t, d, diffFromMax, lowerLimit, upperLimit, diff+modifier)

		possibleWins := diff + modifier

		result *= possibleWins

	}

	return result
}

func day6part2(input []string) float64 {
	t, d := parseInputDay6Part2(input)
	diffFromMax := math.Sqrt(t*t-4*d) / 2
	lowerLimit := math.Ceil(t/2 - diffFromMax)
	upperLimit := math.Floor(t/2 + diffFromMax)
	diff := upperLimit - lowerLimit

	modifier := 1.0
	if math.Abs(diffFromMax-math.Round(diffFromMax)) < 0.0000000001 {
		modifier = -1.0
	}
	fmt.Println(t, d, diffFromMax, lowerLimit, upperLimit, diff+modifier)

	result := diff + modifier

	return result
}

func inputDay6(test bool) []string {
	if test {
		return []string{
			"Time:      7  15   30",
			"Distance:  9  40  200",
		}

	}

	content, err := os.ReadFile("aoc/input/day6")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day6() {

	resultPart1 := day6part1(inputDay6(false))
	fmt.Println(resultPart1)
	resultPart2 := day6part2(inputDay6(false))
	fmt.Printf("%f\n", resultPart2)

}
