package main

import (
	"fmt"
	"strconv"
	"time"
)

func prmsg(n int, s string) {
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func first(n int, c chan string) {
	const nm string = "first-"
	for i := 0; i < 10; i++ {
		s := nm + strconv.Itoa(i)
		prmsg(n, s)
		c <- s
	}
}

func second(n int, c chan string) {
	for i := 0; i < 10; i++ {
		prmsg(n, "second["+<-c+"]")
	}
}

func total(cs chan int, cr chan int) {
	n := <-cs
	fmt.Println("n =", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	cr <- t
}

func main() {
	cs := make(chan int)
	cr := make(chan int)
	go total(cs, cr)
	cs <- 100
	fmt.Println("total:", <-cr)
}
