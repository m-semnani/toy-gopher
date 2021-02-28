package main

import (
	"fmt"
	"log"

	"github.com/m-semnani/toy-gopher/hi"
)

func main() {
	printMe("World!")
	printMe("")
}

func printMe(str string) {
	msg, err := hi.HelloWorld(str)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
