package main

import (
	"fmt"

	"github.com/hisanori-long/golang_handson/tree/main/Chapter3/modules/pointer"
)

func main() {
	message := pointer.Hello("John")
	fmt.Println(message)
}
