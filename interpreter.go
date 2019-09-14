package brainfuck

import (
	"errors"
	"fmt"
)

const (
	size = 30000
)

// Interpret for brainfuck interpreter
func Interpret(str string) (string, error) {
	buff := [size]byte{}
	ptr := 0

	output := []byte{}

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			buff[ptr]++
		case '-':
			buff[ptr]--
		case '.':
			output = append(output, buff[ptr])
			fmt.Printf("%c", buff[ptr])
		case ',':
			var input byte
			_, err := fmt.Scanf("%c", &input)
			if err != nil {
				return "", err
			}
			buff[ptr] = input
		case '[':
			if buff[ptr] == 0 {
				for j := i; j < len(str); j++ {
					if str[j] == ']' {
						i = j
						break
					}
					if j == len(str)-1 {
						return "", errors.New("Could not found matching [")
					}
				}
			}
		case ']':
			if buff[ptr] != 0 {
				for j := i; j > 0; j-- {
					if str[j] == '[' {
						i = j
						break
					}
					if j == 0 {
						return "", errors.New("Could not found matching [")
					}
				}
			}
		default:
			return "", errors.New("Invalid character " + string(str[i]))
		}
	}

	return string(output), nil
}
