package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello.. welcom to Go..")
	// arrytest()
	// sayHello("Guru")
	// showSCores()
	//testString()

	fmt.Println("Table of 2:")
	for i := 1; i <= 10; i++ {
		fmt.Printf("2 x %d = %d\n", i, 2*i)
	}

	fmt.Println("\nString Slice Example:")
	stringSliceExample()
}

func testString() {
	greeting := "Hello, there! How are you, friends?"

	// Length of the string
	fmt.Println("Length:", len(greeting))

	// Convert to upper and lower case
	fmt.Println("Uppercase:", strings.ToUpper(greeting))
	fmt.Println("Lowercase:", strings.ToLower(greeting))

	// Check if string contains a substring
	fmt.Println("Contains 'How':", strings.Contains(greeting, "How"))

	// Replace all occurrences of a substring
	fmt.Println("Replace 'Hello' with 'Hi':", strings.ReplaceAll(greeting, "Hello", "Hi"))

	// Split the string into a slice
	words := strings.Split(greeting, " ")
	fmt.Println("Split words:", words)

	// Join a slice of strings
	fmt.Println("Join words:", strings.Join(words, "-"))

	// Trim spaces from start and end
	paddedString := "  trimmed string  "
	fmt.Println("Trimmed:", strings.TrimSpace(paddedString))

	// Check if string starts or ends with a substring
	fmt.Println("Starts with 'Hello':", strings.HasPrefix(greeting, "Hello"))
	fmt.Println("Ends with 'friends?':", strings.HasSuffix(greeting, "friends?"))

	// Find the index of a substring
	fmt.Println("Index of 'How':", strings.Index(greeting, "How"))

	// Count occurrences of a substring
	fmt.Println("Count of 'e':", strings.Count(greeting, "e"))

}

func stringSliceExample() {
	// Create a string slice
	fruits := []string{"apple", "banana", "cherry", "date"}

	// Print the original slice
	fmt.Println("Original slice:", fruits)

	// Append to the slice
	fruits = append(fruits, "elderberry")
	fmt.Println("After append:", fruits)

	// Slice operations
	fmt.Println("First two fruits:", fruits[:2])
	fmt.Println("Last two fruits:", fruits[len(fruits)-2:])

	// Iterate over the slice
	fmt.Println("Iterating over fruits:")
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}

	// Check if an element exists
	searchFruit := "banana"
	for _, fruit := range fruits {
		if fruit == searchFruit {
			fmt.Printf("%s found in the slice\n", searchFruit)
			break
		}
	}

	// Remove an element (by creating a new slice)
	indexToRemove := 2
	fruits = append(fruits[:indexToRemove], fruits[indexToRemove+1:]...)
	fmt.Println("After removing element at index 2:", fruits)

	// Copy a slice
	fruitsCopy := make([]string, len(fruits))
	copy(fruitsCopy, fruits)
	fmt.Println("Copied slice:", fruitsCopy)
}
