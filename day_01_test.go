package main

import (
	"fmt"
	"testing"
)

func TestMyTest1(t *testing.T)  {
	input := "-1\n+2\n"
	res := day1Part1(input)
	fmt.Println("res = ", res)
}


func TestMyTest2(t *testing.T)  {
	input := "+1\n"+
			"-2\n" +
			"+3\n" +
			"+1\n" +
			"+1\n" +
			"-2\n"


	res := day1Part2(input)
	fmt.Println("res = ", res)
}
