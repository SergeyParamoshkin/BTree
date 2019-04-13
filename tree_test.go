package main

import (
	"testing"
)

func TestTree(t *testing.T) {
	testCases := []int64{5, 2, 10}
	var tree BTree
	for _, node := range testCases {
		tree.Insert(node)
	}
	if tree.root.left.data != 2 {
		t.Fail()
	}
	if tree.root.right.data != 10 {
		t.Fail()
	}
}
