package ex00

import (
	tree "day05/internal/binary_tree"
)

func countToys(node *tree.TreeNode) int {
	if node == nil {
		return 0
	}
	if !node.HasToy {
		return countToys(node.Left) + countToys(node.Right)
	} else {
		return countToys(node.Left) + countToys(node.Right) + 1
	}
}

func AreToysBalanced(t *tree.Tree) bool {
	if t.Root == nil {
		return false
	}

	return countToys(t.Root.Left) == countToys(t.Root.Right)
}
