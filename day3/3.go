package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CountOccurences(val_list []string) [12]int {
	var sums [12]int
	for i := 0; i < len(val_list); i++ {
		text := string(val_list[i])
		for j := 0; j < len(text); j++ {
			val, _ := strconv.Atoi(string(text[j]))
			sums[j] += val
		}
	}
	return sums
}

func A() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var val_list []string
	for i := 0; scanner.Scan(); i++ {
		val_list = append(val_list, scanner.Text())
	}

	sums := CountOccurences(val_list)
	n := len(val_list)

	var gamma string
	var epsilon string
	for _, val := range sums {
		if val > n/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gamma_decimal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon_decimal, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gamma_decimal * epsilon_decimal)
}

func B() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var val_list []string
	for i := 0; scanner.Scan(); i++ {
		val_list = append(val_list, scanner.Text())
	}

	sums := CountOccurences(val_list)

	oxygen_keep := val_list
	oxygen_n := len(oxygen_keep)
	for ind := 0; ind < len(sums); ind++ {
		if len(oxygen_keep) == 1 {
			break
		}
		sums := CountOccurences(oxygen_keep)
		var val_to_keep string
		if float64(sums[ind]) >= float64(oxygen_n)/2.0 {
			val_to_keep = "1"
		} else {
			val_to_keep = "0"
		}
		var new_keep []string
		for _, binary := range oxygen_keep {
			if string(binary[ind]) == val_to_keep {
				new_keep = append(new_keep, binary)
			}
		}
		oxygen_keep = new_keep
		oxygen_n = len(oxygen_keep)
	}
	oxygen_gen_rating := oxygen_keep[0]

	co2_keep := val_list
	co2_n := len(co2_keep)
	for ind := 0; ind < len(sums); ind++ {
		if len(co2_keep) == 1 {
			break
		}
		sums := CountOccurences(co2_keep)
		var val_to_keep string
		if float64(sums[ind]) >= float64(co2_n)/2.0 {
			val_to_keep = "0"
		} else {
			val_to_keep = "1"
		}
		var new_keep []string
		for _, binary := range co2_keep {
			if string(binary[ind]) == val_to_keep {
				new_keep = append(new_keep, binary)
			}
		}
		co2_keep = new_keep
		co2_n = len(co2_keep)
	}
	co2_scrubber_rating := co2_keep[0]

	oxygen_decimal, _ := strconv.ParseInt(oxygen_gen_rating, 2, 64)
	co2_decimal, _ := strconv.ParseInt(co2_scrubber_rating, 2, 64)
	fmt.Println(oxygen_decimal * co2_decimal)
}

func main() {
	A()
	B()
}
