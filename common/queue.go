package common

import "container/list"

type Queue struct {
	list list.List
}

func (q *Queue) Pop() interface{} {
	return q.list.Remove(q.list.Front())
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) Push(value interface{}) {
	q.list.PushBack(value)
}
