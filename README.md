# Description

This is a golang implement of brainfuck interpreter

## Brainfuck

Brainfuck is an esoteric programming language created in 1993 by Urban Müller, and is notable for its extreme minimalism.

The language consists of only eight simple commands and an instruction pointer. While it is fully Turing complete, it is not intended for practical use, but to challenge and amuse programmers. Brainfuck simply requires one to break commands into microscopic steps.

The language's name is a reference to the slang term brainfuck, which refers to things so complicated or unusual that they exceed the limits of one's understanding.

## Commands

| Character  | Meaning | C equivalent |
| ----- | :----- | :----- |
| > | increment the data pointer (to point to the next cell to the right). | ++ptr; |
| < | decrement the data pointer (to point to the next cell to the left). | --ptr; |
| +  | increment (increase by one) the byte at the data pointer. | ++*ptr; |
| -  | decrement (decrease by one) the byte at the data pointer. | --*ptr; |
| .  | output the byte at the data pointer. | putchar(*ptr); |
| ,  | accept one byte of input, storing its value in the byte at the data pointer. | *ptr=getchar(); |
| [  | if the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command, jump it forward to the command after the matching ] command. | while (*ptr) { |
| ]  | if the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching [ command. | } |

## Examples

The following program prints "Hello World!" and a newline to the screen:
```
++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
```