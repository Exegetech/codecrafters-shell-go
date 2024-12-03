package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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

	fmt.Printf("%s: not found\n", cmd)
}

func handlePwd() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return
	}

	fmt.Println(pwd)
}

func handleCd(target string) {
	if strings.HasPrefix(target, "/") {
		if _, err := os.Stat(target); errors.Is(err, os.ErrNotExist) {
			fmt.Printf("cd: %s: No such file or directory\n", target)
			return
		}

		os.Chdir(target)
		return
	}

	if target == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			return
		}

		os.Chdir(home)
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return
	}

	fullPath := filepath.Join(pwd, target)
	os.Chdir(fullPath)
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

func executeCmd(cmd string, args []string) error {
	command := exec.Command(cmd, args...)
	out, err := command.Output()
	if err != nil {
		return fmt.Errorf("Error executing command: %v\n%s", err, out)
	}

	fmt.Print(string(out))
	return nil
}
