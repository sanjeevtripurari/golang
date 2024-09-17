package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v\t...$%8.3f\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v\t...$%8.3f\n", "total:", total)
	fs += fmt.Sprintf("\n%-25v\t...$%8.3f\n", "tip:", b.tip)
	return fs
}

func (b *bill) updateItem(name string, price float64) {
	b.items[name] = price
}
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}
