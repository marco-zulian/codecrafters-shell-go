package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		trimCommand := strings.TrimSpace(command)

		exitRegex := regexp.MustCompile(`^exit(\s\d*)?$`)
		echoRegex := regexp.MustCompile(`^echo(\s.*)?$`)
		typeRegex := regexp.MustCompile(`^type(\s.*)?$`)

		switch {
		case exitRegex.MatchString(trimCommand):
			return
		case echoRegex.MatchString(trimCommand):
			match := echoRegex.FindStringSubmatch(trimCommand)
			echoStr := match[1]

			if len(echoStr) > 1 {
				fmt.Println(match[1][1:])
			} else {
				fmt.Println()
			}
		case typeRegex.MatchString(trimCommand):
			commands := map[string]bool{"type": true, "echo": true, "exit": true}
			match := typeRegex.FindStringSubmatch(trimCommand)

			if len(match[1]) <= 1 {
				fmt.Println("usage: type <command>")
			} else if _, ok := commands[match[1][1:]]; ok {
				fmt.Println(match[1][1:] + " is a shell builtin")
			} else {
				fmt.Println(match[1][1:] + ": not found")
			}
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}
