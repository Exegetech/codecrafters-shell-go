package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	for true {
		fmt.Fprint(os.Stdout, "$ ")

		stdin, err := readFromStdin()
		if err != nil {
			log.Fatal(err)
		}

		path := os.Getenv("PATH")

		switch stdin.cmd {
		case exit.String():
			handleExit()

		case echo.String():
			handleEcho(stdin.args)

		case type_.String():
			handleType(stdin.args[0], path)

		case pwd.String():
			handlePwd()

		case cd.String():
			handleCd(stdin.args[0])

		default:
			fullPath, ok := getFullPath(stdin.cmd, path)
			if ok {
				executeCmd(fullPath, stdin.args)
			} else {
				fmt.Printf("%s: command not found\n", stdin.cmd)
			}
		}
	}
}
