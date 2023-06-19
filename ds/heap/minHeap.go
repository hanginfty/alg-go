package heap

type MinHeap struct {
	items []int
}

func (h *MinHeap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *MinHeap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *MinHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *MinHeap) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < len(h.items)
}

func (h *MinHeap) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < len(h.items)
}

func (h *MinHeap) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

func (h *MinHeap) leftChild(index int) int {
	return h.items[h.getLeftChildIndex(index)]
}

func (h *MinHeap) rightChild(index int) int {
	return h.items[h.getRightChildIndex(index)]
}

func (h *MinHeap) parent(index int) int {
	return h.items[h.getParentIndex(index)]
}

func (h *MinHeap) swap(indexOne, indexTwo int) {
	temp := h.items[indexOne]
	h.items[indexOne] = h.items[indexTwo]
	h.items[indexTwo] = temp
}

func (h *MinHeap) peek() int {
	if len(h.items) == 0 {
		panic("Heap is empty")
	}
	return h.items[0]
}

func (h *MinHeap) poll() int {
	if len(h.items) == 0 {
		panic("Heap is empty")
	}
	item := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.heapifyDown()
	return item
}

func (h *MinHeap) add(item int) {
	h.items = append(h.items, item)
	h.heapifyUp()
}

func (h *MinHeap) heapifyUp() {
	index := len(h.items) - 1
	for h.hasParent(index) && h.parent(index) > h.items[index] {
		h.swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}
}

func (h *MinHeap) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		smallerChildIndex := h.getLeftChildIndex(index)
		if h.hasRightChild(index) && h.rightChild(index) < h.leftChild(index) {
			smallerChildIndex = h.getRightChildIndex(index)
		}
		if h.items[index] < h.items[smallerChildIndex] {
			break
		} else {
			h.swap(index, smallerChildIndex)
		}
		index = smallerChildIndex
	}
}
