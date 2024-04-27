package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "40"
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print(("1から" + x + "までの合計は、"))
	} else {
		return
	}
	t := 0 // total
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println(t, "です")

	if err == nil {
		fmt.Print("1から" + x + "までの偶数の合計は、")
	}

	odd_t := 0
	c := 0
	for {
		c++
		if c > n {
			break
		}
		if c%2 == 1 {
			continue
		}
		odd_t += c
	}
	fmt.Println(odd_t, "です")

}
