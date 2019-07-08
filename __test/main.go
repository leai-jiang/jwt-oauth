package main

import (
	"crypto"
	"fmt"
)

type A struct {
	name string
}

func (a *A) Name() string {
	return a.name
}

type B struct {
	A
	age int
}

func main() {
	var b B
	b.name = "lei"
	b.age = 26

	fmt.Print(b.Name())

	fmt.Println(crypto.SHA256)
}
