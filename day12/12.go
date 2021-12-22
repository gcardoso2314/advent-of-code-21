package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(file_path string) map[string][]string {
	f, _ := os.Open(file_path)
	defer f.Close()

	paths := make(map[string][]string)
	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		path := strings.Split(text, "-")
		paths[path[0]] = append(paths[path[0]], path[1])
		paths[path[1]] = append(paths[path[1]], path[0])
	}

	return paths
}

func isLower(node string) bool {
	return strings.ToLower(node) == node
}

func alreadyVisited(visited []string, node string) bool {
	for _, v := range visited {
		if v == node {
			return true
		}
	}
	return false
}

func A(file_path string) [][]string {
	paths := readFile(file_path)
	var all_possible_paths [][]string
	visited := []string{"start"}
	var stack [][]string
	stack = append(stack, paths["start"])
	for {
		// if nothing left to visit
		if len(stack) <= 0 {
			break
		}
		children := stack[len(stack)-1]
		if len(children) == 0 {
			// No further children to explore
			stack = append(stack[:len(stack)-1])
			// Didn't get to "end" so remove the last visited node
			if len(visited) > 0 {
				visited = append(visited[:len(visited)-1])
			}
		} else {
			child := children[0]
			if child == "end" {
				var new_path []string
				for _, node := range visited {
					new_path = append(new_path, node)
				}
				new_path = append(new_path, child)
				all_possible_paths = append(all_possible_paths, new_path)
				stack[len(stack)-1] = append(children[1:])
			} else if isLower(child) && alreadyVisited(visited, child) {
				// Can't go this way
				stack[len(stack)-1] = append(children[1:])
			} else {
				visited = append(visited, child)
				stack[len(stack)-1] = append(children[1:])
				stack = append(stack, paths[child])
			}
		}
	}
	return all_possible_paths
}

func visitedSmallCaveTwice(visited []string) bool {
	node_visited := make(map[string]bool)
	for _, node := range visited {
		if isLower(node) && node_visited[node] {
			return true
		}
		node_visited[node] = true
	}
	return false
}

func B(file_path string) [][]string {
	paths := readFile(file_path)
	var all_possible_paths [][]string
	visited := []string{"start"}
	var stack [][]string
	stack = append(stack, paths["start"])
	for {
		// if nothing left to visit
		if len(stack) <= 0 {
			break
		}
		children := stack[len(stack)-1]
		if len(children) == 0 {
			// No further children to explore
			stack = append(stack[:len(stack)-1])
			// Didn't get to "end" so remove the last visited node
			if len(visited) > 0 {
				visited = append(visited[:len(visited)-1])
			}
		} else {
			child := children[0]
			if child == "end" {
				var new_path []string
				for _, node := range visited {
					new_path = append(new_path, node)
				}
				new_path = append(new_path, child)
				all_possible_paths = append(all_possible_paths, new_path)
				stack[len(stack)-1] = append(children[1:])
			} else if isLower(child) && alreadyVisited(visited, child) {
				if visitedSmallCaveTwice(visited) || child == "start" || child == "end" {
					// Can't go this way
					stack[len(stack)-1] = append(children[1:])
				} else {
					visited = append(visited, child)
					stack[len(stack)-1] = append(children[1:])
					stack = append(stack, paths[child])
				}
			} else {
				visited = append(visited, child)
				stack[len(stack)-1] = append(children[1:])
				stack = append(stack, paths[child])
			}
		}
	}
	return all_possible_paths
}

func main() {
	fmt.Println(len(B("input.txt")))
}
