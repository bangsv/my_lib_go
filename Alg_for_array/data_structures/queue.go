package Alg_for_array

type ptr int // type alias

type queae struct {
	data []int
	next ptr
	size int
}

func (q *queae) push(data int) {
	q.data = append(q.data, data)
	q.size++
}

func (q *queae) pop() {

	if q == nil || q.size == 0 {
		panic("Empty queue")
	}

	q.data = q.data[1:]
	q.size--
}

func (q *queae) front() int {
	if q == nil || q.size == 0 {
		panic("Empty queue")
	}
	return q.data[0]
}

func (q *queae) back() int {
	if q == nil || q.size == 0 {
		panic("Empty queue")
	}
	return q.data[q.size-1]
}

func (q *queae) empty() bool {
	return q.size == 0
}
