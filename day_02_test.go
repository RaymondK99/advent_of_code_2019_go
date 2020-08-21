package main

import (
	"fmt"
	"testing"
)

func TestDay02_test1(t *testing.T)  {
	input := "abcdef\n"+
		"bababc\n" +
		"abbcde\n" +
		"abcccd\n" +
		"aabcdd\n" +
		"abcdee\n" +
		"ababab"

	res := day2Part1(input)
	fmt.Println("res = ", res)
}

func TestDay02_test2(t *testing.T)  {
	input := "abcdef\n"+
		"bababc\n" +
		"abbcde\n" +
		"abcccd\n" +
		"aabcdd\n" +
		"abcdee\n" +
		"ababab"

	res := day2Part2(input)
	fmt.Println("res = ", res)
}

