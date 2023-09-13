package main

import "fmt"

type stack []any

func (s *stack) push(ele any) {
	*s = append(*s, ele)
}

func (s *stack) pop() (any, error) {
	last := len(*s) - 1
	data := (*s)[last]
	*s = (*s)[:last]
	return data, nil
}

func main() {
	var s stack
	s.push(100)
	s.push(101.1)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
