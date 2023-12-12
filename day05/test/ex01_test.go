package test

import (
	"day05/cmd/ex01"
	tree "day05/internal/binary_tree"
	"reflect"
	"testing"
)

func TestTreeEasy(t *testing.T) {
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
	actual := ex01.UnrollGarland(treeTest)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestTreeEasy01(t *testing.T) {
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
	actual := ex01.UnrollGarland(treeTest)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestTreeEmpty(t *testing.T) {
	treeTest := &tree.Tree{}

	treeTest.Print()

	expected := make([]bool, 0)
	actual := ex01.UnrollGarland(treeTest)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
}
