package ex01

import (
	tree "day05/internal/binary_tree"
)

func levelOrderTraversal(n *tree.TreeNode) []bool {
	var result []bool
	var queue []*tree.TreeNode
	queue = append(queue, n)
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		values := make([]bool, size)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 != 0 {
				values[i] = node.HasToy
			} else {
				values[size-i-1] = node.HasToy
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, values...)
		level++
	}
	return result
}

func unrollGarland(t *tree.Tree) []bool {
	if t.Root == nil {
		return []bool{}
	}

	return levelOrderTraversal(t.Root)
}
