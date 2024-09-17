package main

import (
	"fmt"
	"strings"
)

func intials(n string) (string, string) {
	s := strings.ToUpper(n)

	name := strings.Split(s, " ")
	var initials []string

	for _, v := range name {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"

}
func main() {

	fn1, fn2 := intials("ajay")
	sayHello("Dude")
	fmt.Println(fn1, fn2)
	mapFunc()
}
