package main

import (
	"strconv"
	"strings"
)

func day1Part1(input string)  string {
	list := strings.Split(input,"\n");
	sum := 0
	for i:= 0; i<len(list);i++ {
		if len(list[i]) == 0 {
			continue
		}
		value, _ := strconv.Atoi(list[i])
		sum += value
	}

	return strconv.Itoa(sum)
}

func day1Part2(input string)  string {
	list := strings.Split(input,"\n");
	freq := map[int] int {}
	sum := 0
	i := 0
	for ;; {
		if len(list[i]) > 0 {
			value, _ := strconv.Atoi(list[i])
			sum += value

			if _, ok := freq[sum]; ok {
				// Exists already
				break
			} else {
				freq[sum] = 1
			}
		}

		// Increment counter
		i += 1
		if i == len(list) {
			i = 0
		}
	}

	return strconv.Itoa(sum)
}
