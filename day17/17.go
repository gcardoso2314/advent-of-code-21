package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readFile(path string) (int, int, int, int) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lower_x int
	var upper_x int
	var lower_y int
	var upper_y int
	r := regexp.MustCompile(`-?[0-9]*\.\.-?[0-9]*`)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		occurrences := r.FindAllString(text, -1)

		x_vals := strings.Split(occurrences[0], "..")
		lower_x, _ = strconv.Atoi(x_vals[0])
		upper_x, _ = strconv.Atoi(x_vals[1])

		y_vals := strings.Split(occurrences[1], "..")
		lower_y, _ = strconv.Atoi(y_vals[0])
		upper_y, _ = strconv.Atoi(y_vals[1])

	}
	return lower_x, upper_x, lower_y, upper_y
}

func minXVelocity(lower_x int) int {
	var min_x int
	for {
		if (min_x * min_x) >= 2*lower_x-min_x {
			break
		}
		min_x++
	}
	return min_x
}

func findAllPossibleVelocities(lower_x, upper_x, lower_y, upper_y int) [][2]int {
	x_velocity_min := minXVelocity(lower_x)
	x_velocity_max := upper_x
	y_velocity_min := lower_y
	y_velocity_max := math.Abs(float64(lower_y))

	var possible_velocities [][2]int
	for x_init := x_velocity_min; x_init <= x_velocity_max; x_init++ {
		for y_init := y_velocity_min; y_init <= int(y_velocity_max); y_init++ {
			x_velocity := x_init
			y_velocity := y_init
			pos_x := 0
			pos_y := 0

			var target_hit bool
			for {
				pos_x += x_velocity
				pos_y += y_velocity
				if (pos_x >= lower_x) && (pos_x <= upper_x) && (pos_y >= lower_y) && (pos_y <= upper_y) {
					target_hit = true
					break
				} else if pos_y < lower_y {
					// gone to far
					break
				}
				if x_velocity < 0 {
					x_velocity++
				} else if x_velocity > 0 {
					x_velocity--
				}
				y_velocity--
			}
			if target_hit {
				possible_velocities = append(possible_velocities, [2]int{x_init, y_init})
			}
		}
	}
	return possible_velocities
}

func main() {
	lower_x, upper_x, lower_y, upper_y := readFile("input.txt")

	all_velocities := findAllPossibleVelocities(lower_x, upper_x, lower_y, upper_y)

	sort.Slice(all_velocities, func(i, j int) bool { return all_velocities[i][1] > all_velocities[j][1] })
	highest_point := (all_velocities[0][1] * (all_velocities[0][1] + 1)) / 2
	fmt.Println("Answer to Part 1 is", highest_point)

	fmt.Println("Answer to Part 2 is", len(all_velocities))
}
