package main

import (
	"log"
	"os"
	"strconv"

	"./helper"
)

func main() {
	argc := len(os.Args)
	switch argc {
	case 2:
		helper.ReadMeHelper(os.Args[1])
	case 3:
		id, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("The parameter should be a number!")
		}
		helper.MakeFolder(id, os.Args[2])
	}
}
