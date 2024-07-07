package main

import (
	"fmt"

	dsa "example.com/go-dsa/internal"
)

func main() {
	queue := dsa.New[int](5)

	fmt.Printf("%v\n", queue)
}
