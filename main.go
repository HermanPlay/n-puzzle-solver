package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	npuzzle "n-puzzle/n-puzzle"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadFileToArray(fileName string) ([]int, int, error) {
	var numbers []int

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read file: %s", err)
	}

	// Split the file content into lines
	nums := strings.Split(string(data), " ")

	// Convert each line to an integer and append it to the array
	counter := 0
	for _, num := range nums {
		number, err := strconv.Atoi(num)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to convert line to integer: %s", err)
		}
		numbers = append(numbers, number)
		counter++
	}

	return numbers, counter, nil
}

func generateTestFile(i, n int) error {
	path := fmt.Sprintf("tests/test%v.txt", i)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	} else {
		return fmt.Errorf("file with such number already exists")
	}
	a := make([]int, n*n)
	for i := range a {
		a[i] = i
	}
	rand := rand.New(rand.NewSource(time.Now().UnixMilli()))
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write into file
	str_array := fmt.Sprintln(a)
	_, err = file.WriteString(str_array[1 : len(str_array)-2])
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func generateFile(path string, n int) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	} else {
		return fmt.Errorf("file with such name already exists")
	}
	a := make([]int, n*n)
	for i := range a {
		a[i] = i
	}
	rand := rand.New(rand.NewSource(time.Now().UnixMilli()))
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write into file
	str_array := fmt.Sprintln(a)
	_, err = file.WriteString(str_array[1 : len(str_array)-2])
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// err := generateTestFile(6, 4)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	pathPtr := flag.String("path", "", "(REQUIRED) path to a file to save generated puzzle, requires size")
	sizePtr := flag.Int("size", 0, "(REQUIRED) size of one of the sides of the puzzle")
	genPtr := flag.Bool("gen", false, "(OPTIONAL) generates puzzle in the given path file")
	verbosePtr := flag.Bool("v", false, "(OPTIONAL) output verbose solution path")

	flag.Parse()

	if len(*pathPtr) == 0 {
		panic("path is required")
	}

	if !*genPtr {
		s, n, err := ReadFileToArray(*pathPtr)
		if err != nil {
			log.Panic(err)
		}
		if *sizePtr*(*sizePtr) != n {
			log.Panic("bad size supplied")
		}
		n = *sizePtr
		fmt.Printf("%v n: %v\n", s, n)
		node := npuzzle.CreateNode(s, n, 0)
		node.SetFValue()
		result := npuzzle.AStarSearch(node, 1000)
		if *verbosePtr {
			path := npuzzle.ReconstructPath(result)
			for i := len(path) - 1; i >= 0; i-- {
				fmt.Println(path[i])
			}
		} else {
			fmt.Println(result)
		}

	} else {
		if err := generateFile(*pathPtr, *sizePtr); err != nil {
			log.Fatal(err)
		}
		return
	}
}
