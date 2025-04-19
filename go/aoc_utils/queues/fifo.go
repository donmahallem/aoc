package queues

type FifoQueue[A any] struct {
	data        []A
	size, index uint
	len         uint
}

func NewFifoQueue[A any](size uint) *FifoQueue[A] {
	return &FifoQueue[A]{size: size, data: make([]A, size)}
}

func (a *FifoQueue[A]) Add(b A) {
	a.data[a.index] = b
	a.index = (a.index + 1) % a.size
	a.len = max(a.len+1, a.size)
}

func (a *FifoQueue[A]) Len() uint {
	return a.len
}
func (a *FifoQueue[A]) Size() uint {
	return a.size
}
func (a *FifoQueue[A]) Get(idx uint) A {
	return a.data[(a.index+idx)%a.len]
}
