package stack_test

import (
	"testing"

	"github.com/kevhlee/glox/internal/stack"
)

// TestIter checks to the make sure the stack iterators work correctly.
func TestIter(t *testing.T) {
	var s stack.Stack[int]

	var (
		expectedIndex  = 0
		expectedValues = []int{1, 3, 7, 5, 2}
	)

	for _, expectedValue := range expectedValues {
		s.Push(expectedValue)
	}

	// Test starting from the top

	for index, value := range s.IterTop() {
		if index != expectedIndex {
			t.Errorf("Expected index %d but got %d", expectedIndex, index)
		}
		if expectedValue := expectedValues[len(expectedValues)-(expectedIndex+1)]; value != expectedValue {
			t.Errorf("Expected value %d but got %d", expectedValue, value)
		}
		expectedIndex++
	}

	// Test starting from the bottom

	expectedIndex = 0

	for index, value := range s.IterBot() {
		if index != expectedIndex {
			t.Errorf("Expected index %d but got %d", expectedIndex, index)
		}
		if expectedValue := expectedValues[expectedIndex]; value != expectedValue {
			t.Errorf("Expected value %d but got %d", expectedValue, value)
		}
		expectedIndex++
	}
}

// TestReadWriteOperations checks to make sure the stack works correctly when
// performing different read/write operations on it.
func TestReadWriteOperations(t *testing.T) {
	var s stack.Stack[int]
	if _, ok := s.Peek(); ok {
		t.Error("Stack should be empty")
	}

	s.Push(1)
	s.Push(2)
	if value, ok := s.Peek(); !ok {
		t.Error("Stack should not be empty")
	} else if value != 2 {
		t.Errorf("Expected value 2, got %d instead", value)
	}

	s.Push(3)
	if value, ok := s.Peek(); !ok {
		t.Error("Stack should not be empty")
	} else if value != 3 {
		t.Errorf("Expected value 3, got %d instead", value)
	}

	s.Pop()
	s.Pop()
	s.Pop()
	if _, ok := s.Peek(); ok {
		t.Error("Stack should be empty")
	}
}
