package main

import (
	"KDRC-Client/cmd"
	"log"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("Command failed with :%s\n", err.Error())
	}
}
