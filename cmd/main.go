package main

import (
	"fmt"

	q "example.com/go-dsa/internal/queue"
)

func main() {
	var queue q.Queue[int]
	queue = q.NewArrayQueue[int](5)

	fmt.Printf("%v\n", queue)
}
