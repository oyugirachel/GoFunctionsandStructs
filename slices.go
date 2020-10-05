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
	var d [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}
	fmt.Printf("%q", d)
	// [["row 1 - column 1" "row 1 - column 2" "row 1 - column 3"]
	//  ["row 2 - column 1" "row 2 - column 2" "row 2 - column 3"]]
}
