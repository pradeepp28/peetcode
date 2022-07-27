package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}

	l := dummyHead
	r := head
	for n > 0 {
		r = r.Next
		n = n - 1
	}

	for r != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next
	head = dummyHead.Next
	return head
}

func main() {

}
