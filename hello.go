package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var name1 string = "name1"
	var name2 string = "name two"
	var name3 string

	name4 := "hello one"
	name3 = "Three"

	fmt.Println(name1, name3, name2, name4)
	var i int = 10
	j := 20

	var f float32 = 34.34

	fmt.Println(i, j, f)
	fmt.Print("Hello ")
	fmt.Print("World\n")
	fmt.Printf("Hi %d %2.5f fhow is %s and %s\n", i, f, name2, name4)
	var a [4]int = [4]int{3, 4, 5, 7}
	var b = [4]int{3, 4, 5, 7}

	fmt.Println(a, b)
}
