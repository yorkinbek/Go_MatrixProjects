package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		y := strings.Repeat("#", i)
		fmt.Println(y)
	}
}
