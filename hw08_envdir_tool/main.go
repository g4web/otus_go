package main

import (
	"log"
	"os"
)

func main() {
	environments, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	returnCode := RunCmd(os.Args[2:], environments)
	if returnCode != 0 {
		log.Printf("The command is executed with code %d\n", returnCode)
		log.Fatal()
	}
}
