package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// windowSize is the batching required to compare depths with
func find_increasing_depth_count(scanner bufio.Scanner, windowSize int) (count int, err error) {
	window := make([]int, 0)    // queue impl via int slice
	windowLog := make([]int, 0) // int slice

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text()) // convert string to int
		// fmt.Println("Found depth:", depth)
		if err != nil {
			fmt.Println(err.Error())
			return count, err
		}

		window = append(window, depth)
		// fmt.Print("Window is:", window)
		if len(window) == windowSize {
			windowLog = append(windowLog, sum(window))
			window = window[1:] // pop left element
		}
	}

	prev_depth := 0
	count = -1

	// fmt.Println("Window log is as follow: ", windowLog)

	for _, depth := range windowLog {
		if prev_depth < depth {
			count += 1
		}
		prev_depth = depth
	}
	return
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	partOneAnswer, err := find_increasing_depth_count(*scanner, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part one answer:", partOneAnswer)

	_, err = file.Seek(0, 0) // seek file to start
	if err != nil {
		panic(err)
	}

	partTwoAnswer, err := find_increasing_depth_count(*scanner, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part two answer:", partTwoAnswer)
}

/*
Output:
$> go run solution.go

Part one answer: 1390
Part two answer: 1457
*/
