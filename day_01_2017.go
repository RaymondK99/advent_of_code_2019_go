package main

import (
	"strconv"

	//"strconv"
	//"strings"
)

func day1Part1_2017(input string)  string {
	len := len(input)
	sum := 0
	for i := range input {
		curr := int(input[i] - '0');
		next := int(input[(i+1)%len] - '0');
		if curr == next {
			sum += curr;
		}
	}

	return strconv.Itoa(sum);
}

func day1Part2_2017(input string)  string {
	len := len(input)
	offset := len / 2
	sum := 0
	for i := range input {
		curr := int(input[i] - '0');
		next := int(input[(i+offset)%len] - '0');
		if curr == next {
			sum += curr;
		}
	}

	return strconv.Itoa(sum);
}
