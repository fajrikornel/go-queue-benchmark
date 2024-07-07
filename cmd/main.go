package main

import (
	"fmt"

	q "example.com/go-dsa/internal/queue"
)

func main() {
	queue := q.New[int](5)

	fmt.Printf("%v\n", queue)
}
