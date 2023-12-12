package addTwoNumbers

type ListNode struct {
	Val		int
	Next	*ListNode
}

func makeList(input []int) *ListNode {

	head := &ListNode{input[0], nil}
	tail := head

	for _, digit := range input[1:] {		
		node := &ListNode{digit, nil}
		tail.Next = node
		tail = node
	}

	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res []int

	mem := 0

	for l1 != nil || l2 != nil {
		sum := 0
		
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		sum = (sum + mem) % 10
		mem = (sum + mem) / 10

		res = append(res, sum)
	}
	return makeList(res)
}