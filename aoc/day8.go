package aoc

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Node struct {
	left  string
	right string
}

func parseInputDay8(input []string) (string, map[string]Node) {

	pattern := input[0]

	nodes := make(map[string]Node)
	re := regexp.MustCompile(`(?P<node>\w+) = \((?P<left>\w+), (?P<right>\w+)\)`)

	for _, item := range input[2:] {
		matches := re.FindStringSubmatch(item)
		node := matches[re.SubexpIndex("node")]
		left := matches[re.SubexpIndex("left")]
		right := matches[re.SubexpIndex("right")]

		nodes[node] = Node{left, right}

	}

	return pattern, nodes

}

func day8part1(input []string) int {
	pattern, nodes := parseInputDay8(input)

	counter := 0
	currentNode := "AAA"

	for {
		direction := pattern[counter%len(pattern)]
		fmt.Println(direction, currentNode)
		if direction == 'L' {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}

		counter++
		if currentNode == "ZZZ" {
			break
		}

	}

	return counter
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(input []int) int {
	m := input[0]
	for _, element := range input[1:] {
		m = abs(m*element) / gcd(m, element)
	}
	return m

}

func day8part2(input []string) int {
	pattern, nodes := parseInputDay8(input)

	//1.  Find all start nodes
	currentNodes := []string{}
	startingNodes := []string{}

	for nodeName, _ := range nodes {
		if nodeName[2] == 'A' {
			currentNodes = append(currentNodes, nodeName)
			startingNodes = append(startingNodes, nodeName)
		}
	}
	periods := make([]int, len(startingNodes))
	offsets := make([][]int, len(startingNodes))
	visitedNodes := make([][]string, len(startingNodes))
	for i := range offsets {
		offsets[i] = []int{}
	}
	for i := range visitedNodes {
		visitedNodes[i] = []string{startingNodes[i]}
	}

	// for each starting node, find the period
	for i, _ := range currentNodes {
		counter := 0
		for {
			direction := pattern[counter%len(pattern)]

			// for all current Nodes, step forward
			if direction == 'L' {
				currentNodes[i] = nodes[currentNodes[i]].left
			} else {
				currentNodes[i] = nodes[currentNodes[i]].right
			}
			counter++

			if slices.Contains(visitedNodes[i], currentNodes[i]) && counter > len(pattern)*len(nodes) {

				break
			} else {
				visitedNodes[i] = append(visitedNodes[i], currentNodes[i])
			}

			if currentNodes[i][2] == 'Z' {
				offsets[i] = append(offsets[i], counter)
			}

		}
	}
	values := make([]int, len(offsets))

	for i, offset := range offsets {
		periods[i] = offset[1] - offset[0]
		values[i] = offset[0]
	}
	fmt.Println(periods)

	return lcm(periods)

}

func inputDay8(test bool) []string {
	if test {
		// return []string{
		// 	"RL",
		// 	"",
		// 	"AAA = (BBB, CCC)",
		// 	"BBB = (DDD, EEE)",
		// 	"CCC = (ZZZ, GGG)",
		// 	"DDD = (DDD, DDD)",
		// 	"EEE = (EEE, EEE)",
		// 	"GGG = (GGG, GGG)",
		// 	"ZZZ = (ZZZ, ZZZ)",
		// }
		// return []string{
		// 	"LLR",
		// 	"",
		// 	"AAA = (BBB, BBB)",
		// 	"BBB = (AAA, ZZZ)",
		// 	"ZZZ = (ZZZ, ZZZ)",
		// }
		return []string{
			"LR",
			"",
			"11A = (11B, XXX)",
			"11B = (XXX, 11Z)",
			"11Z = (11B, XXX)",
			"22A = (22B, XXX)",
			"22B = (22C, 22C)",
			"22C = (22Z, 22Z)",
			"22Z = (22B, 22B)",
			"XXX = (XXX, XXX)",
		}

	}

	content, err := os.ReadFile("aoc/input/day8")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day8() {
	// resultPart1 := day8part1(inputDay8(false))
	// fmt.Println(resultPart1)
	resultPart2 := day8part2(inputDay8(false))
	fmt.Println(resultPart2)

	// 30937253904925038111814859 is too high
}
