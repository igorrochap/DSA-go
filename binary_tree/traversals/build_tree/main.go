package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderIndexMap := make(map[int]int)
	for i, val := range inorder {
		inorderIndexMap[val] = i
	}
	rootIndex := len(postorder) - 1
	return build(inorder, postorder, 0, len(inorder)-1, &rootIndex, inorderIndexMap)
}

func build(inorder, postorder []int, inorderStart, inorderEnd int, rootIndex *int, inorderMap map[int]int) *TreeNode {
	if inorderStart > inorderEnd || *rootIndex < 0 {
		return nil
	}
	rootValue := postorder[*rootIndex]
	*rootIndex--

	root := &TreeNode{Val: rootValue}

	inorderIndex := inorderMap[rootValue]

	root.Right = build(inorder, postorder, inorderIndex+1, inorderEnd, rootIndex, inorderMap)
	root.Left = build(inorder, postorder, inorderStart, inorderIndex-1, rootIndex, inorderMap)
	return root
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	buildTree(inorder, postorder)
}
