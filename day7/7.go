package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var positions []int
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		pos := strings.Split(text, ",")
		for _, val := range pos {
			int_value, _ := strconv.Atoi(val)
			positions = append(positions, int_value)
		}
	}

	return positions
}

func costFunctionPart1(positions []int, guess int) int {
	var cost float64
	for _, pos := range positions {
		cost += math.Abs(float64(pos - guess))
	}
	return int(cost)
}

func costFunctionPart2(positions []int, guess int) int {
	var cost float64
	for _, pos := range positions {
		cost += (math.Abs(float64(pos-guess)) * (math.Abs(float64(pos-guess)) + 1)) / 2
	}
	return int(cost)
}

func findOptimalPos(positions []int, cost_function func([]int, int) int) int {
	// Use average as best guess
	var sum_positions int
	for _, pos := range positions {
		sum_positions += pos
	}
	best_guess := sum_positions / len(positions)
	min_cost := cost_function(positions, best_guess)
	for {
		if cost_function(positions, best_guess-1) < min_cost {
			best_guess = best_guess - 1
			min_cost = cost_function(positions, best_guess)
		} else if cost_function(positions, best_guess+1) < min_cost {
			best_guess = best_guess + 1
			min_cost = cost_function(positions, best_guess)
		} else {
			break
		}
	}
	return best_guess
}

func main() {
	positions := readFile("input.txt")
	best_pos_1 := findOptimalPos(positions, costFunctionPart1)
	best_pos_2 := findOptimalPos(positions, costFunctionPart2)
	fmt.Println("Total fuel spent for Part 1 is:", costFunctionPart1(positions, best_pos_1), "for position", best_pos_1)
	fmt.Println("Total fuel spent for Part 2 is:", costFunctionPart2(positions, best_pos_2), "for position", best_pos_2)
}
