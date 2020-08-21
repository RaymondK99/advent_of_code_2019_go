package main

import (
	"fmt"
	"testing"
)

func TestDay18_test1(t *testing.T)  {
	input := "#########\n#b.A.@.a#\n#########"

	res := day18Part1(input)
	fmt.Println("res = ", res)
}

func TestDay18_test2(t *testing.T)  {
	input := "########################\n" +
	"#f.D.E.e.C.b.A.@.a.B.c.#\n" +
		"######################.#\n" +
		"#d.....................#\n" +
		"########################"

	res := day18Part1(input)
	fmt.Println("res = ", res)
}

func TestDay18_test3(t *testing.T)  {
	input := "#################\n#i.G..c...e..H.p#\n########.########\n#j.A..b...f..D.o#\n########@########\n#k.E..a...g..B.n#\n########.########\n#l.F..d...h..C.m#\n#################"

	res := day18Part1(input)
	fmt.Println("res = ", res)
}




