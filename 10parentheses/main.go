package main

import (
	"fmt"
)

// Stack is stack
type Stack struct {
	data []rune
}

func (s *Stack) push(r rune) {
	s.data = append(s.data, r)
}
func (s *Stack) pop() {
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[0 : len(s.data)-1]
}
func (s *Stack) top() rune {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func getStack() *Stack {
	s := &Stack{
		data: []rune{'(', ')'},
	}
	s.push('[')
	s.push(']')
	s.pop()
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", s.top())
	return s
}

func validChar(c rune) bool {
	switch c {
	case '{', '}', '[', ']', '(', ')':
		return true
	}
	return false
}

func invertChar(c rune) rune {
	switch c {
	case '{':
		return '}'
	case '(':
		return ')'
	case '[':
		return ']'
	}
	return 0
}
func isValid(s string) bool {
	stack := &Stack{}
	stack.data = make([]rune, 0, 64)
	for _, c := range s {
		// fmt.Printf("len=%d cap=%d\n", len(stack.data), cap(stack.data))
		if validChar(c) {
			if invertChar(stack.top()) == c {
				stack.pop()
			} else {
				stack.push(c)
			}
		} else {
			return false
		}
	}
	return stack.isEmpty()
}

func main() {
	fmt.Printf("%q", rune(7897))
	valid := isValid("()")
	fmt.Printf("%v\n", valid)
}
