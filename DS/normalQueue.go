package main

import (
	"fmt"
)

type Queue struct {
	pool []string
}


// Check if queue is empty
func (q *Queue) UnderFlow() bool {
	return len(q.pool) == 0
}

// Enqueue operation
func (q *Queue) Enqueue(v string) {
	q.pool = append(q.pool, v)
}

// Dequeue operation
func (q *Queue) Dequeue() (bool, string) {
	if q.UnderFlow() {
		return true, ""
	} else {
		removed := q.pool[0]
		q.pool = q.pool[1:]
		return true, removed
	}
}

func main() {

	queue := Queue{}
	queue.Enqueue("This")
	queue.Enqueue("is")
	queue.Enqueue("a")
	queue.Enqueue("queue")

	// Print queue
	fmt.Println(queue)

	for len(queue.pool) > 0 {
		check, value := queue.Dequeue()
		if check == true {
			fmt.Println(value)
		}
	}
}