package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var greeting string = "The quick brown fox jumps off the wall!!"

	//fmt.Println(strings.Count(greeting, "l"))
	/**
	fmt.Println(strings.ToTitle(greeting))
	fmt.Println(strings.Split(greeting, " "))
	*/
	var ages []int = []int{23, 43, 56, 7, 8, 99, 11}
	sort.Ints(ages)
	fmt.Println(greeting)

	index := sort.SearchInts(ages, 9)

	fmt.Println(ages, index)

	names := []string{"Raj", "Rani", "Ramesh", "Rajni"}
	fmt.Println(names)

	//	fmt.Println(sort.SearchStrings(names, "Tikant"))
	//	cycleSay(names, sayHell)
	//	cycleSay(names, sayBye)

	a1 := areaOfCircle(42.11)
	a2 := areaOfCircle(23.11)

	fmt.Println("Area of circles", a1, a2)

}

func areaOfCircle(r float32) float32 {
	return math.Pi * r * r
}

func cycleSay(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func sayHell(n string) {
	fmt.Println("Hello", n)
}

func sayBye(n string) {
	fmt.Println("Good Bye", n)
}

func loop1() {
	var x int = 0
	for x < 5 {

		fmt.Println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	s := []string{"big", "small", "hello", "high", "low"}

	for i, v := range s {
		fmt.Printf("index: = %v, value: %v\n", i, v)
	}
}
