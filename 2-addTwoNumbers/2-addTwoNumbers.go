package addTwoNumbers

type ListNode struct {
	Val		int
	Next	*ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	var temp *ListNode = head

	mem := 0

	for l1 != nil || l2 != nil {
		sum := 0
		
		if l1 != nil {
			sum = sum + l1.Val
		}

		if l2 != nil {
			sum = sum + l2.Val
		} else {
			
		}

		sum = (sum + mem) % 10
		mem = (sum + mem) / 10

		temp.Val = sum
		temp = temp.Next

		l1 = l1.Next
		l2 = l2.Next
	}
	return head
}