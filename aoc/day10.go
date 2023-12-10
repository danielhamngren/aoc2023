package aoc

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Pipe struct {
	connections []Position
	pipeType    string
}

func parseInputDay10(input []string) ([][]Pipe, Position) {
	rows := len(input)
	cols := len(input[0])

	pipes := make([][]Pipe, rows)
	for i := 0; i < rows; i++ {
		pipes[i] = make([]Pipe, cols)
	}

	var startPos Position
	for row := range input {
		for col := range input[row] {
			fmt.Printf("%c", input[row][col])

			connections := []Position{}

			pipeType := string(input[row][col])

			switch pipeType {
			case "|":
				connections = append(connections, Position{row - 1, col})
				connections = append(connections, Position{row + 1, col})
			case "-":
				connections = append(connections, Position{row, col - 1})
				connections = append(connections, Position{row, col + 1})
			case "L":
				connections = append(connections, Position{row - 1, col})
				connections = append(connections, Position{row, col + 1})
			case "J":
				connections = append(connections, Position{row - 1, col})
				connections = append(connections, Position{row, col - 1})
			case "7":
				connections = append(connections, Position{row, col - 1})
				connections = append(connections, Position{row + 1, col})
			case "F":
				connections = append(connections, Position{row, col + 1})
				connections = append(connections, Position{row + 1, col})
			case "S":
				startPos.row = row
				startPos.col = col
			}

			pipes[row][col] = Pipe{connections, pipeType}

		}

		fmt.Println("")
	}

	// find connections for start pipe

	if startPos.row > 0 && strings.Contains("|F7", pipes[startPos.row-1][startPos.col].pipeType) {
		pipes[startPos.row][startPos.col].connections = append(
			pipes[startPos.row][startPos.col].connections,
			Position{startPos.row - 1, startPos.col})
	}
	if startPos.col < cols-1 && strings.Contains("|JL", pipes[startPos.row+1][startPos.col].pipeType) {
		pipes[startPos.row][startPos.col].connections = append(
			pipes[startPos.row][startPos.col].connections,
			Position{startPos.row + 1, startPos.col})
	}
	if startPos.col > 0 && strings.Contains("-FL", pipes[startPos.row][startPos.col-1].pipeType) {
		pipes[startPos.row][startPos.col].connections = append(
			pipes[startPos.row][startPos.col].connections,
			Position{startPos.row, startPos.col - 1})
	}
	if startPos.col < cols-1 && strings.Contains("-7J", pipes[startPos.row][startPos.col+1].pipeType) {
		pipes[startPos.row][startPos.col].connections = append(
			pipes[startPos.row][startPos.col].connections,
			Position{startPos.row, startPos.col + 1})
	}

	return pipes, startPos
}

func day10part1(input []string) int {
	pipes, startPos := parseInputDay10(input)

	fmt.Println(startPos)
	fmt.Println(pipes)

	visited := []Position{startPos}

	dir1 := pipes[startPos.row][startPos.col].connections[0]
	dir2 := pipes[startPos.row][startPos.col].connections[1]
	steps := 1

	for {
		// fmt.Println(dir1, pipes[dir1.row][dir1.col])
		// fmt.Println(dir2, pipes[dir2.row][dir2.col])
		if dir1.col == dir2.col && dir1.row == dir2.row {
			break
		}

		for _, connection := range pipes[dir1.row][dir1.col].connections {
			if !slices.Contains(visited, connection) {
				visited = append(visited, dir1)
				dir1 = connection
			}
		}
		for _, connection := range pipes[dir2.row][dir2.col].connections {
			if !slices.Contains(visited, connection) {
				visited = append(visited, dir2)
				dir2 = connection
			}
		}
		steps++
		// fmt.Println("visited", visited)
		fmt.Println("steps", steps)
		// if steps > 10 {
		// 	break
		// }
	}

	return steps
}

