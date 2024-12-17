package main

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func Insert(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return &TreeNode{Data: data}
	}

	if data <= node.Data {
		node.Left = Insert(node.Left, data)
	} else {
		node.Right = Insert(node.Right, data)
	}
	return node
}

func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PreOrder(node.Left)
	fmt.Println(node.Data)
	PreOrder(node.Right)
}

func PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PreOrder(node.Left)
	PreOrder(node.Right)
	fmt.Println(node.Data)
}

func InvertTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left, node.Right = InvertTree(node.Right), InvertTree(node.Left)
	return node
}
