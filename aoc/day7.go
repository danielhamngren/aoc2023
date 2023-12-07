package aoc

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bet   int
}

func parseInputDay7(input []string) []Hand {
	hands := []Hand{}

	for _, row := range input {
		items := strings.Split(row, " ")
		cards := items[0]
		bet, _ := strconv.Atoi(items[1])
		hand := Hand{
			cards, bet,
		}
		hands = append(hands, hand)
	}

	return hands
}

func checkCards(cards string, jokers bool) int {
	cardsMap := make(map[rune]int)

	for _, c := range cards {
		val, ok := cardsMap[c]
		if ok {
			cardsMap[c] = val + 1
		} else {
			cardsMap[c] = 1
		}

	}

	jo := 0
	if jokers {
		var ok bool
		jo, ok = cardsMap['J']
		delete(cardsMap, 'J')
		if !ok {
			jo = 0
		}

	}

	values := []int{}
	for _, v := range cardsMap {
		values = append(values, v)

	}
	if len(values) == 0 {
		values = append(values, 0)
	}
	slices.SortFunc(values, func(i, j int) int { return j - i })
	// Check 5 of a kind 7
	fmt.Print(cardsMap, jo)
	values[0] = values[0] + jo

	if values[0] == 5 {
		return 7
	}

	// check 4 or a kind 6
	if values[0] == 4 {
		return 6
	}

	// check full house 5
	if values[0] == 3 && values[1] == 2 {
		return 5
	}

	// check three of a kind 4
	if values[0] == 3 {
		return 4
	}

	// check two pair 3
	if values[0] == 2 && values[1] == 2 {
		return 3
	}

	// check one pair 2
	if values[0] == 2 {
		return 2
	}

	// high card 1
	return 1
}

func firstWins(x, y Hand, jokers bool) int {
	var values map[byte]int
	if jokers {
		values = map[byte]int{
			'A': 14,
			'K': 13,
			'Q': 12,
			'J': 1,
			'T': 10,
			'9': 9,
			'8': 8,
			'7': 7,
			'6': 6,
			'5': 5,
			'4': 4,
			'3': 3,
			'2': 2,
		}
	} else {
		values = map[byte]int{
			'A': 14,
			'K': 13,
			'Q': 12,
			'J': 11,
			'T': 10,
			'9': 9,
			'8': 8,
			'7': 7,
			'6': 6,
			'5': 5,
			'4': 4,
			'3': 3,
			'2': 2,
		}

	}

	for i := range x.cards {
		if values[x.cards[i]] == values[y.cards[i]] {
			continue
		}
		return values[x.cards[i]] - values[y.cards[i]]
	}

	return 0

}

func checkWinningHand(x, y Hand, jokers bool) int {
	// 1. Get the hand type
	handOrderX := checkCards(x.cards, jokers)
	handOrderY := checkCards(y.cards, jokers)

	if handOrderX == handOrderY {
		// 3. If same, handle draw
		return firstWins(x, y, jokers)
	}
	// 2. Check which is better

	return handOrderX - handOrderY

}

func day7part1(input []string) int {
	hands := parseInputDay7(input)
	result := 0

	fmt.Println(hands)

	slices.SortFunc(hands, func(i, j Hand) int { return checkWinningHand(i, j, false) })
	fmt.Println(hands)

	// get total winnings

	for i, hand := range hands {
		result += hand.bet * (i + 1)
	}

	return result
}

func day7part2(input []string) int {
	hands := parseInputDay7(input)
	result := 0

	fmt.Println(hands)

	slices.SortFunc(hands, func(i, j Hand) int { return checkWinningHand(i, j, true) })
	fmt.Println(hands)

	// get total winnings

	for i, hand := range hands {
		result += hand.bet * (i + 1)
	}

	return result
}

func inputDay7(test bool) []string {
	if test {
		return []string{
			"32T3K 765",
			"T55J5 684",
			"KK677 28",
			"KTJJT 220",
			"QQQJA 483",
		}

	}

	content, err := os.ReadFile("aoc/input/day7")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	return input

}

func Day7() {
	resultPart1 := day7part1(inputDay7(false))
	fmt.Println(resultPart1)
	resultPart2 := day7part2(inputDay7(false))
	fmt.Println(resultPart2)
}
