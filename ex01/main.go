package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lethanhtrung7412/df_go_23/ex01/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		log.Fatal("Missing argument")
	}

	result, ok := utils.ReName(args)
	if !ok {
		log.Fatal("Error when reordering name")
	}

	fmt.Println(result)
}