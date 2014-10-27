package main

import (
	"fmt"

	"./server"
)

func main() {
	err := server.Run()
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("All well :)")
	}
}
