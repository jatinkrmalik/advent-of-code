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

func (s *Submarine) pilot(commands []string) {
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

func (s *Submarine) getPosition() (position int) {
	return s.depth * s.horizontalPosition
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
	sub := Submarine{horizontalPosition: 0, depth: 0}
	commands := getCommands("input.txt")
	sub.pilot(commands)
	fmt.Println(sub.getPosition())

}
