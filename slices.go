package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	b := [2]string{"hello", "world!"}
	fmt.Printf("%q\n", b)
	c := [2]string{"hello", "world!"}
	fmt.Println(c)
	// [hello world!]
	fmt.Printf("%s\n", c)
	// [hello world!]
	fmt.Printf("%q\n", c)
	// ["hello" "world!"]
}
