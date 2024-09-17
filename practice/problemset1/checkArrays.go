package main

import "fmt"

func arrytest() {

	var a int = 4
	var arr [3]int = [3]int{34, 56, 78}

	var bg = []int{1, 3, 4, 5}
	bg[2] = 10

	bg = append(bg, 10)

	fmt.Println(arr, len(arr))
	fmt.Println(a, arr, bg)

}
