package ex01

import (
	tree "day05/internal/binary_tree"
	"reflect"
	"testing"
)

func TestCase00(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Insert(true)
	treeTest.Root.InsertLeft(true)
	treeTest.Root.InsertRight(false)
	treeTest.Root.Left.InsertLeft(true)
	treeTest.Root.Left.InsertRight(false)
	treeTest.Root.Right.InsertLeft(true)
	treeTest.Root.Right.InsertRight(true)

	treeTest.Print()

	expected := []bool{true, true, false, true, true, false, true}
	actual := unrollGarland(treeTest)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestCase01(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Insert(false)
	treeTest.Root.InsertLeft(false)
	treeTest.Root.InsertRight(true)
	treeTest.Root.Left.InsertLeft(true)
	treeTest.Root.Left.InsertRight(true)
	treeTest.Root.Left.Left.InsertLeft(true)
	treeTest.Root.Left.Left.InsertRight(false)
	treeTest.Root.Left.Right.InsertLeft(true)
	treeTest.Root.Left.Right.InsertRight(false)
	treeTest.Root.Left.Right.Left.InsertRight(true)

	treeTest.Print()

	expected := []bool{false, false, true, true, true, true, false, true, false, true}
	actual := unrollGarland(treeTest)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestCase02(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Print()

	expected := make([]bool, 0)
	actual := unrollGarland(treeTest)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}
