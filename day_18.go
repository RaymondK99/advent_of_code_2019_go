package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

type Pos struct {
	x,y int
}

// An Node is something we manage in a distance queue.
type Node struct {
	pos Pos
	keys string
 	distance int    // The distance of the item in the queue.
	index    int    // The index of the item in the heap.
}

type DistanceElement struct {
	keys string
	pos Pos
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortString(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, distance so we use greater than here.
	if pq[i].distance == pq[j].distance {
		return len(pq[i].keys) > len(pq[j].keys)
	}
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func isKey(ch uint8) bool {
	return ch >= 'a' && ch <= 'z'
}

func isDoor(ch uint8) bool {
	return ch >= 'A' && ch <= 'Z'
}

func day18Part1(input string) string {

	var start_pos Pos
	num_keys := 0

	// Build map
	lines := strings.Split(input,"\n")
	m := make(map[Pos] uint8)
	for y:=0;y<len(lines);y++ {
		line := lines[y]
		//fmt.Println("Eval line:",line)
		for x:=0;x<len(line);x++ {
			//fmt.Println("Insert element:",line[x]," at:",x,y,"iskey:",unicode.IsLower(rune(line[x])))
			m[Pos{x:x,y:y}] = line[x]
			if line[x] == '@' {
				start_pos = Pos{x:x,y:y}
			} else if isKey(line[x]) {
				num_keys++
			}

		}
	}

	fmt.Println("start node:",start_pos,", num keys=",num_keys)

	// Create first node
	start_node := Node{pos:start_pos, distance: 0,keys:""}

	// Queue
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &start_node)

	// Distance map
	var distances = map[DistanceElement] int {}
	distances[DistanceElement{pos: start_pos, keys: ""}] = 0


	for ;len(pq)>0; {

		// Pop next node to evaluate
		//heap.Init(&pq)
		node := heap.Pop(&pq).(*Node)
		//fmt.Println("Pop node:",*node)
		x := node.pos.x
		y := node.pos.y
		dist := node.distance
		next_dist := dist+1

		// Found solution
		if len(node.keys) == num_keys {
			// Found solution
			fmt.Println("Found keys:",len(node.keys)," num_keys:",num_keys)
			fmt.Println("Found solution at dist:", node.distance)
			break
		}

		// Generate neighbors
		adjacent_pos := [4]Pos {Pos{x:x,y:y-1},Pos{x:x,y:y+1},Pos{x:x-1,y:y},Pos{x:x+1,y:y}}
		adjacent_nodes := make([]Node,0)

		for _,next_pos := range adjacent_pos {
			next_ch := m[next_pos]
			//fmt.Println("Eval pos:",next_pos, ", item:",string(next_ch))

			// Is passable position
			if next_ch == '.' || next_ch == '@' || ( isDoor(next_ch) && strings.Contains(node.keys, strings.ToLower(string(next_ch))) ) {
				adjacent_nodes = append(adjacent_nodes, Node{keys: node.keys,distance: next_dist,pos: next_pos})
			} else if next_ch == '#' {
				// Not passable
				continue
			} else if isKey(next_ch) {
				// Collect key if needed

				if !strings.Contains(node.keys, string(next_ch)) {
					next_keys := sortString(node.keys + string(next_ch))
					adjacent_nodes = append(adjacent_nodes, Node{keys: next_keys, distance: next_dist, pos: next_pos})
				} else {
					adjacent_nodes = append(adjacent_nodes, Node{keys: node.keys,distance: next_dist,pos: next_pos})
				}
			}
		}

		for _,next_node := range adjacent_nodes {

			next_dist_elem := DistanceElement{keys: next_node.keys, pos: next_node.pos}
			prev_dist,ok := distances[next_dist_elem]

			if !ok || ( ok && prev_dist > next_dist ) {
				//fmt.Println("Push new node:",next_node)
				distances[next_dist_elem] = next_dist
				heap.Push(&pq,&Node{keys: next_node.keys,pos:next_node.pos,distance: next_node.distance})
			} else {
				//fmt.Println("Node is visited:",next_node)
			}


		}


	}


	return ""
}
