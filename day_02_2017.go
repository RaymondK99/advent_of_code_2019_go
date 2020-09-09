package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day2Part1_2017(input string)  string {
	lines := strings.Split(input,"\n");
	check_sum := 0

	for _,line := range lines {
		var max,min int = 0,math.MaxInt32

		for _,strValue := range strings.Split(line,"\t") {
			value, _ := strconv.Atoi(strValue)
			if value > max {
				max = value
			}
			if value < min {
				min = value
			}
		}
		check_sum += max - min
	}
	return strconv.Itoa(check_sum);
}

func day2Part2_2017(input string)  string {
	lines := strings.Split(input,"\n");
	sum := 0

	for _,line := range lines {
		nums := make([]int,0)
		for _,strValue := range strings.Split(line,"\t") {
			value, _ := strconv.Atoi(strValue)
			nums = append(nums, value)
		}

		fmt.Println("Eval line:",nums)
		for i:=0;i<len(nums);i++ {
			for j:=i+1;j<len(nums);j++ {
				a := nums[i]
				b := nums[j]
				if a > b && a % b == 0 {
					sum += a / b
					fmt.Println("Add ", a/b)
				} else if b > a && b % a == 0 {
					sum += b / a
					fmt.Println("Add ", b/a)
				}
			}
		}
	}
	return strconv.Itoa(sum);
}
