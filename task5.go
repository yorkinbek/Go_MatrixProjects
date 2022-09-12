package main

import "fmt"

func main() {

	var row int

	fmt.Print("Enter Triangle Pattern Rows = ")
	fmt.Scanln(&row)

	for i := 0; i < row; i++ {

		for j := row; j > i; j-- {
			fmt.Printf(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Printf("#")
			fmt.Printf("#")

		}
		fmt.Println("#")
	}

	for i := row; i > 0; i-- {

		for j := row; j > i; j-- {
			fmt.Printf(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Printf("#")
			fmt.Printf("#")

		}
		fmt.Println("#")
	}
}
