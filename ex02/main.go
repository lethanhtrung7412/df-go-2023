package main

import (
	"fmt"
	"log"

	"github.com/lethanhtrung7412/df-go-2023/ex02/cmd"
)

func main() {
	sortedString, err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sortedString)
}
