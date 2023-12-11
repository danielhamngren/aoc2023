package aoc

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Universe []Galaxy
type Galaxy Position

func parseInputDay11(input []string) Universe {
	u := Universe{}

	for row, line := range input {
		for col, letter := range line {
			if letter == '#' {
				u = append(u, Galaxy{row, col})
			}
		}
	}

	return u
}

func expandUniverse(u Universe, expansion int) Universe {

	minRow := u[0].row
	maxRow := u[0].row
	minCol := u[0].col
	maxCol := u[0].col

	for _, galaxy := range u {
		if galaxy.row > maxRow {
			maxRow = galaxy.row
		}
		if galaxy.row < minRow {
			minRow = galaxy.row
		}
		if galaxy.col < minCol {
			minCol = galaxy.col
		}
		if galaxy.col > maxCol {
			maxCol = galaxy.col
		}
	}

	// check rows
	for row := minRow; row < maxRow; {
		foundGalaxy := false
		fmt.Println("row", row)
		for _, galaxy := range u {
			if galaxy.row == row {
				foundGalaxy = true
			}
		}
		if !foundGalaxy {
			// expand for row
			fmt.Println("Ready to expand")
			maxRow += expansion - 1
			for i, galaxy := range u {
				if galaxy.row > row {
					fmt.Println("galaxy moved", galaxy)
					galaxy.row += expansion - 1
					fmt.Println("to", galaxy)
					u[i] = galaxy
				}
			}
			row += expansion - 1

		}
		row++
	}

	// check cols
	for col := minCol; col < maxCol; {
		foundGalaxy := false
		fmt.Println("col", col)
		for _, galaxy := range u {
			if galaxy.col == col {
				foundGalaxy = true
			}
		}
		if !foundGalaxy {
			// expand for col
			fmt.Println("Ready to expand")
			maxCol += expansion - 1
			for i, galaxy := range u {
				if galaxy.col > col {
					fmt.Println("galaxy moved", galaxy)
					galaxy.col += expansion - 1
					fmt.Println("to", galaxy)
					u[i] = galaxy
				}
			}
			col += expansion - 1

		}
		col++
	}

	return u
}

func displayUniverse(u Universe) {
	fmt.Println(u)
}

type GalaxyPair []Galaxy

func getGalaxyPairs(u Universe) []GalaxyPair {
	pairs := []GalaxyPair{}

	for i := 0; i < len(u); i++ {
		for j := 0; j < i; j++ {
			pairs = append(pairs, GalaxyPair{u[i], u[j]})

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
	universe = expandUniverse(universe, 2)
	displayUniverse(universe)

	galaxyPairs := getGalaxyPairs(universe)
	fmt.Println("galaxies:", len(universe))
	fmt.Println("galaxy pairs:", len(galaxyPairs))

	dist := 0
	for _, pair := range galaxyPairs {
		dist += pairDistance(pair)
	}

	return dist
}

func day11part2(input []string) int {
	universe := parseInputDay11(input)

	displayUniverse(universe)
	universe = expandUniverse(universe, 1000000)
	displayUniverse(universe)

	galaxyPairs := getGalaxyPairs(universe)
	fmt.Println("galaxies:", len(universe))
	fmt.Println("galaxy pairs:", len(galaxyPairs))
	fmt.Println(galaxyPairs)

	dist := 0
	for _, pair := range galaxyPairs {
		d := pairDistance(pair)
		fmt.Println("pair", pair, "distance:", d)
		dist += d
	}

	return dist
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
	if test == 2 {
		return []string{
			"#.#",
			"...",
			"#.#",
		}
	}
	if test == 3 {
		return []string{
			"#.##.#",
		}
	}
	if test == 3 {
		return []string{
			"#",
			".",
			"#",
			"#",
			".",
			"#",
		}
	}
	if test == 4 {
		return []string{
			"#.#",
			"...",
			"...",
			"#.#",
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
	// resultPart1 := day11part1(inputDay11(0))
	// fmt.Println("part1:", resultPart1)
	resultPart2 := day11part2(inputDay11(0))
	fmt.Println("part2:", resultPart2)

}
