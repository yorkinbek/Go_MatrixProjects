package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%v", &n)
	for i := n; i >= 1; i-- {
		y := strings.Repeat("#", i)
		fmt.Println(y)
	}
}
