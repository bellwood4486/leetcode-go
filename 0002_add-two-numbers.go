package leetcode_go

type ListNode struct {
	Val  int
	Next *ListNode
}

var nilNode = &ListNode{
	Val:  0,
	Next: nil,
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	outputBegin := &ListNode{}
	outputCursor := outputBegin

	l1Cursor, l2Cursor := l1, l2
	moveUp := false
	for {
		v := l1Cursor.Val + l2Cursor.Val
		if moveUp {
			moveUp = false
			v += 1
		}
		if v >= 10 {
			moveUp = true
			v %= 10
		}

		outputCursor.Val = v

		if !moveUp && l1Cursor.Next == nil && l2Cursor.Next == nil {
			break
		}

		outputCursor.Next = &ListNode{}
		outputCursor = outputCursor.Next

		l1Cursor, l2Cursor = l1Cursor.Next, l2Cursor.Next
		if l1Cursor == nil {
			l1Cursor = nilNode
		}
		if l2Cursor == nil {
			l2Cursor = nilNode
		}
	}

	return outputBegin
}
