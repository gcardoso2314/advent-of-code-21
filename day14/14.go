package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(path string) (string, map[string]string) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var template string
	instructions := make(map[string]string)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		if i == 0 {
			template = text
		} else if strings.Contains(text, "->") {
			new_instr := strings.Split(text, " -> ")
			instructions[new_instr[0]] = new_instr[1]
		}
	}
	return template, instructions
}

func A(path string) int {
	template, instructions := readFile(path)

	letter_counts := make(map[string]int)
	for _, c := range template {
		letter_counts[string(c)]++
	}

	for step := 0; step < 10; step++ {
		var new_string string
		for i := 0; i < len(template)-1; i++ {
			new_string += string(template[i]) + instructions[template[i:i+2]]
			letter_counts[instructions[template[i:i+2]]]++
		}
		template = new_string + string(template[len(template)-1])
	}

	min := letter_counts[string(template[0])]
	max := letter_counts[string(template[0])]
	for _, count := range letter_counts {
		if count < min {
			min = count
		} else if count > max {
			max = count
		}
	}

	return max - min

}

func B(path string) int {
	template, instructions := readFile(path)

	letter_counts := make(map[string]int)
	for _, c := range template {
		letter_counts[string(c)]++
	}

	polymer_pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		polymer_pairs[template[i:i+2]]++
	}

	for step := 0; step < 40; step++ {
		new_polymers := make(map[string]int)
		for pair, num := range polymer_pairs {
			if instructions[pair] != "" && num > 0 {
				letter_counts[instructions[pair]] += num
				new_polymers[string(pair[0])+instructions[pair]] += num
				new_polymers[instructions[pair]+string(pair[1])] += num
				polymer_pairs[pair] -= num
			}
		}
		for pair, num := range new_polymers {
			polymer_pairs[pair] += num
		}
	}

	min := letter_counts[string(template[0])]
	max := letter_counts[string(template[0])]
	for _, count := range letter_counts {
		if count < min {
			min = count
		} else if count > max {
			max = count
		}
	}

	return max - min
}

func main() {
	fmt.Println(A("input.txt"))
	fmt.Println(B("input.txt"))
}