func day10part2(input []string) int {
	pipes, startPos := parseInputDay10(input)

	fmt.Println(startPos)
	fmt.Println(pipes)

	visited := []Position{startPos}

	current := pipes[startPos.row][startPos.col].connections[0]
	steps := 1

	for {
		// fmt.Println("current", current, "connections", pipes[current.row][current.col].connections, "visited", visited)
		visited = append(visited, current)
		for _, connection := range pipes[current.row][current.col].connections {
			// fmt.Println(connection)
			if !slices.Contains(visited, connection) {
				current = connection
			}
		}

		// fmt.Println(visited)
		fmt.Println(current)
		if slices.Contains(visited, current) {
			break
		}

		steps++

		fmt.Println("steps", steps)
	}

	fmt.Println(visited)

	// make all unvisited into dots.

	for row := range pipes {
		for col := range pipes[row] {
			if !slices.Contains(visited, Position{row, col}) {
				pipes[row][col] = Pipe{[]Position{}, "."}

			}
			fmt.Print(pipes[row][col].pipeType)
		}
		fmt.Print("\n")
	}

	// Traverse all the visited pipes in order. Mark left with O and right with I if pipe type there is "."
	// I think I need a concept of direction from where you go. UP, DOWN, LEFT, RIGHT.

	// if you go UP, left is (0, -1)
	// if you go UP, left is (0, 1)
	// if you go RIGHT, left is (-1, 0)
	// if you go LEFT, left is (1, 0)

	// This might only be relevant if the type is "- or |", no that is not true.

	for i, current := range visited {
		var direction string

		var previous Position
		if i == 0 {
			previous = visited[len(visited)-1]
		} else {
			previous = visited[i-1]
		}

		if current.row > previous.row {
			direction = "DOWN"
		}
		if current.row < previous.row {
			direction = "UP"
		}
		if current.col > previous.col {
			direction = "RIGHT"
		}
		if current.col < previous.col {
			direction = "LEFT"
		}
		fmt.Println(i, current, direction)

		var left Position
		var right Position
		var front Position
		var turningRight bool
		if direction == "DOWN" {
			left = Position{current.row, current.col + 1}
			right = Position{current.row, current.col - 1}
			front = Position{current.row + 1, current.col}
			turningRight = pipes[current.row][current.col].pipeType == "J"
		}
		if direction == "UP" {
			left = Position{current.row, current.col - 1}
			right = Position{current.row, current.col + 1}
			front = Position{current.row - 1, current.col}
			turningRight = pipes[current.row][current.col].pipeType == "F"
		}
		if direction == "LEFT" {
			left = Position{current.row + 1, current.col}
			right = Position{current.row - 1, current.col}
			front = Position{current.row, current.col - 1}
			turningRight = pipes[current.row][current.col].pipeType == "L"
		}
		if direction == "RIGHT" {
			left = Position{current.row - 1, current.col}
			right = Position{current.row + 1, current.col}
			front = Position{current.row, current.col + 1}
			turningRight = pipes[current.row][current.col].pipeType == "7"
		}

		// This is relevant for all the pipe types
		fmt.Println(left.row, left.col, right.row, right.col)
		if left.row >= 0 && left.row < len(pipes) && left.col >= 0 && left.col < len(pipes[0]) {
			if pipes[left.row][left.col].pipeType == "." {
				pipes[left.row][left.col].pipeType = "I"
			}
		}
		if right.row >= 0 && right.row < len(pipes) && right.col >= 0 && right.col < len(pipes[0]) {
			if pipes[right.row][right.col].pipeType == "." {
				pipes[right.row][right.col].pipeType = "O"
			}
		}

		// This is only relevant for bends
		// Set the pipe in front to inner or outer.
		if front.row >= 0 && front.row < len(pipes) && front.col >= 0 && front.col < len(pipes[0]) {

			if pipes[front.row][front.col].pipeType == "." {
				if turningRight {
					pipes[front.row][front.col].pipeType = "I"
				} else {
					pipes[front.row][front.col].pipeType = "O"
				}
			}
		}

	}

	printPipes(pipes)

	// Fill neighbours
	for numberOfDots(pipes) > 0 {
		for row := range pipes {
			for col := range pipes[row] {
				if pipes[row][col].pipeType == "." {
					if row > 0 && pipes[row-1][col].pipeType == "I" {
						pipes[row][col].pipeType = "I"
					} else if row < len(pipes)-1 && pipes[row+1][col].pipeType == "I" {
						pipes[row][col].pipeType = "I"
					} else if col > 0 && pipes[row][col-1].pipeType == "I" {
						pipes[row][col].pipeType = "I"
					} else if col < len(pipes[0])-1 && pipes[row][col+1].pipeType == "I" {
						pipes[row][col].pipeType = "I"
					}

					if row > 0 && pipes[row-1][col].pipeType == "O" {
						pipes[row][col].pipeType = "O"
					} else if row < len(pipes)-1 && pipes[row+1][col].pipeType == "O" {
						pipes[row][col].pipeType = "O"
					} else if col > 0 && pipes[row][col-1].pipeType == "O" {
						pipes[row][col].pipeType = "O"
					} else if col < len(pipes[0])-1 && pipes[row][col+1].pipeType == "O" {
						pipes[row][col].pipeType = "O"
					}

				}
			}
		}
	}
	fmt.Print("\n")
	printPipes(pipes)
	inside, outside := countSides(pipes)
	fmt.Println("inside:", inside, "outside:", outside)

	return inside
}
func printPipes(pipes [][]Pipe) {
	for row := range pipes {
		for col := range pipes[row] {
			fmt.Print(pipes[row][col].pipeType)
		}
		fmt.Print("\n")
	}
}

