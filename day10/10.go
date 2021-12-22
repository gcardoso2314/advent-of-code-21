package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var oppositeSymbol = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
	"<": ">",
}

func checkSymbolOpen(symbol string) bool {
	if symbol == "(" || symbol == "{" || symbol == "[" || symbol == "<" {
		return true
	}
	return false
}

func A(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	symbolScores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	scanner := bufio.NewScanner(file)
	var total_score int
	for i := 0; scanner.Scan(); i++ {
		var open_symbols []string
		for pos, c := range scanner.Text() {
			symbol := string(c)
			if checkSymbolOpen(symbol) {
				open_symbols = append(open_symbols, symbol)
				continue
			} else if pos == 0 {
				total_score += symbolScores[symbol]
				break
			} else if pos > 0 && symbol != oppositeSymbol[open_symbols[len(open_symbols)-1]] {
				total_score += symbolScores[symbol]
				break
			} else {
				open_symbols = open_symbols[:len(open_symbols)-1]
			}

		}
	}

	return total_score
}

func B(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	symbolScores := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	scanner := bufio.NewScanner(file)
	var total_scores []int
	for i := 0; scanner.Scan(); i++ {
		var total_score int
		var open_symbols []string
		var corrupted bool
		for pos, c := range scanner.Text() {
			symbol := string(c)
			if checkSymbolOpen(symbol) {
				open_symbols = append(open_symbols, symbol)
				continue
			} else if pos == 0 {
				corrupted = true
				break
			} else if pos > 0 && symbol != oppositeSymbol[open_symbols[len(open_symbols)-1]] {
				corrupted = true
				break
			} else {
				open_symbols = open_symbols[:len(open_symbols)-1]
			}
		}
		if len(open_symbols) > 0 && !corrupted {
			for j := len(open_symbols) - 1; j >= 0; j-- {
				total_score *= 5
				total_score += symbolScores[oppositeSymbol[open_symbols[j]]]
			}
			total_scores = append(total_scores, total_score)
		}
	}

	sort.Ints(total_scores)

	return total_scores[len(total_scores)/2]

}

func main() {
	fmt.Println(A("input.txt"))
	fmt.Println(B("input.txt"))
}
