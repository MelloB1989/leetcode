package queue

type Queue struct {
	q      []any
	length int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(v any) {
	q.q = append(q.q, v)
	q.length++
}

func (q *Queue) Poll() any {
	if q.length == 0 {
		return -1
	}
	front := q.q[0]
	q.q = q.q[1:]
	q.length--
	return front
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Flush() {
	q.q = []any{}
	q.length = 0
}
