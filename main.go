package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello GO LG")
		time.Sleep(10 * time.Second)
	}
}
