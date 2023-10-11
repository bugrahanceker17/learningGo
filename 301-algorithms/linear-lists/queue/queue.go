package main

type Queue struct {
	list []MyItem
}

type MyItem int

func main() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(4)
	q.Enqueue(31)

	println(q.Dequeue())
	println(q.Dequeue())
}

func NewQueue() *Queue {
	return &Queue{list: []MyItem{}}
}

func (q *Queue) Enqueue(value MyItem) {
	q.list = append(q.list, value)
}

func (q *Queue) Dequeue() MyItem {
	first := q.list[0]
	q.list = q.list[1:]

	return first
}
