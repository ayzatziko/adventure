package main

import (
	"flag"
	"log"
	"os"

	"github.com/ayzatziko/adventure"
)

var (
	filename = flag.String("file", "./sample", "adventure filename")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	f, err := os.Open(*filename)
	if err != nil {
		return err
	}
	defer f.Close()

	adv, _ := adventure.Read(f)
	adv.Run(os.Stdin)
	return nil
}
