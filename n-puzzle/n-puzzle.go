package npuzzle

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"time"
)

type Node struct {
	State  []int
	parent *Node
	N      int
	gValue int // Cost from start state
	hValue int // Heuristic estimate to goal state
	fValue int // f(n) = g(n) + h(n)
}

var verbose bool = false

var Logger *log.Logger = log.Default()

type PriorityQueue []*Node

func CreateNode(state []int, n, g int) *Node {
	node := Node{state, nil, n, g, -1, -1}
	// node.hValue = node.GetManhattanDistance()
	// node.fValue = node.gValue + node.hValue
	return &node
}

func (node Node) String() string {
	// Convert array to a string representation
	arrayString := ""
	for i := 0; i < len(node.State); i++ {
		arrayString += fmt.Sprintf("%d ", node.State[i])
	}
	return arrayString
}

func (node *Node) SetFValue() {
	node.hValue = node.GetManhattanDistance()
	node.fValue = node.gValue + node.hValue
}

func (node Node) GetNeighbours() [][]int {
	index_0 := 0
	for i, val := range node.State {
		if val == 0 {
			index_0 = i
			break
		}
	}

	empty_tile_x := index_0 % node.N
	empty_tile_y := index_0 / node.N

	res := make([][]int, 0)
	// Move left block
	if empty_tile_x > 0 {
		new := make([]int, len(node.State))
		copy(new, node.State)
		new[index_0], new[index_0-1] = new[index_0-1], new[index_0]
		res = append(res, new)
	}
	// Move right block
	if empty_tile_x < node.N-1 {
		new := make([]int, len(node.State))
		copy(new, node.State)
		new[index_0], new[index_0+1] = new[index_0+1], new[index_0]
		res = append(res, new)
	}

	if empty_tile_y > 0 {
		new := make([]int, len(node.State))
		copy(new, node.State)
		new[index_0], new[index_0-node.N] = new[index_0-node.N], new[index_0]
		res = append(res, new)
	}

	if empty_tile_y < node.N-1 {
		new := make([]int, len(node.State))
		copy(new, node.State)
		new[index_0], new[index_0+node.N] = new[index_0+node.N], new[index_0]
		res = append(res, new)
	}

	return res
}

func (node Node) PrettyPrint() {
	fmt.Printf("----------------")
	for i, val := range node.State {
		if i%node.N == 0 {
			fmt.Println()
		}
		fmt.Printf("%v ", val)
	}
	fmt.Printf("\n----------------\n")
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
	heap.Fix(pq, len(*pq)-1)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].fValue < pq[j].fValue
}

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Pop() interface{} {
	last := len(*pq) - 1
	node := (*pq)[last]
	*pq = (*pq)[:last]
	return node
}

func (pq *PriorityQueue) Print() {
	fmt.Println("Current open_set:")
	for _, val := range *pq {
		fmt.Printf("%v ", val.fValue)
	}
	fmt.Println()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Peek() {
	count := 10
	i := 0
	for i < count {
		fmt.Println(pq[i].fValue)
		i++
	}
}

func (node *Node) GetManhattanDistance() int {
	total := 0

	for i, val := range node.State {
		if val == 0 {
			continue
		}

		curr_x := i % node.N
		curr_y := i / node.N
		correct_x := (val - 1) % node.N
		correct_y := (val - 1) / node.N

		v := int(math.Abs(float64(correct_x)-float64(curr_x)) + math.Abs(float64(correct_y)-float64(curr_y)))
		if v == 0 {
			continue
		}

		total += v

		// Calculate linear conflicts
		if curr_x == correct_x {
			total += GetRowLinearConflicts(node.State, node.N, curr_x, i, val)
		}

		if curr_y == correct_y {
			total += GetColumnLinearConflicts(node.State, node.N, curr_y, i, val)
		}
	}
	return total
}

func GetColumnLinearConflicts(state []int, n, curr_y, index, val int) int {
	conflicts := 0
	for i := index; i < (n-curr_y)*n+index; i += n {
		if val > state[i] {
			conflicts++
		}
	}

	return conflicts * 2
}

func GetRowLinearConflicts(state []int, n, curr_x, index, val int) int {
	conflicts := 0
	for i := index; i < n-curr_x+index; i++ {
		if val > state[i] {
			conflicts++
		}
	}
	return conflicts * 2
}

func IsGoal(state []int, n int) bool {
	for i := 1; i < n*n; i++ {
		if state[i-1] != i {
			return false
		}
	}
	return true
}

func ReconstructPath(endNode *Node) []string {
	temp := endNode
	result := make([]string, 0)
	for temp != nil {
		result = append(result, temp.String())
		temp = temp.parent
	}
	return result
}

type OpenSet struct {
	gValue int
	hValue int
	parent *Node
}

func AStarSearch(startNode *Node, maxDepth int) *Node {
	open_set := PriorityQueue{}
	in_open_set := make(map[string]OpenSet)
	closed_set := make(map[string]bool)
	open_set.Push(startNode)
	counter := 0
	if verbose {
		Logger.Println("Starting evaluation!")
	}
	startTime := time.Now()

	for open_set.Len() > 0 {
		currentNode := heap.Pop(&open_set).(*Node)
		open_set_values := in_open_set[currentNode.String()]
		currentNode.gValue = open_set_values.gValue
		currentNode.hValue = open_set_values.hValue
		currentNode.parent = open_set_values.parent

		delete(in_open_set, currentNode.String())

		counter++

		if IsGoal(currentNode.State, currentNode.N) {
			fmt.Printf("Found solution in: %v\n", time.Since(startTime))
			fmt.Printf("Evaluated total: %v nodes\n", counter)
			return currentNode
		}

		closed_set[currentNode.String()] = true

		for _, neighborState := range currentNode.GetNeighbours() {
			neighbor := CreateNode(neighborState, currentNode.N, currentNode.gValue+1)
			if neighbor.gValue > maxDepth {
				continue
			}
			neighbor.parent = currentNode

			if _, exists := closed_set[neighbor.String()]; exists {
				continue
			}
			neighbor.SetFValue()

			if in_open_set[neighbor.String()].gValue != 0 {
				if neighbor.fValue < in_open_set[neighbor.String()].gValue+in_open_set[neighbor.String()].hValue {
					in_open_set[neighbor.String()] = OpenSet{neighbor.gValue, neighbor.hValue, currentNode}
				}
			} else {
				in_open_set[neighbor.String()] = OpenSet{neighbor.gValue, neighbor.hValue, currentNode}
				open_set.Push(neighbor)
			}
		}
	}
	if verbose {
		Logger.Printf("Didnt find solution in: %v\n", time.Since(startTime))
		Logger.Printf("Evaluated total: %v nodes\n", counter)
	}
	return startNode
}
