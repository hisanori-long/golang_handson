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

}

func push(a []int, v int) []int {
	return append(a, v)
}

func pop(a []int) []int {
	return a[:len(a)-1]
}

func shift(a []int) []int {
	return a[1:]
}

func unshift(a []int, v int) []int {
	return append([]int{v}, a...)
}

func insert(a []int, p, v int) []int {
	return append(a[:p], append([]int{v}, a[p:]...)...)
}

func remove(a []int, p int) []int {
	return append(a[:p], a[p+1:]...)
}
