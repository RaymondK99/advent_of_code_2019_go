package main

import (
	"container/list"
	"reflect"
	"strconv"
	"strings"
)

func is_valid(line string) bool {
	words := strings.Split(line," ");
	word_map := make(map[string] int);
	for _,word := range  words {
		_,ok := word_map[word]
		if !ok {
			word_map[word] = 1
		} else {
			return false
		}

	}
	return true
}

func has_anagrams(line string) bool {
	words := strings.Split(line," ");
	anagram_list := list.New()

	for _,word := range  words {
		//println(words[index])
		anagram := make([]int,256)
		// Create anagram for word
		for _,ch := range word {
			anagram[int32(ch)] += 1
		}

		for e := anagram_list.Front(); e != nil; e = e.Next() {
			if reflect.DeepEqual(e.Value, anagram) {
				return true
			}
		}

		anagram_list.PushBack(anagram)

	}
	return false
}


func day4Part1_2017(input string)  string {
	lines := strings.Split(input,"\n")
	sum := 0
	for _,line := range lines {
		if is_valid(line) {
			sum += 1
		}

	}
	return strconv.Itoa(sum);
}

func day4Part2_2017(input string)  string {
	lines := strings.Split(input,"\n")
	sum := 0
	for _,line := range lines {
		if !has_anagrams(line) {
			sum += 1
		}

	}
	return strconv.Itoa(sum);
}
