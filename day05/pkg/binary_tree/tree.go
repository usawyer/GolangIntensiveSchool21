package tree

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Root *TreeNode
}

func (t *Tree) AreToysBalanced() bool {
	if t.Root == nil {
		return false
	}

	return t.Root.Left.countToys() == t.Root.Right.countToys()
}

func (t *Tree) UnrollGarland() []bool {
	if t.Root == nil {
		return []bool{}
	}

	return t.Root.levelOrderTraversal()
}

func (t *Tree) Insert(value bool) {
	if t.Root == nil {
		t.Root = &TreeNode{HasToy: value}
	}
}

func (t *Tree) InsertRandomly(value bool) {
	if t.Root == nil {
		t.Root = &TreeNode{HasToy: value}
	} else {
		t.Root.insertRandomly(value)
	}
}

func (t *Tree) Print() {
	if t.Root == nil {
		fmt.Printf("Tree is empty!")
	} else {
		t.Root.print("")
	}
}

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func (n *TreeNode) countToys() int {
	if n == nil {
		return 0
	}
	if !n.HasToy {
		return n.Left.countToys() + n.Right.countToys()
	} else {
		return n.Left.countToys() + n.Right.countToys() + 1
	}
}

func (n *TreeNode) levelOrderTraversal() []bool {
	var result []bool
	var queue []*TreeNode
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

func (n *TreeNode) InsertLeft(value bool) {
	n.insert(value, true)
}

func (n *TreeNode) InsertRight(value bool) {
	n.insert(value, false)
}

func (n *TreeNode) insert(value bool, isLeft bool) {
	if n == nil {
		return
	}
	if isLeft {
		if n.Left == nil {
			n.Left = &TreeNode{HasToy: value}
		} else {
			n.Left.insert(value, isLeft)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{HasToy: value}
		} else {
			n.Right.insert(value, isLeft)
		}
	}
}

func (n *TreeNode) insertRandomly(value bool) {
	if n == nil {
		return
	}
	if rand.Intn(2) == 0 {
		if n.Left == nil {
			n.Left = &TreeNode{HasToy: value}
		} else {
			n.Left.insertRandomly(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{HasToy: value}
		} else {
			n.Right.insertRandomly(value)
		}
	}
}

func (n *TreeNode) print(indent string) {
	if n == nil {
		return
	}

	if n.HasToy {
		fmt.Printf("🎁\n")
	} else {
		fmt.Printf("⚫ \n")
	}

	if n.Left != nil {
		fmt.Printf("%s|-- Left:", indent)
		n.Left.print(indent + "    ")
	}
	if n.Right != nil {
		fmt.Printf("%s|-- Right:", indent)
		n.Right.print(indent + "    ")
	}
}
