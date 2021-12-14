package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func A(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var counter int
	for i := 0; scanner.Scan(); i++ {
		four_digits := strings.Fields(strings.Split(scanner.Text(), "|")[1])
		for _, digit := range four_digits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				counter++
			}
		}
	}

	return counter
}

func SetIntersection(set map[string]bool, digit string) int {
	var n_intersection int
	for _, char := range digit {
		if set[string(char)] {
			n_intersection++
		}
	}
	return n_intersection
}

func SetEquality(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k, _ := range a {
		if !b[k] {
			return false
		}
	}
	return true
}

func B(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var counter int
	for i := 0; scanner.Scan(); i++ {
		numbers := make(map[int]map[string]bool)
		for i := 0; i < 10; i++ {
			numbers[i] = make(map[string]bool)
		}
		split_text := strings.Split(scanner.Text(), "|")
		sample_digits, four_digits := strings.Fields(split_text[0]), strings.Fields(split_text[1])

		// Get the unique length numbers
		for _, digit := range sample_digits {
			if len(digit) == 2 {
				for _, char := range digit {
					numbers[1][string(char)] = true
				}
			} else if len(digit) == 3 {
				for _, char := range digit {
					numbers[7][string(char)] = true
				}
			} else if len(digit) == 4 {
				for _, char := range digit {
					numbers[4][string(char)] = true
				}
			} else if len(digit) == 7 {
				for _, char := range digit {
					numbers[8][string(char)] = true
				}
			}
		}
		// find all other digits
		for _, digit := range sample_digits {
			if len(digit) == 5 {
				if SetIntersection(numbers[7], digit) == 3 {
					for _, char := range digit {
						numbers[3][string(char)] = true
					}
				} else if SetIntersection(numbers[4], digit) == 2 {
					for _, char := range digit {
						numbers[2][string(char)] = true
					}
				} else if SetIntersection(numbers[4], digit) == 3 {
					for _, char := range digit {
						numbers[5][string(char)] = true
					}
				}
			} else if len(digit) == 6 {
				if SetIntersection(numbers[4], digit) == 4 {
					for _, char := range digit {
						numbers[9][string(char)] = true
					}
				} else if SetIntersection(numbers[1], digit) == 2 {
					for _, char := range digit {
						numbers[0][string(char)] = true
					}
				} else {
					for _, char := range digit {
						numbers[6][string(char)] = true
					}
				}
			}
		}
		for e, digit := range four_digits {
			set := make(map[string]bool)
			for _, char := range digit {
				set[string(char)] = true
			}
			for n, val := range numbers {
				if SetEquality(val, set) {
					counter += n * int(math.Pow(10.0, float64(3-e)))
				}
			}
		}
	}
	return counter
}

func main() {
	fmt.Println(B("input.txt"))
}
