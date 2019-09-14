package brainfuck

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	code := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."

	output, err := Interpret(code)
	if err != nil {
		t.Error(err.Error())
	}
	if output != "Hello World!\n" {
		t.Errorf("wrong output, expected Hello World!\\n and got %v", output)
	}
}

func TestFwdBreakLoop(t *testing.T) {
	code := "++[->+++<]>."

	output, err := Interpret(code)
	if err != nil {
		t.Error(err.Error())
	}
	if output != string([]byte{6}) {
		t.Fail()
	}
}

func TestBackBreakLoop(t *testing.T) {
	code := "++[>+++<-]>."

	output, err := Interpret(code)
	if err != nil {
		t.Error(err.Error())
	}
	if output != string([]byte{6}) {
		t.Fail()
	}
}

func TestJumpForward(t *testing.T) {
	code := "++--[+++]."

	output, err := Interpret(code)
	if err != nil {
		t.Error(err.Error())
	}
	if output != string([]byte{0}) {
		t.Fail()
	}
}

func TestNestedLoop(t *testing.T) {
	code := "++++++++[>++++++<-]++++[>.+<-]"

	output, err := Interpret(code)
	if err != nil {
		t.Error(err.Error())
	}
	if output != "0123" {
		t.Fail()
	}
}

func TestAddTwo(t *testing.T) {
	code := "++>+++++[<+>-]++++++++[<++++++>-]<."

	output, err := Interpret(code)
	if err != nil {
		t.Error(err.Error())
	}
	if output != "7" {
		t.Fail()
	}
}
