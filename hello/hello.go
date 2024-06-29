package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("toto")
	fmt.Println(message)
}
