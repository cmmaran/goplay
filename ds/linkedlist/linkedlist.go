package main

import "fmt"

type node struct {
	data int
	next *node
}

var head *node

func (n *node) append(data int) *node {

	new := &node{
		data: data,
	}

	if head == nil {
		head = new
	} else {
		n.next = new
	}

	return new
}

func (n *node) print() {
	pn := head
	for pn != nil {
		fmt.Println(pn.data)
		pn = pn.next
	}
}

func main() {
	var n *node
	n = n.append(100)
	n = n.append(101)
	n = n.append(102)

	n.print()
}
