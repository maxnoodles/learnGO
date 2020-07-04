package main

import (
	"fmt"
	"learnGo/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, nil},
	}
	root.Left = &nodes[0]
	root.Left.Left = new(tree.Node)

	fmt.Println(root)
	fmt.Println(nodes)

	root.Right.Left = tree.CreateNode(7)
	root.Print()
	root.SetValue(100)
	root.Print()
	root.Left.Right.SetValue(6)
	root.Traverse()
}
