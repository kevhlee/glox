// Package stack contains an implementation of a stack.
package stack

import "iter"

// Stack is a stack data structure.
type Stack[V any] struct {
	data []V
}

// At returns the value at the given index of the stack.
//
// Index 0 represents the bottom of the stack. This function will panic if an
// out-of-bounds index was given.
func (s *Stack[V]) At(index int) V {
	return s.data[index]
}

// Empty checks if the stack is empty.
func (s *Stack[V]) Empty() bool {
	return len(s.data) == 0
}

// IterBot returns an iterator starting from the bottom of the stack to the top.
func (s *Stack[V]) IterBot() iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for i := range s.data {
			if !yield(i, s.data[i]) {
				return
			}
		}
	}
}

// IterTop returns an iterator starting from the top of the stack to the bottom.
func (s *Stack[V]) IterTop() iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		n := len(s.data)

		for i := range s.data {
			if !yield(i, s.data[n-(i+1)]) {
				return
			}
		}
	}
}

// Len returns the length of the stack.
func (s *Stack[V]) Len() int {
	return len(s.data)
}

// Peek returns the top of the stack, assuming the stack is not empty.
//
// This function returns a boolean value to indicate if the stack was not empty.
func (s *Stack[V]) Peek() (V, bool) {
	var zero V
	if s.Empty() {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Pop removes and returns the top of the stack, assuming the stack is not
// empty.
//
// This function returns a boolean value to indicate if the stack was not empty.
func (s *Stack[V]) Pop() (V, bool) {
	var zero V
	if s.Empty() {
		return zero, false
	}

	len := len(s.data) - 1
	top := s.data[len]
	s.data[len] = zero
	s.data = s.data[:len]
	return top, true
}

// Push adds a value to the top of the stack.
func (s *Stack[V]) Push(value V) {
	s.data = append(s.data, value)
}
