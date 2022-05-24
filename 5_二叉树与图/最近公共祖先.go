package main

/*
leetcode 236 middle

question: 已知一个二叉树， 求两个已知节点的公共祖先

构思: 向上追溯出两个节点的路径， 然后再求交点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pathMap = make(map[*TreeNode][]*TreeNode)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//get two path list
	newList := make([]*TreeNode, 0)
	newList = append(newList, root)
	getPathToRoot(root, newList)

	pList := make([]*TreeNode, 0)
	if p != nil {
		pList = pathMap[p]
	} else {
		pList = newList
	}
	qList := make([]*TreeNode, 0)
	if q != nil {
		qList = pathMap[q]
	} else {
		qList = newList
	}
	//get common node, last same node
	var ret *TreeNode
	for i := 0; i < len(qList) && i < len(pList); i++ {
		if qList[i] == pList[i] {
			ret = qList[i]
		} else {
			break
		}
	}

	return ret
}

func getPathToRoot(root *TreeNode, list []*TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		listLeft := make([]*TreeNode, 0)
		listLeft = append(listLeft, list...)
		listLeft = append(listLeft, root.Left)
		pathMap[root.Left] = listLeft
		getPathToRoot(root.Left, listLeft)
	}
	if root.Right != nil {
		listRight := make([]*TreeNode, 0)
		listRight = append(listRight, list...)
		listRight = append(listRight, root.Right)
		pathMap[root.Right] = listRight
		getPathToRoot(root.Right, listRight)
	}
}

//优化点， 现在是递归存储了所有的路径， 如果使用栈来进行先序遍历， 就可以 now left right pop 遍历到想要的结果就立刻pop

func main() {

}
