package main

import (
	"fmt"
)

func main() {

	var row int
	var Star string
	fmt.Print("Enter Pattern Rows = ")
	fmt.Scanln(&row)

	for i := 0; i < row; i++ {
		Star = "#"
		for j := row; j > i; j-- {
			fmt.Printf(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Printf(Star)
		}
		fmt.Println()
	}

}
