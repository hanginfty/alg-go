package tree

import "sync"

type BSTNode struct {
	key         int
	value       int
	left, right *BSTNode
}

type BST struct {
	root *BSTNode
	lock sync.RWMutex
}

func NewBST() *BST {
	return &BST{}
}

func (bst *BST) Insert(key int, value int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	bst.root = bst.insert(bst.root, key, value)
	if bst.root == nil {
		return
	}
}

func (bst *BST) insert(node *BSTNode, key int, value int) *BSTNode {
	if node == nil {
		return &BSTNode{key, value, nil, nil}
	}
	// return bst.insert(node.left, key, value)
	return nil
}
