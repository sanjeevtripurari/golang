package main

import "fmt"

func sayHello(n string) {
	fmt.Println("Hello", n)
	var a int = 100

	p := &a

	fmt.Println(*p, a)

	swap(p)

	fmt.Println(*p, a)

}

func swap(a *int) {
	*a = 200
}

func mapFunc() {

	menu := map[string]float64{
		"soup":   4.00,
		"pie":    5.44,
		"salad":  6.776,
		"toffee": 0.112,
	}

	for i, v := range menu {
		fmt.Println(i, v)
	}
	fmt.Println(menu["salad"])

	menu["salad"] = 9.22

	fmt.Println(menu)
}
