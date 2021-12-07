package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) map[int]int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := make(map[int]int)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		initial_state := strings.Split(text, ",")
		for _, val := range initial_state {
			int_value, _ := strconv.Atoi(val)
			state[int_value] += 1
		}
	}

	return state
}

func advanceOneDay(state map[int]int) map[int]int {
	new_state := make(map[int]int)
	for timer, count := range state {
		if timer == 0 {
			// Set them back to timer=6
			new_state[6] += count
			// Each of them produces one more lanternfish
			new_state[8] += count
		} else {
			new_state[timer-1] += count
		}
	}
	return new_state
}

func main() {
	state := readFile("input.txt")
	n_days := 256
	var total_fish int
	for i := 0; i < n_days; i++ {
		state = advanceOneDay(state)
	}
	for _, count := range state {
		total_fish += count
	}
	fmt.Println("After", n_days, "days there are", total_fish, "fish")
}
