package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func A() {
	lines, _ := readLines("input.txt")
	var counter int
	for i := 1; i < len(lines); i++ {
		new_value, _ := strconv.Atoi(lines[i])
		old_value, _ := strconv.Atoi(lines[i-1])
		if new_value > old_value {
			counter += 1
		}
	}

	fmt.Printf("Solution for part 1 is %d.", counter)
}

func B() {
	lines, _ := readLines("input.txt")
	var converted_lines []int
	for i := 0; i < len(lines); i++ {
		int_value, _ := strconv.Atoi(lines[i])
		converted_lines = append(converted_lines, int_value)
	}
	var counter int
	for i := 0; i < len(converted_lines)-3; i++ {
		old_value := converted_lines[i] + converted_lines[i+1] + converted_lines[i+2]
		new_value := converted_lines[i+1] + converted_lines[i+2] + converted_lines[i+3]
		if new_value > old_value {
			counter += 1
		}
	}

	fmt.Printf("Solution for part 2 is %d.", counter)
}

func main() {
	A()
	B()
}
