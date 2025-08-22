package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		child := BTreeInsertData(root.Left, data)
		root.Left = child
		root.Parent = root
	} else {
		child := BTreeInsertData(root.Right, data)
		root.Right = child
		root.Parent = root
	}
	return root
}
