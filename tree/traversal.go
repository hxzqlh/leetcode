package tree

import "container/list"

// InOrderTraversal
// 中序遍历：迭代
// 空间复杂度：O(n)
func InOrderTraversal(root *TreeNode) []int {
	var ret []int

	stack := list.New()
	p := root
	for p != nil || stack.Len() > 0 {
		if p != nil {
			stack.PushBack(p)
			p = p.Left
		} else {
			e := stack.Back()
			stack.Remove(e)
			p = e.Value.(*TreeNode)
			ret = append(ret, p.Val)
			p = p.Right
		}
	}

	return ret
}

// InOrderTraversal2
// 中序遍历：递归
// 空间复杂度：O(n)
func InOrderTraversal2(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}

	traversal(root)
	return
}

// InOrderTraversal3 Morris 遍历算法
// 中序遍历：迭代
// 空间复杂度：O(1)
//
// 假设当前遍历到的节点为 x:
// 1. 如果 x 无左孩子，先将 x 的值加入答案数组，再访问 xx 的右孩子，即 x = x.right。
// 2. 如果 x 有左孩子，则找到 x 左子树上最右的节点（即左子树中序遍历的最后一个节点，x 在中序遍历中的前驱节点）
//    我们记为 predecessor，根据 predecessor 的右孩子是否为空，进行如下操作:
//    a. 如果 predecessor 的右孩子为空，则将其右孩子指向 x，然后访问 x 的左孩子，即 x = x.left。
//    b. 如果 predecessor 的右孩子不为空，则此时其右孩子指向 x，说明我们已经遍历完 x 的左子树，
//       我们将 predecessor 的右孩子置空，将 x 的值加入答案数组，然后访问 x 的右孩子，即 x = x.right。
// 3. 重复上述操作，直至访问完整棵树。
//
// 诀窍：假设当前遍历到的节点为 x，将 x 的左子树中最右边的节点的右孩子指向 x，
// 这样在左子树遍历完成后我们通过这个指向走回了 x，且能通过这个指向知晓我们已经遍历完成了左子树，
// 而不用再通过栈来维护，省去了栈的空间复杂度。
//
func InOrderTraversal3(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			pre := root.Left
			for pre.Right != nil && pre.Right != root {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = root
				root = root.Left
			} else {
				// we know: pre.Right == root
				res = append(res, root.Val)
				pre.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return
}

// PreOrderTraversal
// 前序遍历：递归
// 空间复杂度：O(n)
func PreOrderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return
}

// PostOrderTraversal
// 后序遍历：递归
// 空间复杂度：O(n)
func PostOrderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}

	traversal(root)
	return
}

// LevelOrder 层次遍历
// LeetCode: #102, https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// 二叉树：[3,9,20,null,null,15,7],
//
//   3
//  / \
// 9  20
//   /  \
//  15   7
//
// 返回其层序遍历结果：
//
// [
// 		[3],
// 		[9,20],
// 		[15,7]
// ]
func LevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)

	var (
		node     *TreeNode
		levelArr []int
	)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node = queue.Remove(queue.Front()).(*TreeNode)
			levelArr = append(levelArr, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, levelArr)
		levelArr = []int{} // clear level arr
	}
	return
}

// LevelOrderBottom 节点值自底向上的层序遍历
// LeetCode: #107, https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
// 二叉树：[3,9,20,null,null,15,7],
//
//   3
//  / \
// 9  20
//   /  \
//  15   7
//
// 返回其节点值自底向上的层序遍历结果：
//
// [
// 		[15,7],
// 		[9,20],
// 		[3],
// ]
func LevelOrderBottom(root *TreeNode) (res [][]int) {
	res = LevelOrder(root)
	// reverse
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return
}

// RightSideView 二叉树的右视图
// LeetCode: #199, https://leetcode-cn.com/problems/binary-tree-right-side-view/
// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，
// 按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
// 二叉树：[1,2,3,null,5,null,4],
//
//   1
//  / \
// 2   3
//  \   \
//   5   4
//
// 返回其节点值自底向上的层序遍历结果：
//
// [1,3,4]
func RightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)

	var (
		node     *TreeNode
		levelArr []int
	)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node = queue.Remove(queue.Front()).(*TreeNode)
			levelArr = append(levelArr, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, levelArr[len(levelArr)-1]) // append the last one
		levelArr = []int{}                           // clear level arr
	}
	return
}
