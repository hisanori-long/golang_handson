package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := "1 2 3 4 5 142 32"
	ar := strings.Split(x, " ")
	t := 0
	for _, v := range ar {
		n, err := strconv.Atoi(v)
		if err != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!")
}
