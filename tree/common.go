package tree

const (
	Null = -1
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Arr2Tree 输入一个层次遍历顺序二叉树对应的整数数组
// 返回构造出的二叉树
// 值为 -1 的元素代表该节点为 nil
// 0 1 2 3 4 5 6
// 1 2 3 4 5 6 7
//
//      1
//   2      3
// 4   5  6   7
//
func Arr2Tree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(arr))
	for i := 0; i < len(nodes); i++ {
		if arr[i] != Null {
			nodes[i] = &TreeNode{Val: arr[i]}
		}
	}

	for i := 0; i < len(arr)/2; i++ {
		if nodes[i] == nil {
			continue
		}
		nodes[i].Left = nodes[i*2+1]
		if i*2+2 < len(arr) {
			nodes[i].Right = nodes[i*2+2]
		}
	}

	return nodes[0]
}
