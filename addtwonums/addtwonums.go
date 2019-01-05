package addtwonums

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var newlst *ListNode
	carry := 0
	sum := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}
		sum = sum + carry
		carry = sum - 10
		tempnode := &ListNode{
			Val:  sum % 10,
			Next: newlst,
		}
		newlst = tempnode
	}
	return newlst

}
