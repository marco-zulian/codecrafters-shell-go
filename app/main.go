package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(command[:len(command)-1] + ": command not found")
}
