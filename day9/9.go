package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Min(a []int) int {
	min := 9999999999
	for _, val := range a {
		if val < min {
			min = val
		}
	}
	return min
}

func A(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cave_map [][]int
	for i := 0; scanner.Scan(); i++ {
		var row []int
		for _, val := range scanner.Text() {
			n, _ := strconv.Atoi(string(val))
			row = append(row, n)
		}
		cave_map = append(cave_map, row)
	}

	var risk_score int
	for i, row := range cave_map {
		for j, val := range row {
			switch {
			case i != 0 && cave_map[i-1][j] <= val:
				break
			case j != 0 && cave_map[i][j-1] <= val:
				break
			case i != len(cave_map)-1 && cave_map[i+1][j] <= val:
				break
			case j != len(row)-1 && cave_map[i][j+1] <= val:
				break
			default:
				risk_score += 1 + val
			}
		}
	}
	return risk_score
}

func ClimbUp(cave_map [][]int, start [2]int) [][2]int {
	var basin [][2]int
	i, j := start[0], start[1]
	val := cave_map[i][j]
	basin = append(basin, [2]int{i, j})

	if i != 0 && cave_map[i-1][j] > val && cave_map[i-1][j] < 9 {
		basin = append(basin, ClimbUp(cave_map, [2]int{i - 1, j})...)
	}
	if j != 0 && cave_map[i][j-1] > val && cave_map[i][j-1] < 9 {
		basin = append(basin, ClimbUp(cave_map, [2]int{i, j - 1})...)
	}
	if i != len(cave_map)-1 && cave_map[i+1][j] > val && cave_map[i+1][j] < 9 {
		basin = append(basin, ClimbUp(cave_map, [2]int{i + 1, j})...)
	}
	if j != len(cave_map[0])-1 && cave_map[i][j+1] > val && cave_map[i][j+1] < 9 {
		basin = append(basin, ClimbUp(cave_map, [2]int{i, j + 1})...)
	}
	return basin
}

func B(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cave_map [][]int
	for i := 0; scanner.Scan(); i++ {
		var row []int
		for _, val := range scanner.Text() {
			n, _ := strconv.Atoi(string(val))
			row = append(row, n)
		}
		cave_map = append(cave_map, row)
	}

	var low_points [][2]int
	for i, row := range cave_map {
		for j, val := range row {
			switch {
			case i != 0 && cave_map[i-1][j] <= val:
				break
			case j != 0 && cave_map[i][j-1] <= val:
				break
			case i != len(cave_map)-1 && cave_map[i+1][j] <= val:
				break
			case j != len(row)-1 && cave_map[i][j+1] <= val:
				break
			default:
				low_points = append(low_points, [2]int{i, j})
			}
		}
	}

	var basin_sizes []int
	for _, low_point := range low_points {
		basin_points := ClimbUp(cave_map, low_point)
		basin_set := make(map[[2]int]bool)
		for _, p := range basin_points {
			basin_set[p] = true
		}
		size := len(basin_set)
		basin_sizes = append(basin_sizes, size)
	}

	sort.Slice(basin_sizes, func(p, q int) bool {
		return basin_sizes[p] > basin_sizes[q]
	})

	return basin_sizes[0] * basin_sizes[1] * basin_sizes[2]
}

func main() {
	fmt.Println(A("input.txt"))
	fmt.Println(B("input.txt"))
}
