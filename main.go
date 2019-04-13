package main

import (
	"fmt"
)

func main() {
	fmt.Println("run BTree")
	var tree BTree
	tree.Insert(0)
	fmt.Println(tree)
}
