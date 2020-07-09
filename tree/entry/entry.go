package main

import (
	"fmt"
	"learnGo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

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

	fmt.Println("\nready Traverse")
	root.Traverse()

	fmt.Println("\nready postOrder")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	count := 0
	root.TraverseFunc(func(node *tree.Node) {
		count++
	})
	fmt.Println("Count: ", count)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("maxNode: ", maxNode)
}
