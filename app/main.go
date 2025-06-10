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

		exitRegex := regexp.MustCompile(`^exit (\d*)$`)
		switch {
		case exitRegex.MatchString(trimCommand):
			return
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}
