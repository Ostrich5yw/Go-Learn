package DataStructure

type Item interface {
}

// Stack the stack of items
type Stack struct {
	Items []Item
}

// New Create a new Stack
func NewStack() *Stack {
	return &Stack{Items: []Item{}}
}

// Push adds an Item to the top of the stack
func (s *Stack) Push(t Item) {
	s.Items = append(s.Items, t)
}

// Pop removes an Item from the top of the stack
func (s *Stack) Pop() *Item {
	item := s.Items[len(s.Items)-1] // 后进先出
	s.Items = s.Items[0 : len(s.Items)-1]
	return &item
}

// Return an Item from the top of the stack
func (s *Stack) Top() *Item {
	item := s.Items[len(s.Items)-1] // 后进先出
	return &item
}

// Return an Item from the top of the stack
func (s *Stack) Empty() bool {
	return len(s.Items) == 0
}
