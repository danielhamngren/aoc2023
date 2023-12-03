package main

import (
	"flag"
	"fmt"
	"log"
	"main/aoc"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string) int {
	re := regexp.MustCompile(`[0-9]`)
	result := 0

	for i := range input {

		digits := re.FindAllString(input[i], -1)

		cal, err := strconv.Atoi(digits[0] + digits[len(digits)-1])

		if err != nil {
			panic(err)
		}

		result += cal

	}

	return result

}

func digitConv(x string) string {
	digits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}
	return digits[x]
}

func handleLine(x string) string {
	var first = make(map[string]int)
	var last = make(map[string]int)

	digits := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"0":     "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}

	for key, _ := range digits {
		first[key] = strings.Index(x, key)
		last[key] = strings.LastIndex(x, key)
	}

	var firstKey string
	firstIndex := 10000
	for key, val := range first {
		if val < firstIndex && val > -1 {
			firstIndex = val
			firstKey = key
		}
	}

	var lastKey string
	lastIndex := -1
	for key, val := range last {
		if val > lastIndex {
			lastIndex = val
			lastKey = key
		}
	}
	return digits[firstKey] + digits[lastKey]
}

func part2(input []string) int {
	re := regexp.MustCompile(`[0-9]`)
	result := 0

	// fmt.Println(input)

	reDigits := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|zero")

	for i := range input {
		input[i] = handleLine(input[i])
		input[i] = reDigits.ReplaceAllStringFunc(input[i], digitConv)
	}

	// fmt.Println(input)

	for i := range input {
		digits := re.FindAllString(input[i], -1)

		cal, err := strconv.Atoi(digits[0] + digits[len(digits)-1])

		if err != nil {
			panic(err)
		}

		result += cal

	}

	return result
}

func day1() {
	// input := []string{
	// 	"1abc2",
	// 	"pqr3stu8vwx",
	// 	"a1b2c3d4e5f",
	// 	"treb7uchet",
	// }

	// input := []string{
	// 	"two1nine",
	// 	"eightwothree",
	// 	"abcone2threexyz",
	// 	"xtwone3four",
	// 	"4nineeightseven2",
	// 	"zoneight234",
	// 	"7pqrstsixteen",
	// }

	// Read input from file
	content, err := os.ReadFile("input1-1")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	// result1 := part1(input)
	// fmt.Println(result1)

	result2 := part2(input)
	fmt.Println(result2)

}

func checkRound(round string) (int, int, int) {
	reDigits := regexp.MustCompile(`\d+`)
	cubes := strings.Split(round, ",")
	red := 0
	green := 0
	blue := 0
	for i := range cubes {
		number, _ := strconv.Atoi(reDigits.FindString(cubes[i]))
		if strings.Contains(cubes[i], "red") {
			red += number
		}
		if strings.Contains(cubes[i], "green") {
			green += number
		}
		if strings.Contains(cubes[i], "blue") {
			blue += number
		}
	}
	return red, green, blue
}

func getRoundBalls(round string) (int, int, int) {
	reDigits := regexp.MustCompile(`\d+`)
	cubes := strings.Split(round, ",")
	red := 0
	green := 0
	blue := 0
	for i := range cubes {
		number, _ := strconv.Atoi(reDigits.FindString(cubes[i]))
		if strings.Contains(cubes[i], "red") {
			red = number
		}
		if strings.Contains(cubes[i], "green") {
			green = number
		}
		if strings.Contains(cubes[i], "blue") {
			blue = number
		}
	}
	return red, green, blue
}

func checkGamePower(x string) int {
	tokens := strings.Split(x, ":")

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	reveals := strings.Split(tokens[1], ";")
	for i := range reveals {
		red, green, blue := getRoundBalls(reveals[i])
		if red > maxRed {
			maxRed = red
		}
		if green > maxGreen {
			maxGreen = green
		}
		if blue > maxBlue {
			maxBlue = blue
		}
	}

	return maxRed * maxGreen * maxBlue
}
func checkValid(x string) (int, bool) {

	tokens := strings.Split(x, ":")
	gameId, _ := strconv.Atoi(strings.Fields(tokens[0])[1])

	reveals := strings.Split(tokens[1], ";")
	for i := range reveals {
		red, green, blue := checkRound(reveals[i])
		if red > 12 || green > 13 || blue > 14 {
			return gameId, false
		}
	}

	return gameId, true
}

func day2part1(input []string) int {
	// validIds := []int{}
	idSum := 0

	for i := range input {
		id, valid := checkValid(input[i])
		if valid {
			idSum += id
		}
	}

	return idSum
}

func day2part2(input []string) int {
	totalPower := 0

	for i := range input {
		power := checkGamePower(input[i])
		totalPower += power
	}

	return totalPower
}

func day2() {

	// test_input := []string{
	// 	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	// 	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	// 	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	// 	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	// 	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	// }

	//Read input from file
	content, err := os.ReadFile("input2")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	resultPart1 := day2part1(input)
	fmt.Println(resultPart1)
	resultPart2 := day2part2(input)
	fmt.Println(resultPart2)
}

func main() {
	dayPtr := flag.Int("day", 0, "day of the problem")
	flag.Parse()

	switch *dayPtr {
	case 1:
		day1()
	case 2:
		day2()
	case 3:
		aoc.Day3()
	case 4:
		aoc.Day4()
	default:
		fmt.Println(*dayPtr, "Not yet implemented")

	}

}
