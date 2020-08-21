package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func genMap(input string) map[int] []int {
	m := map[int] []int  {}

	lines := strings.Split(input, "\n")
	sort.Strings(lines)
	var activeGuard,startMin int

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		r := regexp.MustCompile(`\[(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2})\] (.*)`)
		res := r.FindStringSubmatch(lines[i])

		/*
			year, _ := strconv.Atoi(res[1])
			months, _ := strconv.Atoi(res[2])
			day, _ := strconv.Atoi(res[3])
			hour, _ := strconv.Atoi(res[4]) */
		min, _ := strconv.Atoi(res[5])
		action := res[6]

		if strings.Contains(action,"begins") {
			rAction := regexp.MustCompile(`Guard #(\d*) begins shift`)
			res := rAction.FindStringSubmatch(action)
			activeGuard,_ = strconv.Atoi(res[1])
			//fmt.Printf("     Guard no:%d\n", activeGuard)

			// Add active guard
			if _,ok := m[activeGuard]; !ok {
				m[activeGuard] = make([]int,60)
			}

		} else if strings.Contains(action,"falls asleep") {
			startMin = min
		} else if strings.Contains(action, "wakes up") {
			for i:=startMin ; i<min ; i++ {
				m[activeGuard][i] += 1
			}
		}
	}

	return m
}
func day4Part1(input string)  string {
	// Create map
	m := genMap(input)

	// Find active guard no
	maxId := 0
	maxNumMinutes := 0
	for key,value := range m {
		sum := 0
		for i := range value {
			//fmt.Println(i)
			sum += value[i]
		}
		if maxNumMinutes < sum {
			maxNumMinutes = sum
			maxId = key
		}
	}

	// Find which minutes
	maxMinute := 0
	maxMinuteId := 0
	for i := range m[maxId] {
		if m[maxId][i] > maxMinute {
			maxMinute = m[maxId][i]
			maxMinuteId = i
		}
	}

	return strconv.Itoa(maxMinuteId * maxId)
}

func day4Part2(input string)  string {
	// Find active guard no
	m := genMap(input)

	maxMinute := 0
	maxId := 0
	maxMinuteId := 0
	for key,value := range m {
		for i := range value {
			if value[i] > maxMinute {
				maxMinute = value[i]
				maxId = key
				maxMinuteId = i
			}
		}
	}

	return strconv.Itoa(maxId * maxMinuteId)
}



