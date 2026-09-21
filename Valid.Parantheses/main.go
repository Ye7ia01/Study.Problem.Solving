package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}
	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item, nil
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func isValid(s string) bool {
	myStack := NewStack[string]()
	paranthesis := map[string]string{"(": ")", "[": "]", "{": "}"}

	for i := 0; i < len(s); i++ {
		if _, ok := paranthesis[string(s[i])]; ok {
			myStack.Push(string(s[i])) // Is open brancket
		} else {
			curr, _ := myStack.Pop()
			closedParanthesis := paranthesis[curr]
			// Match closing bracket with top open bracket
			if closedParanthesis != string(s[i]) {
				return false
			}
		}
	}
	if myStack.Len() > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
