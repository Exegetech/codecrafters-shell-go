package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type builtin int

const (
	echo builtin = iota
	exit
	type_
)

func (b builtin) String() string {
	switch b {
	case echo:
		return "echo"
	case exit:
		return "exit"
	case type_:
		return "type"
	default:
		return "unknown"
	}
}

var builtins = map[string]bool{
	echo.String():  true,
	exit.String():  true,
	type_.String(): true,
}

func main() {
	for true {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := readFromStdin()
		if err != nil {
			log.Fatal(err)
		}

		path := os.Getenv("PATH")
		cmd := input[0]
		args := input[1:]

		switch cmd {
		case exit.String():
			handleExit()

		case echo.String():
			handleEcho(args)

		case type_.String():
			handleType(args[0], path)

		default:
			fullPath, ok := getFullPath(cmd, path)
			if ok {
				executeCmd(fullPath, args)
			} else {
				fmt.Printf("%s: command not found\n", cmd)
			}
		}
	}
}

func readFromStdin() ([]string, error) {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return []string{}, fmt.Errorf("Error reading from stdin: %v", err)
	}

	withoutDelim := input[:len(input)-1]
	return strings.Split(withoutDelim, " "), nil
}

func handleExit() {
	os.Exit(0)
}

func handleEcho(args []string) {
	joined := strings.Join(args, " ")
	fmt.Println(joined)
}

func handleType(cmd string, path string) {
	if ok := builtins[cmd]; ok {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}

	fullPath, ok := getFullPath(cmd, path)
	if ok {
		fmt.Printf("%s is %s\n", cmd, fullPath)
		return
	}

	fmt.Printf("%s: command not found\n", cmd)
}

func getFullPath(cmd, path string) (string, bool) {
	paths := strings.Split(path, ":")

	for _, p := range paths {
		fullPath := filepath.Join(p, cmd)
		if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
			continue
		}

		return fullPath, true
	}

	return "", false
}

func executeCmd(cmd string, args []string) {
	fmt.Println("Executing", cmd, args)
}
