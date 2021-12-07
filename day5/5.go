package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) [][2][2]int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][2][2]int
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		line := strings.Split(text, " -> ")
		start, end := strings.Split(line[0], ","), strings.Split(line[1], ",")
		x1, _ := strconv.Atoi(start[0])
		x2, _ := strconv.Atoi(end[0])
		y1, _ := strconv.Atoi(start[1])
		y2, _ := strconv.Atoi(end[1])
		parsed_line := [2][2]int{{x1, y1}, {x2, y2}}
		lines = append(lines, parsed_line)
	}

	return lines
}

func drawDiagram(lines [][2][2]int, include_diagonal bool) [][]int {
	// Find max x and y values
	var x_max int
	var y_max int
	for _, line := range lines {
		x1, y1, x2, y2 := line[0][0], line[0][1], line[1][0], line[1][1]
		if x1 > x_max {
			x_max = x1
		}
		if x2 > x_max {
			x_max = x2
		}
		if y1 > y_max {
			y_max = y1
		}
		if y2 > y_max {
			y_max = y2
		}
	}
	diagram := make([][]int, x_max+1)
	for i := 0; i <= x_max; i++ {
		diagram[i] = make([]int, y_max+1)
	}

	// Iterate through line segments
	for _, line := range lines {
		x1, y1, x2, y2 := line[0][0], line[0][1], line[1][0], line[1][1]
		if !(x1 == x2 || y1 == y2) && !include_diagonal {
			continue
		}
		var min_x int
		var max_x int
		var min_y int
		var max_y int
		if x1 <= x2 {
			min_x = x1
			max_x = x2
		} else {
			min_x = x2
			max_x = x1
		}
		if y1 <= y2 {
			min_y = y1
			max_y = y2
		} else {
			min_y = y2
			max_y = y1
		}
		var is_diagonal bool
		if x1 != x2 && y1 != y2 {
			is_diagonal = true
		}
		if !is_diagonal {
			for i := min_x; i <= max_x; i++ {
				for j := min_y; j <= max_y; j++ {
					diagram[i][j]++
				}
			}
		} else {
			for i := 0; i <= max_x-min_x; i++ {
				if x1 == min_x {
					if y1 == min_y {
						diagram[x1+i][y1+i]++
					} else {
						diagram[x1+i][y1-i]++
					}
				} else {
					if y1 == min_y {
						diagram[x1-i][y1+i]++
					} else {
						diagram[x1-i][y1-i]++
					}
				}
			}
		}

	}
	return diagram
}

func countDangerousPoints(diagram [][]int) int {
	var dangerous_points int
	for _, row := range diagram {
		for _, val := range row {
			if val > 1 {
				dangerous_points++
			}
		}
	}
	return dangerous_points
}

func main() {
	lines := readFile("input.txt")
	diagram := drawDiagram(lines, true)
	// fmt.Println(diagram)
	fmt.Println(countDangerousPoints(diagram))
}
