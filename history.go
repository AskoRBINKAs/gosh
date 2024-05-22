package main

import "fmt"

var history []string

func AddToHistory(cmd string) {
	history = append(history, cmd)
}

func PrintHistory() {
	for idx, cmd := range history {
		fmt.Println(idx, cmd)
	}
}
