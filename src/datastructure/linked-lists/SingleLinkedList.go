package linked_lists

type Node struct {
	data int
	next *Node
}

func insert(root *Node, item int) *Node {
	if root == nil {
		root = &Node{data: item}
	} else {
		root.next = insert(root.next, item)
	}
	return root
}

func insertM(root *Node, items []int) *Node {
	for _, v := range items {
		root = insert(root, v)
	}
	return root
}

func getElements(first *Node) []int {
	var elements []int
	traverser := first
	for ; traverser != nil; traverser = traverser.next {
		elements = append(elements, traverser.data)
	}
	return elements
}

func deleteItem(root *Node, i int) *Node {
	if root == nil {
	} else if root.data == i {
		root = nil
	} else if root.next != nil && root.next.data == i {
		root.next = root.next.next
	} else {
		deleteItem(root.next, i)
	}
	return root
}

func reverse(first *Node) *Node {
	traverser := first
	var middle *Node = nil
	var previous *Node = nil
	for ; traverser != nil; {
		previous = middle
		middle = traverser
		traverser = traverser.next
		middle.next = previous
	}
	return middle
}
