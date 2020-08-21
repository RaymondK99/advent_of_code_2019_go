package main

import (
	"strconv"
	"strings"
)

type Point struct {
	x,y int
}

func buildMap(input string) map [Point] []int {
	lines := strings.Split(input, "\n")
	m := map [ Point] []int  {}

	for i:=0;i< len(lines);i++ {
		// #1 @ 1,3: 4x4
		if len(lines[i]) == 0 {
			continue
		}
		cols := strings.Split(lines[i]," ")
		pos0 := strings.Split(cols[0], "#")
		pos1 := strings.Split(cols[2],",")
		pos2 := strings.Split(cols[3],"x")
		num,_ := strconv.Atoi(pos0[1])
		x,_ := strconv.Atoi(pos1[0])
		y,_ := strconv.Atoi(strings.Replace(pos1[1],":","",1))

		width,_ := strconv.Atoi(pos2[0])
		height,_ := strconv.Atoi(pos2[1])

		y_end := y + height
		x_end := x + width
		x_start := x

		for ;y<y_end;y++ {
			x = x_start
			for ; x < x_end; x++ {
				point := Point{x:x, y:y}
				if _, ok := m[point]; ok {
					m[point] = append(m[point],num)

				} else {
					l := make([]int,1)
					l[0] = num
					m[point] = l
				}
			}
		}
	}

	return m
}

func day3Part1(input string)  string {
	m := buildMap(input)

	sum := 0
	for point := range m {
		l,_ := m[point]

		if len(l) > 1 {
			sum++
		}
	}

	return strconv.Itoa(sum)
}

func day3Part2(input string)  string {
	m := buildMap(input)
	table := make([]int, len(strings.Split(input,"\n"))+1)

	for point := range m {
		l,_ := m[point]

		if len(l) > 1 {
			for i:=0;i<len(l);i++ {
				index := l[i]
				table[index] += 1
			}
		}
	}

	for i:=1;i< len(table);i++ {
		if table[i] == 0 {
			return strconv.Itoa(i)
		}
	}

	return strconv.Itoa(0)
}
