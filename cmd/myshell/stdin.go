package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	i := 0
	for i < len(input) {
		if input[i] == ' ' {
			break
		}

		i++
	}

	if i == len(input) {
		return stdin{
			cmd:  input,
			args: []string{},
		}, nil
	}

	cmd := input[:i]
	rest := input[i+1:]
	args := parseString(rest)

	return stdin{
		cmd:  cmd,
		args: args,
	}, nil
}

func parseString(input string) []string {
	var tokens []string
	var currentToken strings.Builder
	inQuotes := false
	inDoubleQuotes := false
	isEscaped := false

	for _, char := range input {
		if isEscaped {
			currentToken.WriteRune(char)
			isEscaped = false
			continue
		}

		switch char {
		case '\'':
			if inDoubleQuotes {
				currentToken.WriteRune(char)
			} else {
				inQuotes = !inQuotes
			}

		case '"':
			inDoubleQuotes = !inDoubleQuotes

		case ' ':
			if inQuotes || inDoubleQuotes {
				currentToken.WriteRune(char)
			} else {
				if currentToken.Len() > 0 {
					tokens = append(tokens, currentToken.String())
					currentToken.Reset()
				}
			}

		case '\\':
			if inDoubleQuotes {
				currentToken.WriteRune(char)
			} else {
				isEscaped = true
			}

		default:
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
