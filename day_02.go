package main

import (
	"strconv"
	"strings"
)

func day2Part1(input string)  string {
	lines := strings.Split(input,"\n")
	twosTotal := 0
	threesTotal := 0

	for i:=0;i< len(lines);i++ {
		line := lines[i]
		twos := 0
		threes := 0

		var count [256]int

		for j:=0; j< len(line);j++ {
			ch := line[j] - 'a'
			count[ch] += 1
		}

		for j:=0;j<len(count);j++ {
			if count[j] == 2 {
				twos++
			} else if count[j] == 3 {
				threes++
			}
		}

		if twos > 0 {
			twosTotal++
		}

		if threes > 0 {
			threesTotal++
		}

	}

	return strconv.Itoa(threesTotal * twosTotal)
}


func commonString(str1 string, str2 string) (int,string) {
	res := ""

	for i:=0;i<len(str1) && i < len(str2);i++ {
		if str1[i] == str2[i] {
			res += string(str1[i])
		}
	}

	missing := len(str1) - len(res)
	return missing, res
}

func day2Part2(input string)  string {
	lines := strings.Split(input,"\n")

	for i:=0;i< len(lines);i++ {
		for j:=0;j< len(lines);j++ {
			if i == j {
				continue
			}

			missing, subStr := commonString(lines[i], lines[j])

			if missing == 1 {
				return subStr
			}

		}

	}

	return ""
}

