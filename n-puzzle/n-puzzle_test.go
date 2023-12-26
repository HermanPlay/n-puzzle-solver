package npuzzle

import (
	"fmt"
	"math"
	"n-puzzle/utils"
	"testing"
)

func TestN3Solvable1(t *testing.T) {
	i := 1
	fileName := fmt.Sprintf("tests/test%v.txt", i)
	s, n, err := utils.ReadFileToArray(fileName)
	if err != nil {
		t.Fatalf("Testfile %v failed: %v", i, err)
	}
	n = int(math.Sqrt(float64(n)))
	fmt.Printf("%v n: %v\n", s, n)
	node := CreateNode(s, n, 0)
	node.SetFValue()
	result := AStarSearch(node, 300)
	if !IsGoal(result.State, result.N) {
		t.Fatalf("Search found wrong result: %v", result.State)
	}
}

func TestN3Solvable2(t *testing.T) {
	i := 2
	fileName := fmt.Sprintf("tests/test%v.txt", i)
	s, n, err := utils.ReadFileToArray(fileName)
	if err != nil {
		t.Fatalf("Testfile %v failed: %v", i, err)
	}
	n = int(math.Sqrt(float64(n)))
	fmt.Printf("%v n: %v\n", s, n)
	node := CreateNode(s, n, 0)
	node.SetFValue()
	result := AStarSearch(node, 300)
	if !IsGoal(result.State, result.N) {
		t.Fatalf("Search found wrong result: %v", result.State)
	}
}

func TestN3Solvable3(t *testing.T) {
	i := 3
	fileName := fmt.Sprintf("tests/test%v.txt", i)
	s, n, err := utils.ReadFileToArray(fileName)
	if err != nil {
		t.Fatalf("Testfile %v failed: %v", i, err)
	}
	n = int(math.Sqrt(float64(n)))
	fmt.Printf("%v n: %v\n", s, n)
	node := CreateNode(s, n, 0)
	node.SetFValue()
	result := AStarSearch(node, 300)
	if !IsGoal(result.State, result.N) {
		t.Fatalf("Search found wrong result: %v", result.State)
	}
}

func TestN3Solvable4(t *testing.T) {
	i := 4
	fileName := fmt.Sprintf("tests/test%v.txt", i)
	s, n, err := utils.ReadFileToArray(fileName)
	if err != nil {
		t.Fatalf("Testfile %v failed: %v", i, err)
	}
	n = int(math.Sqrt(float64(n)))
	fmt.Printf("%v n: %v\n", s, n)
	node := CreateNode(s, n, 0)
	node.SetFValue()
	result := AStarSearch(node, 300)
	if !IsGoal(result.State, result.N) {
		t.Fatalf("Search found wrong result: %v", result.State)
	}
}

func TestN3Solvable5(t *testing.T) {
	i := 5
	fileName := fmt.Sprintf("tests/test%v.txt", i)
	s, n, err := utils.ReadFileToArray(fileName)
	if err != nil {
		t.Fatalf("Testfile %v failed: %v", i, err)
	}
	n = int(math.Sqrt(float64(n)))
	fmt.Printf("%v n: %v\n", s, n)
	node := CreateNode(s, n, 0)
	node.SetFValue()
	result := AStarSearch(node, 300)
	if !IsGoal(result.State, result.N) {
		t.Fatalf("Search found wrong result: %v", result.State)
	}
}

func TestN3Solvable6(t *testing.T) {
	i := 6
	fileName := fmt.Sprintf("tests/test%v.txt", i)
	s, n, err := utils.ReadFileToArray(fileName)
	if err != nil {
		t.Fatalf("Testfile %v failed: %v", i, err)
	}
	n = int(math.Sqrt(float64(n)))
	fmt.Printf("%v n: %v\n", s, n)
	node := CreateNode(s, n, 0)
	node.SetFValue()
	result := AStarSearch(node, 300)
	if !IsGoal(result.State, result.N) {
		t.Fatalf("Search found wrong result: %v", result.State)
	}
}
