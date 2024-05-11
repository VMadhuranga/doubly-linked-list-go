package main

type node struct {
	data any
	next *node
	prev *node
}

type doublyLinkedList struct {
	head, tail *node
	size       int
}

func (dl *doublyLinkedList) pushToHead(data any) {
	newNode := &node{
		data: data,
		next: nil,
		prev: nil,
	}

	if dl.head == nil {
		dl.tail = newNode
	} else {
		newNode.next = dl.head
	}

	dl.head = newNode
	dl.size++
}

func (dl *doublyLinkedList) popFromHead() *node {
	if dl.head == nil {
		return &node{}
	}

	node := dl.head
	dl.head = dl.head.next
	dl.head.prev = nil
	dl.size--

	return node
}

func (dl *doublyLinkedList) pushToTail(data any) {
	newNode := &node{
		data: data,
		next: nil,
		prev: nil,
	}

	if dl.head == nil {
		dl.head = newNode
	} else {
		newNode.prev = dl.tail
		dl.tail.next = newNode
	}

	dl.tail = newNode
	dl.size++
}

func (dl *doublyLinkedList) popFromTail() *node {
	if dl.head == nil {
		return &node{}
	}

	node := dl.tail
	dl.tail = dl.tail.prev
	dl.tail.next = nil
	dl.size--

	return node
}

func (dl *doublyLinkedList) getNodeAtPosition(position int) *node {
	if position >= dl.size || position < 0 {
		panic("position out of range")
	}

	i := 0
	currentNode := dl.head
	for currentNode != nil {
		if i == position {
			break
		}
		currentNode = currentNode.next
		i++
	}

	return currentNode
}

func (dl *doublyLinkedList) pushToPosition(data any, position int) {
	nodeAtPosition := dl.getNodeAtPosition(position)

	newNode := &node{
		data: data,
		next: nil,
		prev: nil,
	}

	nodeAtPosition.prev.next = newNode
	newNode.next = nodeAtPosition
	nodeAtPosition.prev = newNode

	dl.size++
}

func (dl *doublyLinkedList) popFromPosition(position int) *node {
	nodeAtPosition := dl.getNodeAtPosition(position)

	nodeAtPosition.prev.next = nodeAtPosition.next
	nodeAtPosition.next.prev = nodeAtPosition.prev

	nodeAtPosition.next = nil
	nodeAtPosition.prev = nil

	dl.size--

	return nodeAtPosition
}

func (dl *doublyLinkedList) getSize() int {
	return dl.size
}
