package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreOrder(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewTree(ary)
	expected := []int{1, 2, 4, 5, 3, 6, 7}
	assert.Equal(t, expected, tree.PreOrder())
}

func TestLevelOrder(t *testing.T) {
	ary := []int{1, 2, 3, 4, -1, 6, 7}
	tree := NewTree(ary)

	fmt.Println(tree.LevelOrder())
}

func TestCount(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewTree(ary)
	expected := 7
	assert.Equal(t, expected, tree.Count())
}

func TestFlatten(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewTree(ary)

	tree.Flatten()
	fmt.Println(tree.LevelOrder())
}
