package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ananrafs1/cliit/plugin/shared"
)

func main() {
	shared.Serve(ExecDummy1{})()
}

type ExecDummy1 struct{}

func (e ExecDummy1) GetActionMetadata() map[string]map[string]string {
	return map[string]map[string]string{
		"Summarize": {
			"left":  "insert first argument",
			"right": "insert second argument",
		},
		"Multiply": {
			"Factor 1": "insert first Factor",
			"Factor 2": "insert second Factor",
		},
	}
}

func (e ExecDummy1) Execute(act string, params map[string]string) <-chan string {
	out := make(chan string, 2)

	go func(p map[string]string) {
		time.Sleep(1 * time.Second)
		switch act {
		case "Summarize":
			sum := 0
			for _, v := range p {
				val, _ := strconv.Atoi(v)
				sum += val
			}
			out <- fmt.Sprintf("[%s] %v is %d", act, p, sum)
		case "Multiply":
			mul := 1
			for _, v := range p {
				val, _ := strconv.Atoi(v)
				mul *= val
			}
			out <- fmt.Sprintf("[%s] %v is %d", act, p, mul)
		}
		close(out)

	}(params)

	return out
}
