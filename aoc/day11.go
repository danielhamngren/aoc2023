package aoc

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Universe [][]string

func parseInputDay11(input []string) Universe {
	u := make(Universe, len(input))

	for row, line := range input {
		u[row] = make([]string, len(line))
		for col, letter := range line {
			u[row][col] = string(letter)
		}
	}

	return u
}

func expandUniverse(u Universe) Universe {
	rows := len(u)
	cols := len(u[0])

	row := 0
	col := 0

	// Expand rows
	for row < rows {
		allEmpty := true
		for col < cols {
			if u[row][col] != "." {
				allEmpty = false
			}
			col++
		}

		if allEmpty {
			dst := make([]string, cols)
			for i := range dst {
				dst[i] = "."
			}
			u = slices.Insert(u, row, dst)
			row++
			rows++
		}
		row++
		col = 0
	}

	// Expand cols
	rows = len(u)
	cols = len(u[0])

	row = 0
	col = 0

	for col < cols {
		allEmpty := true
		for row < rows {
			if u[row][col] != "." {
				allEmpty = false
			}
			row++
		}

		if allEmpty {
			for irow := range u {
				u[irow] = slices.Insert(u[irow], col, ".")
			}
			col++
			cols++
		}
		col++
		row = 0
	}

	return u

}

func displayUniverse(u Universe) {
	for row := range u {
		for col := range u[row] {
			fmt.Print(u[row][col])
		}

		fmt.Print("\n")
	}

	fmt.Print("\n")
}

func getGalaxyPositions(u Universe) []Position {
	ps := []Position{}
	for row := range u {
		for col := range u[row] {
			if u[row][col] == "#" {
				ps = append(ps, Position{row, col})
			}
		}
	}

	return ps
}

type GalaxyPair []Position

func getGalaxyPairs(positions []Position) []GalaxyPair {
	pairs := []GalaxyPair{}

	for i := 0; i < len(positions); i++ {
		for j := 0; j < i; j++ {
			pairs = append(pairs, []Position{positions[i], positions[j]})

		}
	}
	return pairs

}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pairDistance(pair GalaxyPair) int {
	a := pair[0]
	b := pair[1]

	return intAbs(a.row-b.row) + intAbs(a.col-b.col)

}

func day11part1(input []string) int {
	universe := parseInputDay11(input)

	displayUniverse(universe)
	universe = expandUniverse(universe)
	displayUniverse(universe)

	galaxyPositions := getGalaxyPositions(universe)
	fmt.Println(galaxyPositions)
	galaxyPairs := getGalaxyPairs(galaxyPositions)
	fmt.Println(len(galaxyPairs))

	dist := 0
	for _, pair := range galaxyPairs {
		dist += pairDistance(pair)
	}

	return dist
}

func day11part2(input []string) int {
	universe := parseInputDay11(input)

	fmt.Println(universe)
	return 0
}

func inputDay11(test int) []string {
	if test == 1 {
		return []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		}
	}
	content, err := os.ReadFile("aoc/input/day11")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day11() {
	resultPart1 := day11part1(inputDay11(0))
	fmt.Println("part1:", resultPart1)
	// resultPart2 := day11part2(inputDay10(0))
	// fmt.Println("part2:", resultPart2)

}
