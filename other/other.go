package other

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{Val: 0, Next: nil}
	tmp := res
	flag := 0
	for l1 != nil || l2 != nil {
		tmp.Next = &ListNode{Val: 0, Next: nil}
		tmp = tmp.Next
		if l1 != nil {
			tmp.Val = tmp.Val + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val = tmp.Val + l2.Val
			l2 = l2.Next
		}

		if flag == 1 {
			tmp.Val = tmp.Val + 1
			flag = 0
		}
		if tmp.Val >= 10 {
			flag = 1
			tmp.Val = tmp.Val - 10
		}
	}

	if flag == 1 {
		tmp.Next = &ListNode{Val: 0, Next: nil}
		tmp = tmp.Next
		tmp.Val = tmp.Val + 1
	}
	return res.Next
}

func palindromelen(s string, m int) int {
	f := 1
	l := m - 1
	r := m + 1
	for l > 0 && r < len(s) {
		if s[l] == s[r] {
			f = f + 1
		} else {
			break
		}
		l = l - 1
		r = r + 1
	}
	return f

}

func longestPalindrome(s string) string {
	resi := 0
	rmax := 0
	for i, _ := range s {
		tmp := palindromelen(s, i)
		if tmp > rmax {
			rmax = tmp
			resi = i
		}
	}
	return s[resi : resi+rmax]
}
