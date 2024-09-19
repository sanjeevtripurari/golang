package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// stringfunc()
	// pointerfunc()
	// arrayfunc()
	rangefunc()
}

func rangefunc() {

	var numbers [10]int

	for i := 0; i < 10; i++ {
		numbers[i] = rand.Intn(10)
	}

	fmt.Println(numbers)

	for i := range numbers {
		fmt.Printf("slice item %d, %d\n", i, numbers[i])
	}

	countryCapMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	for country := range countryCapMap {
		fmt.Println("Capital of ", country, "is", countryCapMap[country])
	}

	for country, capital := range countryCapMap {
		fmt.Println("Capital of", country, "is", capital)
	}

}

func arrayfunc() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
	for i := 0; i <= len(numbers); i++ {
		if i == 100 {
			break
		}
		numbers = append(numbers, i)
		fmt.Printf("%d ", i)
	}
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func pointerfunc() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("Address of variable: %x, value: %d\n", &a, a)
	fmt.Printf("Address in pointer: %x, value: %d\n", ip, *ip)
}

func stringfunc() {

	var greeting = "Hello World!"
	fmt.Printf("normal string: ")
	fmt.Printf("%s\n", greeting)

	fmt.Print("Hex bytes: ")

	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}
	fmt.Println()

}
