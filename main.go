package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello, World")
		fmt.Println("Update me and see me rebuild!")
		time.Sleep(time.Second)
	}
}
