package main

import (
	"brainfuck"
	"log"
)

func main() {
	hello := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
	if _, err := brainfuck.Interpret(hello); err != nil {
		log.Fatal(err.Error())
	}
}
