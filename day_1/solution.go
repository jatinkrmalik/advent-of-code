package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// windowSize is the batching required to compare depths with
func FindIncreasingDepthCount(scanner bufio.Scanner, windowSize int) (count int, err error) {
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

func seekFileToStart(file os.File) {
	_, err := file.Seek(0, 0) // seek file to start
	if err != nil {
		panic(err)
	}
}

func getFileAndScanner(fileName string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return file, bufio.NewScanner(file)
}

func main() {
	file, scanner := getFileAndScanner("input.txt")
	defer file.Close()

	partOneAnswer, err := FindIncreasingDepthCount(*scanner, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part one answer:", partOneAnswer)

	seekFileToStart(*file)

	partTwoAnswer, err := FindIncreasingDepthCount(*scanner, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part two answer:", partTwoAnswer)
}

/*
Output: Day 1
$> go run solution.go

Part one answer: 1390
Part two answer: 1457
*/
