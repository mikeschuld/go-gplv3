package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PrintFile(file *os.File) {
	defer func(lic *os.File) {
		err := lic.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	s := bufio.NewScanner(file)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	err := s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
