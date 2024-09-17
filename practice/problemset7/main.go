package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(w, str)
	}
}

func main() {
	go display("geeks")
	display("hello")
	go func() {
		fmt.Println("welcome")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("good bye main")
}
