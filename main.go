package main

import "fmt"

type test struct {
	int
	string
}

func main() {
	a := test{10, "10"}
	fmt.Println(a)
	b := test{int: 10, string: "10"}
	fmt.Println(b)
}
