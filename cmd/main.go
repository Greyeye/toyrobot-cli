package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	robot := NewRobot(false, TableXMaxLimit, 0, TableYMaxLimit, 0, NORTH)
	ih := InputHandler{Reader: os.Stdin}
	fmt.Println("robot started, please enter your command.\nAcceptable commands:\n    place *,*,direction\n    move\n    left\n    right\n    report")
	for {
		fmt.Print(">")
		robot.Controller(ih)
	}
}

// InputHandler handles the input
//
// to initialize InputHandler{Reader: os.Stdin}
// to mock InputHandler{strings.NewReader("test input")}
type InputHandler struct {
	Reader io.Reader
}

// userInput uses the Reader from InputHandler to read input.
func (s *InputHandler) userInput() string {
	scanner := bufio.NewScanner(s.Reader)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}
