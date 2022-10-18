package list

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
	len  int
}

func NewList(items []int) List {
	var p *Node = nil
	for i := len(items) - 1; i >= 0; i-- {
		newNode := Node{items[i], p}
		p = &newNode
	}

	return List{
		head: p,
		len:  len(items),
	}
}
