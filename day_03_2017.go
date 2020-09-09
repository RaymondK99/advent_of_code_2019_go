package main

import (
	"math"
	"strconv"
)

func day3Part1_2017(input string)  string {
	num,_ := strconv.Atoi(input)

	ring_number := int32(math.Sqrt(float64(num)))
	abs_dist := ring_number - 1
	println("ring_number=",ring_number,", abs_dist=",abs_dist)
	


	return strconv.Itoa(0);
}
