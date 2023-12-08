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
		fmt.Println(matches, item)
		node := matches[re.SubexpIndex("node")]
		left := matches[re.SubexpIndex("left")]
		right := matches[re.SubexpIndex("right")]

		fmt.Println(node, left, right)
		nodes[node] = Node{left, right}

	}

	return pattern, nodes

}

func day8part1(input []string) int {
	pattern, nodes := parseInputDay8(input)
	fmt.Println(pattern, nodes)

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

func day8part2(input []string) int {
	pattern, nodes := parseInputDay8(input)
	fmt.Println(pattern, nodes)

	//1.  Find all start nodes
	currentNodes := []string{}
	startingNodes := []string{}

	for nodeName, _ := range nodes {
		if nodeName[2] == 'A' {
			currentNodes = append(currentNodes, nodeName)
			startingNodes = append(startingNodes, nodeName)
		}
	}
	fmt.Println("starting nodes", currentNodes)
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
			// fmt.Println(direction, currentNode)
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

			// fmt.Println("cur", currentNodes[i], "start", startingNodes[i])

			// if counter > 10 {
			// 	break
			// }
		}
		fmt.Println("visited", visitedNodes[i])

	}
	fmt.Println(periods)
	fmt.Println(offsets)
	values := make([]int, len(offsets))

	result := 1
	for i, offset := range offsets {
		periods[i] = offset[1] - offset[0]
		values[i] = offset[0]
	}

	for {
		// find minimum value, add period to that value
		min := values[0]
		var ind int
		for i, val := range values {
			if val < min {
				ind = i
				min = val
			}
		}
		values[ind] += periods[ind]

		allEqual := true
		for i, _ := range values {
			if values[i] != values[0] {
				allEqual = false
				break
			}
		}

		fmt.Println(values)
		if allEqual {
			return 0
		}

	}
	fmt.Println(values)

	return result
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
