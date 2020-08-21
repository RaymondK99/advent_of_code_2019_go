package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() string {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	return input
}

func main() {

	if len(os.Args) < 3 {
		os.Exit(1)
	}

	// Read input arguments
	argsWithoutProg := os.Args[1:]
	day,_ := strconv.Atoi(argsWithoutProg[0])
	part,_:= strconv.Atoi(argsWithoutProg[1])

	// Read input data
	input := readInput()

	var res string
	switch day {
	case 1:
		if part == 1 {
			res = day1Part1(input)
		} else if part == 2 {
			res = day1Part2(input)
		}
		break;
	case 2:
		if part == 1 {
			res = day2Part1(input)
		} else if part == 2 {
			res = day2Part2(input)
		}
		break;
	case 3:
		if part == 1 {
			res = day3Part1(input)
		} else if part == 2 {
			res = day3Part2(input)
		}
		break;
	case 4:
		if part == 1 {
			res = day4Part1(input)
		} else if part == 2 {
			res = day4Part2(input)
		}
		break;
	case 18:
		if part == 1 {
			res = day18Part1(input)
		} else if part == 2 {
			res = day18Part1(input)
		}
		break;
	default:
		res = "Not impl"
		break;

	}

	fmt.Println(res)
}



