package aoc

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type ScratchCard struct {
	winningNumbers []int
	currentNumbers []int
	copies         int
}

func checkWinning(sc ScratchCard) int {
	matches := 0
	for _, current := range sc.currentNumbers {
		for _, winning := range sc.winningNumbers {
			if winning == current {
				matches++
			}
		}
	}

	if matches == 0 {
		return 0
	}
	return int(math.Pow(2, float64(matches-1)))
}

func checkWinningPart2(sc ScratchCard) int {
	winnings := 0
	for _, current := range sc.currentNumbers {
		for _, winning := range sc.winningNumbers {
			if winning == current {
				winnings++
			}
		}
	}

	return winnings
}

func parseInputDay4(input []string) []ScratchCard {
	// Parse input
	scratchcards := []ScratchCard{}

	for _, row := range input {
		content := strings.Split(row, ":")
		numbersRaw := strings.Split(content[1], "|")

		winningNumbersRaw := strings.Split(strings.TrimSpace(numbersRaw[0]), " ")
		currentNumbersRaw := strings.Split(strings.TrimSpace(numbersRaw[1]), " ")

		winningNumbers := []int{}
		currentNumbers := []int{}

		for _, number := range winningNumbersRaw {
			if number == "" {
				continue
			}
			x, _ := strconv.Atoi(number)
			winningNumbers = append(winningNumbers, x)
		}
		for _, number := range currentNumbersRaw {
			if number == "" {
				continue
			}
			x, _ := strconv.Atoi(number)
			currentNumbers = append(currentNumbers, x)
		}

		scratchcards = append(scratchcards, ScratchCard{winningNumbers, currentNumbers, 1})
	}

	return scratchcards

}

func day4part1(input []string) int {
	scratchcards := parseInputDay4(input)
	result := 0

	for _, sc := range scratchcards {
		winning := checkWinning(sc)
		result += winning
	}

	return result
}

func day4part2(input []string) int {
	scratchcards := parseInputDay4(input)
	result := 0

	for i, sc := range scratchcards {
		result += sc.copies
		winnings := checkWinningPart2(sc)
		for j := 0; j < winnings; j++ {
			scratchcards[i+1+j].copies += sc.copies
		}
	}

	return result
}

func input(test bool) []string {
	if test {
		return []string{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
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

	resultPart1 := day4part1(input(false))
	fmt.Println(resultPart1)
	resultPart2 := day4part2(input(false))
	fmt.Println(resultPart2)

}
