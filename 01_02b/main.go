package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	msgArr := strings.Split(msg, " ")
	for _, item := range msgArr {
		var newWord string
		for i, char := range item {
			newWord += strings.Repeat(string(char), i+1)
		}
		print(newWord)
	}

}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
