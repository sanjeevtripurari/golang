package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")
	mybill.updateTip(10)
	mybill.updateItem("corn ", 11)
	mybill.updateItem("soup ", 10)
	mybill.updateItem("momo ", 30)
	mybill.updateItem("juice ", 12)

	fmt.Println(mybill.format())
}
