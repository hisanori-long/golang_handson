package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "12"
	fmt.Print(x + "月は、")
	switch n, err := strconv.Atoi(x); n {
	case 0:
		fmt.Println("月が正しくありません")
		fmt.Println(err)
	case 3, 4, 5:
		fmt.Println("春です")
	case 6, 7, 8:
		fmt.Println("夏です")
	case 9, 10, 11:
		fmt.Println("秋です")
	case 12, 1, 2:
		fmt.Println("冬です")
	default:
		fmt.Println("月が正しくありません")
	}

	y := 5 //⭐️
	switch y {
	case f(1):
		fmt.Println("* first case *")
	case f(2):
		fmt.Println("* second case *")
	case f(3):
		fmt.Println("* third case *")
	default:
		fmt.Println("* default case *")
	}
}

func f(n int) int {
	fmt.Println("No, ", n)
	return n
}
