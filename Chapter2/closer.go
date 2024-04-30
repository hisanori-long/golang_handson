package main

import "fmt"

func main() {
	data := "*新しい値*"
	m1 := modify(data)
	data = "+new data+"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())

	fmt.Printf("data: %s\n", data)
	n := 123
	b := true
	s := "hello"
	p := 999
	fmt.Printf("number: %d, boolean: %t, string: %s\n", p, n, b, s)
}

func modify(d string) func() []string {
	m := []string{
		"1st", "2nd",
	}
	return func() []string {
		return append(m, d)
	}
}
