package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	pos := head
	visited := map[*ListNode]bool{}
	for pos != nil && pos.Next != nil {

		if _, ok := visited[pos]; ok {
			return true
		}

		visited[pos] = true
		pos = pos.Next
	}

	return false
}
