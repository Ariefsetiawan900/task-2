package main

import (
	"fmt"
	"strings"
)

type triangle interface {
	rataKiri() int
	rataKanan() int
}

type input struct {
	value int
}

func (c input) rataKiri() int {
	for i := 1; i <= c.value; i++ {
		if c.value > 10 || c.value < 5 {
			fmt.Println("invalid data low/high")
		} else {
			fmt.Printf("%s\n", strings.Repeat("*", i))
		}
	}
	return c.value
}

func (c input) rataKanan() int {
	for i := c.value; i >= 1; i-- {
		switch c.value {
		case 5:
			fmt.Printf("%5s\n", strings.Repeat("*", i))
		case 6:
			fmt.Printf("%6s\n", strings.Repeat("*", i))
		case 7:
			fmt.Printf("%7s\n", strings.Repeat("*", i))
		case 8:
			fmt.Printf("%8s\n", strings.Repeat("*", i))
		case 9:
			fmt.Printf("%9s\n", strings.Repeat("*", i))
		case 10:
			fmt.Printf("%10s\n", strings.Repeat("*", i))
		default:
			fmt.Println("invalid data low/high")
		}

	}
	return c.value

}

func main() {
	var bangunDatar triangle
	bangunDatar = input{7}
	bangunDatar.rataKiri()
	fmt.Println("")
	bangunDatar.rataKanan()
}
