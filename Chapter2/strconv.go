package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "2001f"

	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	} else {
		fmt.Println(x, "is not a valid number")
	}
}
