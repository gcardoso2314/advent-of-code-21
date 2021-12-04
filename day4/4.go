package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) ([][5][5]int, []int) {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var number_draws []int
	var bingo_cards [][5][5]int
	var bingo_index int
	var bingo_card [5][5]int
	var row int
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		if i == 0 {
			// First line contains the number draws (comma separated)
			for _, num := range strings.Split(text, ",") {
				num_int, _ := strconv.Atoi(num)
				number_draws = append(number_draws, num_int)
			}
		} else if i%6 == 2 {
			// Create new bingo card
			bingo_index = i / 6
			row = 0

			// New bingo card every 6 lines starting at line 2
			if bingo_index != 0 {
				bingo_cards = append(bingo_cards, bingo_card)
			}

			for j, num := range strings.Fields(text) {
				num_int, _ := strconv.Atoi(num)
				bingo_card[row][j] = num_int
			}
			row++
		} else if text != "" {
			for j, num := range strings.Fields(text) {
				num_int, _ := strconv.Atoi(num)
				bingo_card[row][j] = num_int
			}
			row++
		}
	}
	// Add last card
	bingo_cards = append(bingo_cards, bingo_card)

	return bingo_cards, number_draws
}

func itemExists(arr []int, val int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

func allNumbersDrawn(numbers_drawn []int, arr [5]int) bool {
	for _, val := range arr {
		if !itemExists(numbers_drawn, val) {
			return false
		}
	}
	return true
}

func checkBingo(numbers_drawn []int, bingo_card [5][5]int) bool {
	for i := 0; i < 5; i++ {
		// Check bingo in rows
		if allNumbersDrawn(numbers_drawn, bingo_card[i]) {
			return true
		}
		// Check bingo in columns
		var col [5]int
		for j := 0; j < 5; j++ {
			col[j] = bingo_card[j][i]
		}
		if allNumbersDrawn(numbers_drawn, col) {
			return true
		}
	}
	return false
}

func addRemainingCardNumbers(numbers_drawn []int, bingo_card [5][5]int) int {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !itemExists(numbers_drawn, bingo_card[i][j]) {
				sum += bingo_card[i][j]
			}
		}
	}
	return sum
}

func removeElement(arr []int, element int) []int {
	var new_arr []int
	for i, val := range arr {
		if val == element {
			new_arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return new_arr
}

func A(bingo_cards [][5][5]int, number_draws []int) int {
	var final_answer int
	var numbers_drawn []int
	for _, num := range number_draws {
		numbers_drawn = append(numbers_drawn, num)
		if len(numbers_drawn) < 5 {
			continue
		}
		bingo_index := -1
		for i, card := range bingo_cards {
			if checkBingo(numbers_drawn, card) {
				bingo_index = i
				break
			}
		}

		if bingo_index != -1 {
			final_answer = num * addRemainingCardNumbers(numbers_drawn, bingo_cards[bingo_index])
			break
		}
	}
	return final_answer
}

func B(bingo_cards [][5][5]int, number_draws []int) int {
	var final_answer int
	var numbers_drawn []int

	// Create a slice containing bingo cards still waiting for bingo
	cards_remaining := make([]int, len(bingo_cards))
	for i := 0; i < len(bingo_cards); i++ {
		cards_remaining[i] = i
	}

	for _, num := range number_draws {
		numbers_drawn = append(numbers_drawn, num)
		if len(numbers_drawn) < 5 {
			continue
		}
		// Check which cards had bingo with this number
		var cards_with_bingo []int
		for _, card_index := range cards_remaining {
			if checkBingo(numbers_drawn, bingo_cards[card_index]) {
				cards_with_bingo = append(cards_with_bingo, card_index)
			}
		}

		// If only one card left and it has bingo then this is the final answer
		if len(cards_remaining) == 1 && len(cards_with_bingo) > 0 {
			final_answer = num * addRemainingCardNumbers(numbers_drawn, bingo_cards[cards_with_bingo[0]])
			break
		}
		// Otherwise remove all cards that just got bingo and keep iterating
		for _, bingo_index := range cards_with_bingo {
			cards_remaining = removeElement(cards_remaining, bingo_index)
		}
	}
	return final_answer
}

func main() {
	bingo_cards, number_draws := readFile("input.txt")
	fmt.Println("Solution for Part 1 is", A(bingo_cards, number_draws))
	fmt.Println("Solution for Part 2 is", B(bingo_cards, number_draws))

}
