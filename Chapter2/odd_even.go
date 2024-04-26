package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "11"
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print(x + "は、")
	} else {
		return
	}
	switch {
	case n%2 == 0:
		fmt.Println("偶数です")
	case n%2 == 1:
		fmt.Println("奇数です")
	}
}
