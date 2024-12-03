package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

type stdin struct {
	cmd  string
	args []string
}

func readFromStdin() (stdin, error) {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return stdin{}, fmt.Errorf("Error reading from stdin: %v", err)
	}

	input = input[:len(input)-1]
	inputs := parseString(input)

	cmd := inputs[0]
	args := inputs[1:]

	return stdin{
		cmd:  cmd,
		args: args,
	}, nil
}

func parseString(input string) []string {
	args := []string{}
	arg := []byte{}

	current := 0

	consume := func() byte {
		ch := input[current]
		current += 1
		return ch
	}

	atEnd := func() bool {
		return current >= len(input)
	}

	peek := func() byte {
		return input[current]
	}

	for !atEnd() {
		ch := consume()
		switch ch {
		case ' ':
			if len(arg) != 0 {
				args = append(args, string(arg))
				arg = []byte{}
			}

		case '\'':
			for !atEnd() {
				ch = consume()
				if ch == '\'' {
					break
				}

				arg = append(arg, ch)
			}

		case '"':
			for !atEnd() {
				ch = consume()
				if ch == '"' {
					break
				}

				if ch == '\\' {
					ch := peek()
					if ch == '"' || ch == '\\' || ch == '$' || ch == '\n' {
						arg = append(arg, ch)
						consume()
					} else {
						arg = append(arg, '\\')
					}

					continue
				}

				arg = append(arg, ch)
			}

		case '\\':
			arg = append(arg, consume())

		default:
			arg = append(arg, ch)
		}
	}

	if len(arg) != 0 {
		args = append(args, string(arg))
	}

	return args
}
