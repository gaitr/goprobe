package main

import (
	"log"
	"os"
)

func init() {
}

func main() {

	name, err := os.Hostname()
	if err != nil {
		os.Exit(-1)
	}

	log.Printf("(goprobe) Hello, %s!\n", name)

}
