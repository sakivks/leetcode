package main

import "fmt"

func main() {
	tree := &TreeNode{4,
		&TreeNode{2,
			&TreeNode{1, nil, nil},
			&TreeNode{3, nil, nil}},
		&TreeNode{7,
			&TreeNode{6, nil, nil},
			&TreeNode{9, nil, nil}}}

	// fmt.Println(preorderTraversal(tree))
	fmt.Println(invertTree(tree))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// func preorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}

// 	result := []int{}
// 	result = append(result, root.Val)
// 	if root.Left != nil {
// 		result = append(result, preorderTraversal(root.Left)...)
// 	}
// 	if root.Right != nil {
// 		result = append(result, preorderTraversal(root.Right)...)
// 	}

// 	return result
// }
