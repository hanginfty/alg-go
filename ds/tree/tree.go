package tree

type Node struct {
	Val         int
	Left, Right *Node
}

type Tree = *Node

func NewTree(ary []int) Tree {
	capacity := len(ary)
	if capacity == 0 {
		return nil
	}

	nodes := make([]*Node, capacity)

	for i := 0; i < capacity/2; i++ {
		if ary[i] == -1 {
			continue
		}

		var node = nodes[i]
		if nodes[i] == nil {
			node = &Node{ary[i], nil, nil}
			nodes[i] = node
		}

		l := 2*i + 1
		r := 2*i + 2
		if l < capacity && nodes[l] == nil && ary[l] != -1 {
			node.Left = &Node{ary[l], nil, nil}
			nodes[l] = node.Left
		}
		if r < capacity && nodes[r] == nil && ary[r] != -1 {
			node.Right = &Node{ary[r], nil, nil}
			nodes[r] = node.Right
		}
	}

	return nodes[0]
}

func (tree Tree) PreOrder() []int {
	var res []int
	if tree == nil {
		return res
	}

	res = append(res, tree.Val)

	res = append(res, tree.Left.PreOrder()...)
	res = append(res, tree.Right.PreOrder()...)

	return res
}

func (tree Tree) LevelOrder() [][]int {
	if tree == nil {
		return [][]int{}
	}

	var res = [][]int{}
	var q = []*Node{tree}

	for len(q) > 0 {
		size := len(q)
		level := []int{}

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}

func (tree Tree) Height() int {
	if tree == nil {
		return 0
	}

	l := tree.Left.Height()
	r := tree.Right.Height()
	if l > r {
		return 1 + l
	} else {
		return 1 + r
	}
}

func (tree Tree) IsLeaf() bool {
	return tree.Left == nil && tree.Right == nil
}

func (tree Tree) Count() int {
	if tree == nil {
		return 0
	}

	l := tree.Left.Count()
	r := tree.Right.Count()

	return l + r + 1
}

func IsSameTree(t1, t2 Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return IsSameTree(t1.Left, t2.Left) && IsSameTree(t1.Right, t2.Right)
}

func (tree Tree) Flatten() {
	if tree == nil {
		return
	}

	tree.Left.Flatten()
	tree.Right.Flatten()

	l := tree.Left
	r := tree.Right

	tree.Left = nil
	tree.Right = l

	p := tree
	for p.Right != nil {
		p = p.Right
	}
	p.Right = r
}
