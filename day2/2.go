package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var depth int
	var horizontal int
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		instructions := strings.Fields(text)

		direction := instructions[0]
		value, _ := strconv.Atoi(instructions[1])
		switch direction {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	fmt.Println(depth * horizontal)
}

func B() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var depth int
	var horizontal int
	var aim int
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		instructions := strings.Fields(text)

		direction := instructions[0]
		value, _ := strconv.Atoi(instructions[1])
		switch direction {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	fmt.Println(depth * horizontal)
}

func main() {
	A()
	B()
}
