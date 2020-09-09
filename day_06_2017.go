package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rebalance_banks(banks []int) {
	max := 0
	max_id := 0
	for i,v := range banks {
		if v > max {
			max = v
			max_id = i
		}
	}

	banks[max_id] = 0
	for i:=0;i<max;i++ {
		index := (i + max_id + 1) % len(banks)
		banks[index]++
	}
}

func banks_to_string(banks []int) string {
	var s_state string = ""

	for _, blocks := range banks {
		s_state += fmt.Sprintf("%03d,",blocks)
	}

	return s_state
}

func day6Part1_2017_solve(input string) (int,int) {
	start_state := strings.Split(input,"\t")
	banks := make([]int,16)
	states := map[string] int  {}

	// Init state vector
	for bank_id, blocks := range start_state {
		banks[bank_id],_ = strconv.Atoi(blocks)
	}

	// Add for state
	states[banks_to_string(banks)] = 0

	for iterations:=1;;iterations++ {

		fmt.Println(banks)

		// Re-balance
		rebalance_banks(banks)

		// Add to dict
		state := banks_to_string(banks)

		cycle,ok := states[state]

		if ok {
			fmt.Printf("Loop detected at iteration:%d, curr iteration:%d",cycle,iterations)
			return iterations, cycle
		} else {
			states[state] = iterations
		}

	}

	return 0,0
}

func day6Part1_2017(input string)  string {
	iterations,_ := day6Part1_2017_solve(input)
	return strconv.Itoa(iterations)
}

func day6Part2_2017(input string)  string {
	iterations,cycle := day6Part1_2017_solve(input)
	return strconv.Itoa(iterations-cycle)
}


