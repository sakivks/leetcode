package main

import "fmt"

func main() {
	tree1 := &TreeNode{1,
		&TreeNode{11, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			nil}}
	tree2 := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			nil,
			&TreeNode{7, nil, nil}}}

	fmt.Println(mergeTrees(tree1, tree2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{root1.Val + root2.Val, mergeTrees(root1.Left, root2.Left), mergeTrees(root1.Right, root2.Right)}
}
