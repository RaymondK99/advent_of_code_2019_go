package main

import (
	"fmt"
	"strconv"
	"strings"
)


func parse_tower(input string)  (string,int) {
	var height int
	var name string
	fmt.Sscanf(strings.TrimSpace(input),"%s (%d)",&name,&height)

	return name,height
}

type Tower struct {
	name string
	height int
	childs[] string
	parent string
}

type Pair struct {
	a string
	b string
}

func day7Part1_2017(input string)  string {
	lines := strings.Split(input,"\n");
	tower_list := make([]string,0)
	parent_list := make(map[string] string)
	for _,line := range lines {
		cols := strings.Split(line,"->")
		tower := strings.TrimSpace(cols[0])

		name,height := parse_tower(tower)

		tower_list = append(tower_list, strings.TrimSpace(name))

		fmt.Println(name,height)
		if len(cols) == 2 {
			childs := strings.TrimSpace(cols[1])
			for _,child := range strings.Split(childs,",") {
				//println("   child_name:",strings.TrimSpace(child))
				parent_list[strings.TrimSpace(child)] = strings.TrimSpace(name)
			}

		}
	}

	//fmt.Println(tower_list)
	//fmt.Println(parent_list)

	element := tower_list[0]
	found := false

	for ;found == false; {
		fmt.Println("query:",element);
		parent,ok := parent_list[element]
		if !ok {
			// Found node without parent
			fmt.Println("Node ",element," has no parent")
			return element
			break
		} else {
			// Has parent
			fmt.Println("Node ",element," has ",parent," as parent")
			element = parent
		}
	}


	return "No solution"
}

func day7Part2_2017(input string)  string {
	return strconv.Itoa(1)
}


