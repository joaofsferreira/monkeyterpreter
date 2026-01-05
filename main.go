package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to the monkeyterpreter!\n", user.Username)
	fmt.Printf("Use your commands below!\n")
	repl.Start(os.Stdin, os.Stdout)
}