func numberOfDots(pipes [][]Pipe) int {
	dots := 0

	for row := range pipes {
		for col := range pipes[row] {
			if pipes[row][col].pipeType == "." {
				dots++
			}
		}
	}
	return dots
}

func countSides(pipes [][]Pipe) (int, int) {
	inside := 0
	outside := 0

	for row := range pipes {
		for col := range pipes[row] {
			if pipes[row][col].pipeType == "I" {
				inside++
			}
			if pipes[row][col].pipeType == "O" {
				outside++
			}
		}
	}
	return inside, outside
}

func inputDay10(test int) []string {
	if test == 1 {
		return []string{
			"-L|F7",
			"7S-7|",
			"L|7||",
			"-L-J|",
			"L|-JF",
		}
	}
	if test == 2 {
		return []string{
			"7-F7-",
			".FJ|7",
			"SJLL7",
			"|F--J",
			"LJ.LJ",
		}
	}
	if test == 3 {
		return []string{
			"...........",
			".S-------7.",
			".|F-----7|.",
			".||.....||.",
			".||.....||.",
			".|L-7.F-J|.",
			".|..|.|..|.",
			".L--J.L--J.",
			"...........",
		}
	}
	if test == 32 {
		return []string{
			"...........",
			"...........",
			"...........",
			".S-------7.",
			".|F-----7|.",
			".||.....||.",
			".||.....||.",
			".||.....||.",
			".|L-7.F-J|.",
			".|..|.|..|.",
			".L--J.L--J.",
			"...........",
		}
	}
	if test == 4 {
		return []string{
			"..........",
			".S------7.",
			".|F----7|.",
			".||OOOO||.",
			".||OOOO||.",
			".|L-7F-J|.",
			".|II||II|.",
			".L--JL--J.",
			"..........",
		}
	}
	if test == 5 {
		return []string{
			".F----7F7F7F7F-7....",
			".|F--7||||||||FJ....",
			".||.FJ||||||||L7....",
			"FJL7L7LJLJ||LJ.L-7..",
			"L--J.L7...LJS7F-7L7.",
			"....F-J..F7FJ|L7L7L7",
			"....L7.F7||L7|.L7L7|",
			".....|FJLJ|FJ|F7|.LJ",
			"....FJL-7.||.||||...",
			"....L---J.LJ.LJLJ...",
		}
	}
	if test == 6 {
		return []string{
			"FF7FSF7F7F7F7F7F---7",
			"L|LJ||||||||||||F--J",
			"FL-7LJLJ||||||LJL-77",
			"F--JF--7||LJLJ7F7FJ-",
			"L---JF-JLJ.||-FJLJJ7",
			"|F|F-JF---7F7-L7L|7|",
			"|FFJF7L7F-JF7|JL---7",
			"7-L-JL7||F7|L7F-7F7|",
			"L.L7LFJ|||||FJL7||LJ",
			"L7JLJL-JLJLJL--JLJ.L",
		}
	}
	content, err := os.ReadFile("aoc/input/day10")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day10() {
	// resultPart1 := day10part1(inputDay10(0))
	// fmt.Println("part1:", resultPart1)
	resultPart2 := day10part2(inputDay10(0))
	fmt.Println("part2:", resultPart2)

}
