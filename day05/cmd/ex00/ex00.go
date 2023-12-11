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

func areToysBalanced(t *tree.Tree) bool {
	if t.Root == nil {
		return false
	}

	return countToys(t.Root.Left) == countToys(t.Root.Right)
}

//func generateRandomTree(nodes int) *tree.Tree {
//	var generatedTree tree.Tree
//	for i := 0; i < nodes; i++ {
//		generatedTree.InsertRandomly(rand.Intn(2) == 0)
//	}
//	return &generatedTree
//}
//
//func main() {
//	checkTree := generateRandomTree(5)
//	fmt.Println(AreToysBalanced(checkTree))
//	checkTree.Print()
//}
