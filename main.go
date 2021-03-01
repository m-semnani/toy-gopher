package main

import (
	"fmt"
	"log"
	"os"

	"github.com/m-semnani/toy-gopher/hi"
)

func main() {
	args := os.Args[1:]
	fmt.Printf("%v\n", args)

	for _, msg := range args {
		printMe(msg)
	}
}

func printMe(str string) {
	msg, err := hi.HelloWorld(str)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(msg)
}
