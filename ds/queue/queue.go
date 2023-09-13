package main

import (
	"errors"
	"fmt"
)

var errQueueEmpty = errors.New("queue is empty")

type queue struct {
	elems []any
}

func (q *queue) enqueue(ele any) {
	q.elems = append(q.elems, ele)
}

func (q *queue) dequeue() (any, error) {
	if len(q.elems) == 0 {
		return nil, errQueueEmpty
	}
	ele := q.elems[0]
	q.elems = q.elems[1:]

	return ele, nil
}

func (q *queue) peek() (any, bool) {
	if len(q.elems) == 0 {
		return nil, false
	}

	return q.elems[0], true
}

func (q *queue) clear() {
	*q = queue{}
}

func main() {
	q := new(queue)
	q.enqueue(100)
	q.enqueue(101)
	q.enqueue(102.3)

	for i := 0; ; i++ {
		if ele, ok := q.peek(); ok {
			fmt.Printf("queue peek: %v\n", ele)
		}

		ele, err := q.dequeue()
		if err != nil {
			if errors.Is(err, errQueueEmpty) {
				break
			}
			fmt.Println(err)
		}
		fmt.Printf("enqueue: index: %d, value: %v\n", i, ele)

	}

	q.clear()
}
