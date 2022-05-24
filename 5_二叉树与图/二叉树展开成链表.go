package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
构思

 a
b c

只需要
b.right = c
a.right = b

无b
不用做

无c
a.right = b


从底部开始，对每个节点都进行上面的操作即可


解题反思: 思路没有问题， 需要细节考虑更加全面，争取一把过

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	nodes := make([]*TreeNode, 0)

	all(root, &nodes)

	for i := len(nodes) - 1; i >= 0; i-- {
		exec(nodes[i])
	}
}

func all(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	*nodes = append(*nodes, root)
	start := 0
	end := 1

	for start < end {
		for i := start; i < end; i++ {
			if (*nodes)[i].Left != nil {
				*nodes = append(*nodes, (*nodes)[i].Left)
				end++
			}
			if ((*nodes)[i]).Right != nil {
				*nodes = append(*nodes, (*nodes)[i].Right)
				end++
			}
		}
		start = end
	}
}

//移动节点
func exec(root *TreeNode) {
	//left 为空则不处理
	if root == nil || root.Left == nil {
		return
	}
	//只有left
	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
	} else {
		var rootLeft *TreeNode = root.Left

		for rootLeft.Right != nil {
			rootLeft = rootLeft.Right
		}
		rootLeft.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}
