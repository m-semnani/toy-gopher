package main

import "fmt"

type MaxHeap struct {
	array []int
}

func NewMaxHeap(inputArray []int) MaxHeap {
	newHeap := &MaxHeap{}
	for _, v := range inputArray {
		newHeap.Insert(v)
	}
	return *newHeap
}

func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	if h.lastIndex() != 0 {
		h.siftUp(h.lastIndex())
	}
}

func (h *MaxHeap) Max() int {
	if len(h.array) == 0 {
		return -1
	}

	result := h.array[0]
	l := h.lastIndex()
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	if len(h.array) > 1 {
		h.siftDown()
	}
	return result
}

func (h *MaxHeap) siftUp(i int) {
	for i > 0 && h.array[i] > h.array[h.parent(i)] {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *MaxHeap) siftDown() {
	i := 0
	for i < h.lastIndex() {
		l, r := h.left(i), h.right(i)
		if l < 0 && r < 0 {
			return
		}

		// choose the bigger between left or right
		c := r
		if r < 0 || h.array[l] > h.array[r] {
			c = l
		}

		if h.array[c] > h.array[i] {
			h.swap(c, i)
			i = c
		} else {
			return
		}
	}
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) left(i int) int {
	left := (i * 2) + 1
	if left > h.lastIndex() {
		return -1
	}
	return left
}

func (h *MaxHeap) right(i int) int {
	right := (i * 2) + 2
	if right > h.lastIndex() {
		return -1
	}
	return right
}

func (h *MaxHeap) lastIndex() int {
	return len(h.array) - 1
}

func (h *MaxHeap) String() string {
	return fmt.Sprintf("%q", h.array)
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	heap := NewMaxHeap([]int{12, 5, 1, 8, 15, 2, 7, 3, 4, 10})

	for i := 0; i < 12; i++ {
		fmt.Printf("%d, ", heap.array)
		fmt.Printf("My Max is: %d\n", heap.Max())
	}
}
