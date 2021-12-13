package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SubmarineController interface {
	forward(int)
	down(int)
	up(int)
	getPosition() int
}

type Submarine struct {
	horizontalPosition, depth int
}

func (s *Submarine) forward(steps int) {
	s.horizontalPosition += steps
}

func (s *Submarine) down(steps int) {
	s.depth += steps
}

func (s *Submarine) up(steps int) {
	s.depth -= steps
}

func (s *Submarine) getPosition() (position int) {
	return s.depth * s.horizontalPosition
}

type SubmarineV2 struct {
	horizontalPosition, depth, aim int
}

func (s *SubmarineV2) forward(steps int) {
	s.horizontalPosition += steps
	s.depth += s.aim * steps
}

func (s *SubmarineV2) down(steps int) {
	s.aim += steps
}

func (s *SubmarineV2) up(steps int) {
	s.aim -= steps
}

func (s *SubmarineV2) getPosition() (position int) {
	return s.depth * s.horizontalPosition
}

func pilot(s SubmarineController, commands []string) {
	for _, command := range commands {
		commandList := strings.Split(command, " ")
		// fmt.Println("Commands are:", commandList[0], "-", commandList[1])

		cmd := strings.ToLower(commandList[0])
		param, _ := strconv.Atoi(commandList[1])

		switch cmd {
		case "up":
			s.up(param)

		case "down":
			s.down(param)

		case "forward":
			s.forward(param)

		default:
			panic(fmt.Sprintf("Invalid command: %s", commandList))
		}
	}
}

func getCommands(fileName string) (commands []string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}
	return
}

func main() {
	commands := getCommands("input.txt")

	var sub SubmarineController
	sub = &Submarine{horizontalPosition: 0, depth: 0}
	pilot(sub, commands)
	fmt.Println("Part one answer:", sub.getPosition())

	var subV2 SubmarineController
	subV2 = &SubmarineV2{horizontalPosition: 0, depth: 0, aim: 0}
	pilot(subV2, commands)
	fmt.Println("Part two answer:", subV2.getPosition())
}

/*
Output: Day 2
$> go run solution.go

Part one answer: 1427868
Part two answer: 1568138742
*/
