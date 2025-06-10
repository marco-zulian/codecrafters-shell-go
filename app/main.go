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
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}
