package tree

/*
给定一个二叉树和sum， 求所有根节点到叶子节点的路径之和， 使得路径之和恰好等于sum
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	generate(root, targetSum, path, &ret)
	return ret
}

//递归处理
func generate(root *TreeNode, targetSum int, path []int, ret *[][]int) {
	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil && sumArr(path) == targetSum {
		*ret = append(*ret, path)
		return
	}
	if root.Left != nil {
		generate(root.Left, targetSum, path, ret)
	}
	if root.Right != nil {
		generate(root.Right, targetSum, path, ret)
	}
}

func sumArr(path []int) int {
	sum := 0
	for _, item := range path {
		sum += item
	}
	return sum
}
