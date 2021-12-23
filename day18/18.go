package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SnailNum struct {
	index int
	data  int
	left  *SnailNum
	right *SnailNum
}

func (root *SnailNum) FillIndex(start_index int) int {
	curr_index := start_index
	if root.left != nil {
		curr_index = root.left.FillIndex(curr_index)
		curr_index = root.right.FillIndex(curr_index)
	} else {
		root.index = curr_index
		curr_index++
	}
	return curr_index
}

func (root *SnailNum) PrintNum() string {
	var num_rep string
	if root.left != nil {
		num_rep += "["
		num_rep += root.left.PrintNum()
		num_rep += ","
		num_rep += root.right.PrintNum()
		num_rep += "]"
	} else {
		num_rep += fmt.Sprint(root.data)
	}
	return num_rep
}

func (root *SnailNum) SumSnailNum(num SnailNum) *SnailNum {
	return &SnailNum{left: root, right: &num}
}

func (root *SnailNum) Add(num, index int) {
	if index == 0 {
		// No numbers with index 0 by design
		return
	} else if root.index == index && root.left == nil {
		root.data += num
	} else if root.left != nil {
		root.left.Add(num, index)
		root.right.Add(num, index)
	}
}

func (root *SnailNum) Reduce() {
	for {
		root.FillIndex(1) // Create new indices to know where to add any exploding numbers
		left_index, right_index, left_num, right_num := root.ReduceExplode()
		if left_index > 0 && right_index > 0 {
			root.Add(left_num, left_index-1)
			root.Add(right_num, right_index+1)
			continue
		}
		has_split := root.ReduceSplit()
		if !has_split {
			break
		}
	}
}

func (root *SnailNum) ReduceSplit() bool {
	if root.data >= 10 {
		root.left = &SnailNum{data: root.data / 2}
		root.right = &SnailNum{data: root.data - root.data/2}
		root.data = 0
		return true
	} else if root.left != nil {
		has_split := root.left.ReduceSplit()
		if !has_split {
			has_split = root.right.ReduceSplit()
		}
		return has_split
	}
	return false
}

func (root *SnailNum) ReduceExplode(level_opt ...int) (int, int, int, int) {
	level := 0
	if len(level_opt) > 0 {
		level = level_opt[0]
	}

	if level < 3 && root.left != nil {
		left_index, right_index, left_num, right_num := root.left.ReduceExplode(level + 1)
		if left_index > 0 && right_index > 0 {
			return left_index, right_index, left_num, right_num
		}
		left_index, right_index, left_num, right_num = root.right.ReduceExplode(level + 1)
		return left_index, right_index, left_num, right_num
	} else if level == 3 && root.left != nil {
		switch {
		case root.left.left != nil:
			if root.left.left.index > 0 && root.left.right.index > 0 {
				// This is a pair nested in 4 pairs
				left_index := root.left.left.index
				right_index := root.left.right.index
				left_num := root.left.left.data
				right_num := root.left.right.data
				root.left.left = nil
				root.left.right = nil
				root.left.data = 0
				return left_index, right_index, left_num, right_num
			}
		case root.right.left != nil:
			if root.right.left.index > 0 && root.right.right.index > 0 {
				// This is a pair nested in 4 pairs
				left_index := root.right.left.index
				right_index := root.right.right.index
				left_num := root.right.left.data
				right_num := root.right.right.data
				root.right.left = nil
				root.right.right = nil
				root.right.data = 0
				return left_index, right_index, left_num, right_num
			}
		}
	}
	return 0, 0, 0, 0
}

func (root *SnailNum) Magnitude() int {
	var magnitude int
	if root.left != nil {
		magnitude += 3*root.left.Magnitude() + 2*root.right.Magnitude()
	} else {
		magnitude += root.data
	}
	return magnitude
}

func parseSnailNum(num string) *SnailNum {
	start_char := string(num[0])
	if start_char == "[" {
		// It's a new pair
		unclosed_brackets := 0
		for i, c := range num[1:] {
			if string(c) == "," && unclosed_brackets == 0 {
				return &SnailNum{
					left:  parseSnailNum(num[1 : 1+i]),
					right: parseSnailNum(num[i+2 : len(num)-1]),
				}
			} else if string(c) == "[" {
				unclosed_brackets++
			} else if string(c) == "]" {
				unclosed_brackets--
			}
		}
	} else {
		// It's a number
		value, _ := strconv.Atoi(num)
		return &SnailNum{data: value}
	}
	return &SnailNum{}
}

func A(path string) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var snail_num *SnailNum
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			snail_num = parseSnailNum(scanner.Text())
		} else {
			new_snail_num := parseSnailNum(scanner.Text())
			snail_num = snail_num.SumSnailNum(*new_snail_num)
			snail_num.Reduce()
		}
	}
	fmt.Println(snail_num.PrintNum())
	fmt.Println(snail_num.Magnitude())
}

func B(path string) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var all_nums []string
	for i := 0; scanner.Scan(); i++ {
		all_nums = append(all_nums, scanner.Text())
	}

	var max_magnitude int
	for i := 0; i < len(all_nums)-1; i++ {
		for j := 1; j < len(all_nums); j++ {
			// This is so inefficient... not sure how to work with pointers so creating new structs
			// every iteration
			num_1 := parseSnailNum(all_nums[i])
			num_2 := parseSnailNum(all_nums[j])
			sum_1 := num_1.SumSnailNum(*num_2)
			sum_1.Reduce()
			if sum_1.Magnitude() > max_magnitude {
				max_magnitude = sum_1.Magnitude()
			}

			num_1 = parseSnailNum(all_nums[i])
			num_2 = parseSnailNum(all_nums[j])
			sum_2 := num_2.SumSnailNum(*num_1)
			sum_2.Reduce()

			if sum_2.Magnitude() > max_magnitude {
				max_magnitude = sum_2.Magnitude()
			}
		}
	}
	fmt.Println(max_magnitude)

}

func main() {
	A("input.txt")
	B("input.txt")
}
