package main

import (
	"strconv"
	"strings"
)

func day5Part1_2017(input string)  string {
	lines := strings.Split(input,"\n")
	list := make([]int,0)

	for _,value := range lines {
		int_value,_ := strconv.Atoi(value)
		list = append(list, int_value)
	}

	offset := 0
	steps := 0
	for ; offset >= 0 && offset < len(list); {
		//fmt.Println("Current offset =",offset,", value=",list[offset])
		old_offset := offset
		offset += list[offset]
		// Increment old offset by 1
		list[old_offset] += 1
		steps += 1
	}

	return strconv.Itoa(steps);
}

func day5Part2_2017(input string)  string {
	lines := strings.Split(input,"\n")
	list := make([]int,0)

	for _,value := range lines {
		int_value,_ := strconv.Atoi(value)
		list = append(list, int_value)
	}

	offset := 0
	steps := 0
	for ; offset >= 0 && offset < len(list); {
		//fmt.Println("Current offset =",offset,", value=",list[offset])
		old_offset := offset
		offset += list[offset]
		// Increment old offset by 1
		if list[old_offset] >= 3 {
			list[old_offset] -= 1
		} else {
			list[old_offset] += 1
		}
		steps += 1
	}

	return strconv.Itoa(steps);
}
