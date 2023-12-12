package test

import (
	tree "day05/pkg/binary_tree"
	"testing"
)

func TestTreeBalanced(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Insert(false)
	treeTest.Root.InsertLeft(false)
	treeTest.Root.Left.InsertLeft(false)
	treeTest.Root.Left.InsertRight(true)
	treeTest.Root.InsertRight(true)

	treeTest.Print()

	expected := true
	actual := treeTest.AreToysBalanced()

	if expected != actual {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
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

	treeTest.Print()

	expected := true
	actual := treeTest.AreToysBalanced()

	if expected != actual {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestTreeUnbalanced(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Insert(true)
	treeTest.Root.InsertLeft(true)
	treeTest.Root.InsertRight(false)

	treeTest.Print()

	expected := false
	actual := treeTest.AreToysBalanced()

	if expected != actual {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestTreeUnbalanced01(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Insert(false)
	treeTest.Root.InsertLeft(true)
	treeTest.Root.Left.InsertRight(true)
	treeTest.Root.InsertRight(false)
	treeTest.Root.Right.InsertRight(true)

	treeTest.Print()

	expected := false
	actual := treeTest.AreToysBalanced()

	if expected != actual {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestTreeBalancedEmpty(t *testing.T) {
	treeTest := &tree.Tree{}

	expected := false
	actual := treeTest.AreToysBalanced()

	if expected != actual {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}
