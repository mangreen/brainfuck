package brainfuck

import (
	"testing"
)

func TestInterpret(t *testing.T) {
	hello := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."

	output, err := Interpret(hello)
	if err != nil {
		t.Error(err.Error())
	}
	if output != "Hello World!\n" {
		t.Errorf("wrong output, expected Hello World!\\n and got %v", output)
	}
}
