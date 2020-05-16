package ctci

// Stack can be implemented using a linkedList or an array
type Stack struct {
	Top *StackNode
}

// StackNode represents an object in a stack
type StackNode struct {
	Data interface{}
	Next *StackNode
}

// Push adds to top of stack
func (s *Stack) Push(i interface{}) {
	sn := &StackNode{
		Data: i,
	}

	if s.Top == nil {
		s.Top = sn
		return
	}

	sn.Next = s.Top
	s.Top = sn
}

// Pop removes the top value from a Stack, and returns it (consumes the node)
func (s *Stack) Pop() *StackNode {
	t := s.Top
	if t.Next != nil {
		s.Top = s.Top.Next
	}
	return t
}

// Peek returns the top node in the stack without consuming it
func (s *Stack) Peek() *StackNode {
	return s.Top
}

// IsEmpty checks if a stack is empty
func (s *Stack) IsEmpty() bool {
	if s.Top == nil {
		return true
	}
	return false
}
