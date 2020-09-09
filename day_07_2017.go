package main

import (
	"fmt"
	"strconv"
	"strings"
)


func parse_tower(input string)  (string,int) {
	var c1,c2,c3,c4,height int
	fmt.Sscanf(strings.TrimSpace(input),"%c%c%c%c (%d)",&c1,&c2,&c3,&c4,&height)
	name := fmt.Sprintf("%c%c%c%c",c1,c2,c3,c4)

	return name,height
}

type Tower struct {
	name string
	height int
	childs[] string
	parent string
}

func day7Part1_2017(input string)  string {
	lines := strings.Split(input,"\n");
	for _,line := range lines {
		cols := strings.Split(line,"->");
		tower := strings.TrimSpace(cols[0])

		name,height := parse_tower(tower)

		fmt.Println(name,height)
		if len(cols) == 2 {
			childs := strings.TrimSpace(cols[1])
			for _,child := range strings.Split(childs,",") {
				println("   child_name:",strings.TrimSpace(child))
			}

		}
	}
	return strconv.Itoa(1)
}

func day7Part2_2017(input string)  string {
	return strconv.Itoa(1)
}


