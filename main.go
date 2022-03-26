package main

import "fmt"

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	var x int = 10
	fmt.Printf("1. value of x is %d and pointer is %v\n", x, &x)

	var y *int
	fmt.Printf("2. value of y is %d and pointer is %v\n", y, &y)

	y = &x
	fmt.Printf("3. value of y is %d and pointer is %v\n", y, &y)
}
