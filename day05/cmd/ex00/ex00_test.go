package ex00

import (
	tree "day05/internal/binary_tree"
	"testing"
)

func TestCase00(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Insert(false)
	treeTest.Root.InsertLeft(false)
	treeTest.Root.Left.InsertLeft(false)
	treeTest.Root.Left.InsertRight(true)
	treeTest.Root.InsertRight(true)

	treeTest.Print()

	expected := true
	actual := areToysBalanced(treeTest)

	if expected != actual {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestCase01(t *testing.T) {
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
	actual := areToysBalanced(treeTest)

	if expected != actual {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestCase02(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Insert(true)
	treeTest.Root.InsertLeft(true)
	treeTest.Root.InsertRight(false)

	treeTest.Print()

	expected := false
	actual := areToysBalanced(treeTest)

	if expected != actual {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestCase03(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Insert(false)
	treeTest.Root.InsertLeft(true)
	treeTest.Root.Left.InsertRight(true)
	treeTest.Root.InsertRight(false)
	treeTest.Root.Right.InsertRight(true)

	treeTest.Print()

	expected := false
	actual := areToysBalanced(treeTest)

	if expected != actual {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}
