package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(path string) [10][10]int {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var octopuses [10][10]int
	for i := 0; scanner.Scan(); i++ {
		for j, n := range scanner.Text() {
			x, _ := strconv.Atoi(string(n))
			octopuses[i][j] = x
		}
	}

	return octopuses
}

func increaseAdjacent(octopuses [10][10]int, flashing_oct [2]int) [10][10]int {
	i, j := flashing_oct[0], flashing_oct[1]

	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x >= 0 && x < 10 && y >= 0 && y < 10 && (x != i || y != j) {
				octopuses[x][y]++
				if octopuses[x][y] == 10 { // flashing
					octopuses = increaseAdjacent(octopuses, [2]int{x, y})
				}
			}
		}
	}

	return octopuses
}

func A(path string) int {
	octopuses := readFile(path)

	var flashes int
	for step := 1; step <= 100; step++ {
		// Increase all values and adjacent values if flashing
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopuses[i][j]++
				if octopuses[i][j] == 10 {
					octopuses = increaseAdjacent(octopuses, [2]int{i, j})
				}
			}
		}
		// Count flashes and reset
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octopuses[i][j] > 9 {
					flashes++
					octopuses[i][j] = 0
				}
			}
		}
	}
	return flashes
}

func B(path string) int {
	octopuses := readFile(path)

	var step int
	for {
		step++
		all_flash := true

		// Increase all values and adjacent values if flashing
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopuses[i][j]++
				if octopuses[i][j] == 10 {
					octopuses = increaseAdjacent(octopuses, [2]int{i, j})
				}
			}
		}
		// Count flashes and reset
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octopuses[i][j] > 9 {
					octopuses[i][j] = 0
					continue
				}
				all_flash = false
			}
		}
		if all_flash {
			break
		}
	}
	return step
}

func main() {
	fmt.Println(B("input.txt"))
}
