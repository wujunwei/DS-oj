package sample

type myHeap struct {
	arr []int
}

func (m *myHeap) Len() int {
	return len(m.arr)
}

func (m *myHeap) Less(i, j int) bool {
	return m.arr[i] < m.arr[j]
}

func (m *myHeap) Swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

func (m *myHeap) Push(x interface{}) {
	m.arr = append(m.arr, x.(int))
}

func (m *myHeap) Pop() interface{} {
	top := m.arr[len(m.arr)-1]
	m.arr = m.arr[:len(m.arr)-1]
	return top
}

//var h *myHeap
//heap.init(h)
//heap.Push(h, 123)
