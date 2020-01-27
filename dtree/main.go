package main

import (
	"log"
	"os"
)

func main() {
	if err := new(Runner).Run(os.Stdin, os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
