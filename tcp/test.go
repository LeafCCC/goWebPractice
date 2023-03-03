package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "From 1: SH"
	b := strings.Split(a, " ")
	fmt.Println(b)
	fmt.Println(len(b))

	if b[2] == "SH" {
		fmt.Println("ok")
	}
}
