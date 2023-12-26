package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
