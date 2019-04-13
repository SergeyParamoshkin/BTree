package main

// BTree type tree
type BTree struct {
	root *node
}

// node
type node struct {
	left  *node
	right *node
	data  int64
}

// Insert method on tree
func (t *BTree) Insert(data int64) {
	if t.root == nil {
		t.root = &node{data: data}
	} else {
		t.root.insert(data)
	}
}

/*
Insert method add value
X it is root item
if K > X, right sub tree
if K < X, lrft sub tree
if K = X, replace current V node
*/
func (n *node) insert(data int64) {
	if data <= n.data {
		if n.left == nil {
			n.left = &node{data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{data: data}
		} else {
			n.right.insert(data)
		}
	}
}
