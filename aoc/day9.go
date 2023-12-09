package aoc

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type History struct {
	measurements []int
}

func parseInputDay9(input []string) []History {
	output := []History{}
	re := regexp.MustCompile(`-?\d+`)

	for _, line := range input {
		m := []int{}
		matches := re.FindAllString(line, -1)
		for _, matchStr := range matches {
			matchInt, _ := strconv.Atoi(matchStr)
			m = append(m, matchInt)
		}

		output = append(output, History{m})

	}

	return output
}

func checkAllZero(s []int) bool {
	for _, e := range s {
		if e != 0 {
			return false
		}
	}
	return true
}

func getSliceDifference(s []int) []int {
	result := make([]int, len(s)-1)
	for i := range s[1:] {
		result[i] = s[i+1] - s[i]
	}
	return result
}

func predictNext(m []int) int {
	fmt.Println(m)
	diffs := getSliceDifference(m)

	if checkAllZero(m) {
		return 0
	} else if len(m) == 1 {
		fmt.Println("-- LENGHT IS ONE --, undefined behaviour!!!")
		// return m[0]
		return 0
	} else {
		return predictNext(diffs) + m[len(m)-1]
	}

}

func day9part1(input []string) int {
	histories := parseInputDay9(input)

	result := 0

	for _, history := range histories {
		next := predictNext(history.measurements)
		fmt.Println("next", next)
		result += next
	}

	return result
}

func day9part2(input []string) int {
	history := parseInputDay9(input)

	fmt.Println(history)

	return 0

}

func inputDay9(test bool) []string {
	if test {
		return []string{"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
			"0 1 1 0",
		}
	}

	content, err := os.ReadFile("aoc/input/day9")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day9() {
	resultPart1 := day9part1(inputDay9(false))
	fmt.Println(resultPart1)
	// resultPart2 := day9part2(inputDay9(false))
	// fmt.Println(resultPart2)

}
