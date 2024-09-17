package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose the option (a for add, s for save, t - add tip): ", reader)

	switch opt {
	case "s":

		b.save()
		fmt.Println("saved the bill", b.name)
	case "a":
		fmt.Println("you choose a")
		name, _ := getInput("item name: ", reader)
		price, _ := getInput("item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added")
		promptOptions(b)
	case "t":
		fmt.Println("you choose t")
		tip, _ := getInput("tip: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tipe must be number")
			promptOptions(b)
		}
		b.updateTip(t)
		promptOptions(b)
	default:
		fmt.Println("that is not valid option..")
		promptOptions(b)

	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}
