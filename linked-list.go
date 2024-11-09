package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
	Size int
}

func (l *List) Add(node *Node) {
	if l.Head == nil {
		l.Head = node
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	l.Size++
}

func (l *List) GetList() []int {
	var result []int
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func (l *List) Delete(node *Node) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == node.Value {
		l.Head = l.Head.Next
		l.Size--
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Value == node.Value {
			current.Next = current.Next.Next
			l.Size--
			return
		}
		current = current.Next
	}
}

func (l *List) Get(node *Node) *Node {
	current := l.Head
	for current != nil {
		if current.Value == node.Value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *List) DeleteByIndex(index int) {
	if index < 0 || index >= l.Size {
		return
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Size--
		return
	}
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	l.Size--
}

func (l *List) GetByIndex(index int) *Node {
	if index < 0 || index >= l.Size {
		return nil
	}
	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current
}

func (l *List) SwapPairs() {
	l.Head = swapPairs(l.Head)
}

func swapPairs(head *Node) *Node {
	dummy := &Node{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {
		firstNode := head
		secondNode := head.Next

		prev.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		prev = firstNode
		head = firstNode.Next
	}
	return dummy.Next
}

func (l *List) ReverseKGroup(k int) {
	l.Head = reverseKGroup(l.Head, k)
}

func reverseKGroup(head *Node, k int) *Node {
	count := 0
	current := head

	for current != nil && count != k {
		current = current.Next
		count++
	}

	if count == k {
		current = reverseKGroup(current, k)
		for count > 0 {
			temp := head.Next
			head.Next = current
			current = head
			head = temp
			count--
		}
		head = current
	}
	return head
}

func main() {
	list := &List{}
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}
	node5 := &Node{Value: 5}
	node6 := &Node{Value: 6}
	node7 := &Node{Value: 7}

	// Adding nodes to the list
	list.Add(node1)
	list.Add(node2)
	list.Add(node3)
	list.Add(node4)
	list.Add(node5)
	list.Add(node6)
	list.Add(node7)
	fmt.Println("Initial List:", list.GetList())

	// Swap pairs in the list
	list.SwapPairs()
	fmt.Println("List after SwapPairs:", list.GetList())

	// Reverse nodes in groups of 3
	list.ReverseKGroup(3)
	fmt.Println("List after ReverseKGroup (k=3):", list.GetList())

	// Get the value of a specific node
	node := list.Get(node4)
	fmt.Printf("Node with value %d found\n", node.Value)

	// Get node by index
	nodeByIndex := list.GetByIndex(2)
	fmt.Printf("Node at index 2 has value: %d\n", nodeByIndex.Value)

	// Delete a node
	list.Delete(node4)
	fmt.Println("List after deleting node with value 4:", list.GetList())

	// Delete node by index
	list.DeleteByIndex(2)
	fmt.Println("List after deleting node at index 2:", list.GetList())

	// Final list size
	fmt.Println("Final List Size:", list.Size)
}
