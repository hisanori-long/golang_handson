package main

import (
	"fmt"
)

func main() {
	a := [5]int{10, 20, 30, 40, 50}
	b := a[0:3]
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("a[1]に200を代入")
	a[1] = 200
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("bに4000を追加")
	b = append(b, 4000)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("push 5000")
	b, _ = push(b, 5000, 5400, 5600)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("pop")
	b, _ = pop(b)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("insert 6000 at 2")
	b = insert(b, 6000, 2)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("remove 6000")
	f := func(a []int) ([]int, int) {
		return a, len(a)
	}
	fmt.Println(f(b))

}

func push(a []int, v ...int) (s []int, l int) {
	s = append(a, v...)
	return
}

func pop(a []int) ([]int, int) {
	return a[:len(a)-1], len(a) - 1
}

func shift(a []int) []int {
	return a[1:]
}

func unshift(a []int, v int) []int {
	return append([]int{v}, a...)
}

func insert(a []int, v int, p int) (s []int) {
	s = append(a, 0)
	s = append(s[:p+1], s[p:len(s)-1]...)
	s[p] = v
	return
}

func remove(a []int, p int) []int {
	return append(a[:p], a[p+1:]...)
}
