package main

import (
	"log"
	"os"

	"github.com/mikeschuld/go-gplv3/io"
)

func main() {
	lic, err := os.Open("LICENSE")
	if err != nil {
		log.Fatal(err)
	}

	io.PrintFile(lic)
}
