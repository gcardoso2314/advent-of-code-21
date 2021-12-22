package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func readFile(path string) [][]int {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var risk_map [][]int
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		var row []int
		for _, c := range text {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		risk_map = append(risk_map, row)
	}

	return risk_map
}

func enlargeMap(risk_map [][]int, factor int) [][]int {
	var enlarged_map [][]int
	for s := 0; s < factor; s++ {
		for _, row := range risk_map {
			var new_row []int
			for i := 0; i < factor; i++ {
				for ind := 0; ind < len(row); ind++ {
					if row[ind]+i+s > 9 {
						new_row = append(new_row, (row[ind]+i+s)%9)
					} else {
						new_row = append(new_row, row[ind]+i+s)
					}
				}
			}
			enlarged_map = append(enlarged_map, new_row)
		}
	}

	return enlarged_map
}

func heuristic(a, b [2]int) float64 {
	return math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1]))
}

type Node struct {
	pos      [2]int
	priority int
}

func getNeighbours(node [2]int, max int) [][2]int {
	var neighbours [][2]int
	if node[0] > 0 {
		up := [2]int{node[0] - 1, node[1]}
		neighbours = append(neighbours, up)
	}
	if node[0] < max {
		down := [2]int{node[0] + 1, node[1]}
		neighbours = append(neighbours, down)
	}
	if node[1] > 0 {
		left := [2]int{node[0], node[1] - 1}
		neighbours = append(neighbours, left)
	}
	if node[1] < max {
		right := [2]int{node[0], node[1] + 1}
		neighbours = append(neighbours, right)
	}
	return neighbours
}

func aStarSearch(risk_map [][]int) int {
	max := len(risk_map) - 1
	start := [2]int{0, 0}
	target := [2]int{max, max}
	queue := []Node{{
		pos:      start,
		priority: 0,
	}}
	cost_accrued := make(map[[2]int]int)
	cost_accrued[start] = 0
	predecessor := make(map[[2]int][2]int)

	for {
		if len(queue) == 0 {
			break
		}
		sort.Slice(queue, func(i, j int) bool { return queue[i].priority < queue[j].priority })
		current := queue[0]
		if current.pos == target {
			break
		}

		neighbours := getNeighbours(current.pos, max)
		for _, n := range neighbours {
			new_cost := cost_accrued[current.pos] + risk_map[n[0]][n[1]]
			cost, exists := cost_accrued[n]
			if !exists || new_cost < cost {
				predecessor[n] = current.pos
				cost_accrued[n] = new_cost
				queue = append(queue, Node{
					pos:      n,
					priority: new_cost + int(heuristic(n, target)),
				})
			}
		}
		queue = queue[1:]
	}
	return cost_accrued[target]
}

func main() {
	risk_map := readFile("input.txt")
	enlarged_map := enlargeMap(risk_map, 5)
	fmt.Println(aStarSearch(risk_map))
	fmt.Println(aStarSearch(enlarged_map))
}
