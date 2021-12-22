package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) (map[[2]int]bool, []string, []int) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	points := make(map[[2]int]bool)
	var fold_direction []string
	var fold_line []int
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if strings.HasPrefix(line, "fold along x") {
			x, _ := strconv.Atoi(strings.Split(line, "=")[1])
			fold_direction = append(fold_direction, "x")
			fold_line = append(fold_line, x)
		} else if strings.HasPrefix(line, "fold along y") {
			y, _ := strconv.Atoi(strings.Split(line, "=")[1])
			fold_direction = append(fold_direction, "y")
			fold_line = append(fold_line, y)
		} else if strings.Contains(line, ",") {
			var new_point [2]int
			for j, c := range strings.Split(line, ",") {
				n, _ := strconv.Atoi(c)
				new_point[j] = n
			}
			points[new_point] = true
		}
	}

	return points, fold_direction, fold_line
}

func fold(points map[[2]int]bool, direction string, line int) map[[2]int]bool {
	for p, _ := range points {
		if (direction == "x" && p[0] > line) || (direction == "y" && p[1] > line) {
			var new_point [2]int
			if direction == "x" {
				new_point = [2]int{2*line - p[0], p[1]}

			} else {
				new_point = [2]int{p[0], 2*line - p[1]}
			}
			points[new_point] = true
			delete(points, p)

		}
	}
	return points
}

func A(path string) int {
	points, fold_direction, fold_line := readFile(path)

	direction := fold_direction[0]
	line := fold_line[0]

	points = fold(points, direction, line)

	// Count points
	var counter int
	for _, v := range points {
		if v {
			counter++
		}
	}

	return counter

}

func maxValues(points map[[2]int]bool) (int, int) {
	var max_x int
	var max_y int
	for p, v := range points {
		if v {
			if p[0] > max_x {
				max_x = p[0]
			}
			if p[1] > max_y {
				max_y = p[1]
			}
		}
	}

	return max_x, max_y
}

func B(path string) {
	points, fold_direction, fold_line := readFile(path)

	for i, direction := range fold_direction {
		fmt.Println("folding across", direction, "at", fold_line[i])
		points = fold(points, direction, fold_line[i])
	}

	max_x, max_y := maxValues(points)

	for i := 0; i <= max_y; i++ {
		var line string
		for j := 0; j <= max_x; j++ {
			if points[[2]int{j, i}] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func main() {
	fmt.Println(A("input.txt"))
	B("input.txt")
}
