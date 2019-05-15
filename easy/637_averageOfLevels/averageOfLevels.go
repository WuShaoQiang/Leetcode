package averageOfLevels

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.

示例 1:

输入:
    3
   / \
  9  20
    /  \
   15   7
输出: [3, 14.5, 11]
解释:
第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].

注意：

    节点值的范围在32位有符号整数范围内。
*/

// 执行用时 : 28 ms, 在Average of Levels in Binary Tree的Go提交中击败了97.44% 的用户
// 内存消耗 : 8.1 MB, 在Average of Levels in Binary Tree的Go提交中击败了9.09% 的用户

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	res := make([]float64, 0)
	// 将第一个点放入队列
	quene := []*TreeNode{root}

	for {
		// 确定这一次遍历的个数
		l := len(quene)
		sum := 0
		// 因为下面有判断quene长度,因此这里不可能为0,所以不用怕除数为0
		count := l
		// 当遍历个数用完的时候表明这一层遍历完了,应当退出循环重新确定下一层
		for l > 0 {
			tmp := quene[l-1]
			sum += tmp.Val
			// 在遍历的时候添加下一层的节点
			if tmp.Left != nil {
				quene = append(quene, tmp.Left)
			}
			if tmp.Right != nil {
				quene = append(quene, tmp.Right)
			}

			l--
		}
		// 记录
		res = append(res, float64(sum)/float64(count))
		// 将遍历过的节点去除
		quene = quene[count:]

		// 如果没有新的节点产生,说明到尾了
		if len(quene) == 0 {
			break
		}
	}
	return res
}
