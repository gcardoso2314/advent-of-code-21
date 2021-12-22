package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(path string) string {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		return parseHexadecimal(scanner.Text())
	}
	return ""
}

var bin_map = map[string]string{
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func parseHexadecimal(hex string) string {
	var output string
	for _, c := range hex {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			output += bin_map[string(c)]
		} else {
			bin_num := strconv.FormatInt(int64(num), 2)
			// Pad with leading zeros to make 4 digits
			for i := 0; i < 4-len(bin_num); i++ {
				output += "0"
			}
			output += bin_num
		}
	}
	return output
}

func versionNumber(bin_string string) int64 {
	num, _ := strconv.ParseInt(bin_string[:3], 2, 64)
	return num
}

func sum(values []int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}

func product(values []int) int {
	product := values[0]
	for i := 1; i < len(values); i++ {
		product *= values[i]
	}
	return product
}

func min(values []int) int {
	min_val := values[0]
	for _, v := range values {
		if v < min_val {
			min_val = v
		}
	}
	return min_val
}

func max(values []int) int {
	max_val := values[0]
	for _, v := range values {
		if v > max_val {
			max_val = v
		}
	}
	return max_val
}

func greaterThan(values []int) int {
	if values[0] > values[1] {
		return 1
	} else {
		return 0
	}
}

func lessThan(values []int) int {
	if values[0] < values[1] {
		return 1
	} else {
		return 0
	}
}

func equalTo(values []int) int {
	if values[0] == values[1] {
		return 1
	} else {
		return 0
	}
}

func parsePacket(bin_string string) (int64, int, int) {
	version_num, _ := strconv.ParseInt(bin_string[:3], 2, 64)
	type_id, _ := strconv.ParseInt(bin_string[3:6], 2, 64)

	fmt.Println("Packet contains version_num=", version_num)
	var reader_point int
	var value int
	if type_id == 4 {
		// Literal value
		fmt.Println("It's a literal value packet")
		i := 0
		var literal_binary string
		for {
			next_group := bin_string[6+(5*i) : 6+(5*(i+1))]

			literal_binary += next_group[1:]
			if string(next_group[0]) == "0" {
				break
			}
			i++
		}
		literal_value, _ := strconv.ParseInt(literal_binary, 2, 64)
		fmt.Println(literal_value)
		reader_point = 6 + (5 * (i + 1))
		return version_num, reader_point, int(literal_value)
	} else {
		var literal_values []int
		length_type_id := string(bin_string[6])
		if length_type_id == "0" {
			// Next 15 bits contain total length of subpackets
			length, _ := strconv.ParseInt(bin_string[7:22], 2, 64)
			fmt.Println("It's a packet that contains", length, "bits of subpackets")
			reader_point += 22
			for {
				if reader_point == 22+int(length) {
					break
				}
				child_version_num, chars_seen, value := parsePacket(bin_string[reader_point : 22+length])
				version_num += child_version_num
				reader_point += chars_seen
				literal_values = append(literal_values, value)
			}

		} else {
			// Next 11 bits tell you the number of packets contained
			n_packets, _ := strconv.ParseInt(bin_string[7:18], 2, 64)
			fmt.Println("It's a packet that contains", n_packets, "subpackets")
			reader_point = 18
			for i := 0; i < int(n_packets); i++ {
				child_version_num, chars_seen, value := parsePacket(bin_string[reader_point:])
				version_num += child_version_num
				reader_point += chars_seen
				literal_values = append(literal_values, value)
			}
		}
		switch type_id {
		case 0:
			value = sum(literal_values)
		case 1:
			value = product(literal_values)
		case 2:
			value = min(literal_values)
		case 3:
			value = max(literal_values)
		case 5:
			value = greaterThan(literal_values)
		case 6:
			value = lessThan(literal_values)
		case 7:
			value = equalTo(literal_values)
		}
	}
	return version_num, reader_point, value
}

func main() {
	bin_string := readFile("input.txt")
	// bin_string := parseHexadecimal("9C0141080250320F1802104A08")
	fmt.Println(parsePacket(bin_string))
}
