package main

import (
	"fmt"
	"testing"
)

func TestDay03_test1(t *testing.T)  {
	input := "#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"

	res := day3Part1(input)
	fmt.Println("res = ", res)
}

func TestDay03_test2(t *testing.T)  {
	input := "#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"

	res := day3Part2(input)
	fmt.Println("res = ", res)
}