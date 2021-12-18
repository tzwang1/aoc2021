package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type DecodeResults struct {
	version int
	new_pos int
}

type DecodeResultsPart2 struct {
	result  int
	new_pos int
}

func ReadInput(file_name string) string {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	packets := ""
	for scanner.Scan() {
		cur_line := scanner.Text()
		packets += cur_line
	}
	return packets
}

func ConvertHexPacketsToBinary(packets string) string {
	packets_map := make(map[string]string)
	packets_map["0"] = "0000"
	packets_map["1"] = "0001"
	packets_map["2"] = "0010"
	packets_map["3"] = "0011"
	packets_map["4"] = "0100"
	packets_map["5"] = "0101"
	packets_map["6"] = "0110"
	packets_map["7"] = "0111"
	packets_map["8"] = "1000"
	packets_map["9"] = "1001"
	packets_map["A"] = "1010"
	packets_map["B"] = "1011"
	packets_map["C"] = "1100"
	packets_map["D"] = "1101"
	packets_map["E"] = "1110"
	packets_map["F"] = "1111"

	binary_packets := ""
	for _, c := range packets {
		binary_packets += packets_map[string(c)]
	}

	return binary_packets
}

func DecodePacketPart1(packet string, start_pos int) DecodeResults {
	if start_pos >= len(packet) {
		return DecodeResults{0, start_pos}
	}
	version_sum := 0
	i := start_pos

	version := packet[i : i+3]
	decimal_version, err := strconv.ParseInt(version, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert version: %s to decimal.\n", version)
	}
	version_sum += int(decimal_version)
	i += 3
	type_id := packet[i : i+3]
	i += 3
	decimal_type, err := strconv.ParseInt(type_id, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert type id %s to decimal.\n", decimal_type)
	}
	if decimal_type == 4 {
		for i < len(packet) && string(packet[i]) != "0" {
			i += 5
		}
		i += 5
	} else {
		if i < len(packet) {
			if string(packet[i]) == "0" {
				i++
				sub_packet_len := packet[i : i+15]
				decimal_sub_packet_len, err := strconv.ParseInt(sub_packet_len, 2, 64)
				if err != nil {
					log.Fatalf("Could not convert sub packet len %s to decimal.\n", sub_packet_len)
				}
				i += 15
				cur_pos := i
				for cur_pos-i != int(decimal_sub_packet_len) {
					decode_results := DecodePacketPart1(packet, cur_pos)
					version_sum += decode_results.version
					cur_pos = decode_results.new_pos
				}
				i = cur_pos
			} else {
				i++
				num_sub_packets := packet[i : i+11]
				decimal_num_sub_packets, err := strconv.ParseInt(num_sub_packets, 2, 64)
				if err != nil {
					log.Fatalf("Could not convert num sub packets %s to decimal.\n", decimal_num_sub_packets)
				}
				i += 11
				cur_num_sub_packets := 0
				cur_pos := i
				for cur_num_sub_packets < int(decimal_num_sub_packets) {
					decode_results := DecodePacketPart1(packet, cur_pos)
					version_sum += decode_results.version
					cur_num_sub_packets++
					cur_pos = decode_results.new_pos
				}
				i = cur_pos
			}
		}
	}
	return DecodeResults{version_sum, i}
}

func Part1(packets string) int {
	binary_packets := ConvertHexPacketsToBinary(packets)
	return DecodePacketPart1(binary_packets, 0).version
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AssignInitialResultValue(operation int) int {
	if operation == 0 {
		return 0
	} else if operation == 1 {
		return 1
	} else if operation == 2 {
		return math.MaxInt64
	} else if operation == 3 {
		return math.MinInt64
	} else {
		return -1
	}
}

func PerformOperation(operation int, cur_result int, decode_results DecodeResultsPart2) int {
	if decode_results.result == -1 {
		return cur_result
	}
	if cur_result == -1 {
		return decode_results.result
	}

	if operation == 0 {
		return cur_result + decode_results.result
	} else if operation == 1 {
		return cur_result * decode_results.result
	} else if operation == 2 {
		return Min(cur_result, decode_results.result)
	} else if operation == 3 {
		return Max(cur_result, decode_results.result)
	} else if operation == 5 {
		if cur_result > decode_results.result {
			return 1
		} else {
			return 0
		}
	} else if operation == 6 {
		if cur_result < decode_results.result {
			return 1
		} else {
			return 0
		}
	} else if operation == 7 {
		if cur_result == decode_results.result {
			return 1
		} else {
			return 0
		}
	} else {
		log.Fatalf("Got an invalied operation: %d\n", operation)
		return -1
	}
}

func DecodePacketPart2(packet string, start_pos int) DecodeResultsPart2 {
	if start_pos >= len(packet) {
		return DecodeResultsPart2{-1, start_pos}
	}
	result := 0
	i := start_pos

	i += 3
	type_id := packet[i : i+3]
	i += 3
	decimal_type, err := strconv.ParseInt(type_id, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert type id %d to decimal.\n", int(decimal_type))
	}
	if decimal_type == 4 {
		result_str := ""
		for i < len(packet) && string(packet[i]) != "0" {
			result_str += string(packet[i+1 : i+5])
			i += 5
		}
		result_str += string(packet[i+1 : i+5])
		result_int, err := strconv.ParseInt(result_str, 2, 64)
		if err != nil {
			log.Fatalf("Could not convert string: %s to decimal\n", result_str)
		}
		result = int(result_int)
		i += 5
	} else {
		result = AssignInitialResultValue(int(decimal_type))
		if i < len(packet) {
			if string(packet[i]) == "0" {
				i++
				sub_packet_len := packet[i : i+15]
				decimal_sub_packet_len, err := strconv.ParseInt(sub_packet_len, 2, 64)
				if err != nil {
					log.Fatalf("Could not convert sub packet len %s to decimal.\n", sub_packet_len)
				}
				i += 15
				cur_pos := i
				for cur_pos-i != int(decimal_sub_packet_len) {
					decode_results := DecodePacketPart2(packet, cur_pos)
					result = PerformOperation(int(decimal_type), result, decode_results)
					cur_pos = decode_results.new_pos
				}
				i = cur_pos
			} else {
				i++
				num_sub_packets := packet[i : i+11]
				decimal_num_sub_packets, err := strconv.ParseInt(num_sub_packets, 2, 64)
				if err != nil {
					log.Fatalf("Could not convert num sub packets %d to decimal.\n", decimal_num_sub_packets)
				}
				i += 11
				cur_num_sub_packets := 0
				cur_pos := i
				for cur_num_sub_packets < int(decimal_num_sub_packets) {
					decode_results := DecodePacketPart2(packet, cur_pos)
					result = PerformOperation(int(decimal_type), result, decode_results)
					cur_num_sub_packets++
					cur_pos = decode_results.new_pos
				}
				i = cur_pos
			}
		}
	}
	return DecodeResultsPart2{result, i}
}

func Part2(packets string) int {
	binary_packets := ConvertHexPacketsToBinary(packets)
	return DecodePacketPart2(binary_packets, 0).result
}

func main() {
	packets := ReadInput("input.txt")
	fmt.Println(Part1(packets))
	fmt.Println(Part2(packets))
}

/*
Part1
tiny_input = 6
small_input = 9
small_input1 = 14

example_input1 = 16
example_input2 = 12
example_input3 = 23
example_input4 = 31

Part2
*/
