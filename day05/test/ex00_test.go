package test

import (
	tree "day05/pkg/binary_tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeBalanced(t *testing.T) {
	treeTest := &tree.Tree{}
	treeTest.Insert(false)
	treeTest.Root.InsertLeft(false)
	treeTest.Root.Left.InsertLeft(false)
	treeTest.Root.Left.InsertRight(true)
	treeTest.Root.InsertRight(true)

	// Uncomment the method to display the tree
	//treeTest.Print()

	assert.Equal(t, true, treeTest.AreToysBalanced())
}

func TestTreeBalanced01(t *testing.T) {
	treeTest := &tree.Tree{}
	treeTest.Insert(true)
	treeTest.Root.InsertLeft(true)
	treeTest.Root.Left.InsertLeft(true)
	treeTest.Root.Left.InsertRight(false)
	treeTest.Root.InsertRight(false)
	treeTest.Root.Right.InsertLeft(true)
	treeTest.Root.Right.InsertRight(true)

	// Uncomment the method to display the tree
	//treeTest.Print()

	assert.Equal(t, true, treeTest.AreToysBalanced())
}

func TestTreeUnbalanced(t *testing.T) {
	treeTest := &tree.Tree{}
	treeTest.Insert(true)
	treeTest.Root.InsertLeft(true)
	treeTest.Root.InsertRight(false)

	// Uncomment the method to display the tree
	//treeTest.Print()

	assert.Equal(t, false, treeTest.AreToysBalanced())
}

func TestTreeUnbalanced01(t *testing.T) {
	treeTest := &tree.Tree{}
	treeTest.Insert(false)
	treeTest.Root.InsertLeft(true)
	treeTest.Root.Left.InsertRight(true)
	treeTest.Root.InsertRight(false)
	treeTest.Root.Right.InsertRight(true)

	// Uncomment the method to display the tree
	//treeTest.Print()

	assert.Equal(t, false, treeTest.AreToysBalanced())
}

func TestTreeBalancedEmpty(t *testing.T) {
	treeTest := &tree.Tree{}
	assert.Equal(t, false, treeTest.AreToysBalanced())
}
