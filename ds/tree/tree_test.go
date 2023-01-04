package tree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	ary := []int{1, 2, 3, 4, -1, 6, 7}

	tree := NewTree(ary)

	fmt.Println(tree.LevelOrder())
}
