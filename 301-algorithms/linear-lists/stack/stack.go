package main

type Stack struct {
	cap   uint64
	depth uint64
	list  []MyItem
}

type MyItem int

func main() {
	s := newStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	println(s.Pop())
}

func newStack(cap uint64) *Stack {
	return &Stack{
		cap:   cap,
		depth: 0,
		list:  make([]MyItem, cap),
	}
}

func (s *Stack) Pop() MyItem {
	if s.depth > 0 {
		s.depth--
		return s.list[s.depth]
	}

	return -1
}

func (s *Stack) Push(values MyItem) {
	if s.depth >= s.cap {
		panic("out of cap")
	}
	s.list[s.depth] = values
	s.depth++
}
