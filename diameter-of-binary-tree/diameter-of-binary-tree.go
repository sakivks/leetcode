package main

import "fmt"

func main() {
	tree := &TreeNode{4,
		&TreeNode{2,
			&TreeNode{1, nil, nil},
			&TreeNode{3, nil, nil}},
		&TreeNode{7, nil, nil}}

	fmt.Println(diameterOfBinaryTree(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	height(root, &max)
	return max
}

func height(head *TreeNode, max *int) int {
	if head == nil {
		return 0
	}
	leftHeight := height(head.Left, max)
	rightHeight := height(head.Right, max)

	if *max < leftHeight+rightHeight {
		*max = leftHeight + rightHeight
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

//basic solution with array of nodes
func diameterOfBinaryTree_basic(root *TreeNode) int {
	nodes := inorderTraversal(root)
	diameters, max := make([]int, len(nodes)), 0

	for i, node := range nodes {
		diameters[i] = height_basic(node.Left) + height_basic(node.Right)
	}

	for _, dia := range diameters {
		if dia > max {
			max = dia
		}
	}

	return max
}

func inorderTraversal(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}

	result := []*TreeNode{}
	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}
	result = append(result, root)
	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}

	return result
}

func height_basic(head *TreeNode) int {
	if head == nil {
		return 0
	}
	leftHeight := height_basic(head.Left)
	rightHeight := height_basic(head.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
