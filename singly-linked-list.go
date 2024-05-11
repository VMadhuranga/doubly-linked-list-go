package main

import (
	"fmt"
)

type singlyLinkedListNode struct {
	data int
	next *singlyLinkedListNode
}

type singlyLinkedList struct {
	head *singlyLinkedListNode
	tail *singlyLinkedListNode
	size int
}

func (sList *singlyLinkedList) insertNodeAtTail(nodeData int) {
	node := &singlyLinkedListNode{
		data: nodeData,
		next: nil,
	}

	if sList.head == nil {
		sList.head = node
	} else {
		sList.tail.next = node
	}

	sList.tail = node
	sList.size++
}

func (sList *singlyLinkedList) insertNodeAtHead(nodeData int) {
	node := &singlyLinkedListNode{
		data: nodeData,
		next: nil,
	}

	if sList.head == nil {
		sList.head = node
		sList.tail = node
	} else {
		node.next = sList.head
		sList.head = node
	}

	sList.size++
}

func (sList *singlyLinkedList) insertNodeAtPosition(nodeData, position int) {
	if sList.head == nil {
		fmt.Println("Empty list")
	} else if position >= sList.size {
		fmt.Println("Cannot insert, position out of range")
	} else if position == 0 {
		sList.insertNodeAtHead(nodeData)
	} else {
		node := &singlyLinkedListNode{
			data: nodeData,
			next: nil,
		}

		currentNode := sList.head

		for i := 0; i < position; i++ {
			if i+1 == position {
				node.next = currentNode.next
				currentNode.next = node
				break
			}

			currentNode = currentNode.next
		}

		sList.size++
	}
}

func (sList *singlyLinkedList) deleteNodeAtHead() {
	if sList.head == nil {
		fmt.Println("Empty list")
	} else {
		sList.head = sList.head.next
		if sList.size == 1 {
			sList.tail = sList.tail.next
		}
		sList.size--
	}
}

func (sList *singlyLinkedList) deleteNodeAtTail() {
	if sList.head == nil {
		fmt.Println("Empty list")
	} else if sList.size == 1 {
		sList.tail = sList.tail.next
		sList.head = sList.head.next
		sList.size--
	} else {
		currentNode := sList.head

		for i := 0; i < sList.size; i++ {
			if i+2 == sList.size {
				currentNode.next = currentNode.next.next
				sList.tail = currentNode
				sList.size--
				break
			}

			currentNode = currentNode.next
		}
	}
}

func (sList *singlyLinkedList) deleteNodeAtPosition(position int) {
	if sList.head == nil {
		fmt.Println("Empty list")
	} else if position >= sList.size {
		fmt.Println("Cannot delete, position out of range")
	} else if position == 0 {
		sList.deleteNodeAtHead()
	} else if position == sList.size-1 {
		sList.deleteNodeAtTail()
	} else {
		currentNode := sList.head

		for i := 0; i < position; i++ {
			if i+1 == position {
				currentNode.next = currentNode.next.next
				sList.size--
				break
			}

			currentNode = currentNode.next
		}
	}
}

func (sList singlyLinkedList) printNodeAtPosition(position int) {
	if sList.head == nil {
		fmt.Println("Empty list")
	} else if position >= sList.size {
		fmt.Println("Cannot print data, position out of range")
	} else {
		currentNode := sList.head

		for i := 0; i < sList.size; i++ {
			if position == i {
				fmt.Println(currentNode.data)
				break
			}

			currentNode = currentNode.next
		}
	}
}

func (sList *singlyLinkedList) reverseList() {
	if sList.head == nil {
		fmt.Println("Empty list")
	} else if sList.size == 1 {
		fmt.Println("Cannot reverse list of size 1")
	} else {
		for i := sList.size - 1; i >= 0; i-- {
			currentNode := sList.head
			defer sList.insertNodeAtTail(currentNode.data)
			sList.deleteNodeAtHead()
		}
	}
}

func (sList *singlyLinkedList) removeDuplicates() {
	if sList.head == nil {
		fmt.Println("Empty list")
	} else {
		var newSList singlyLinkedList

		nodeDataMap := make(map[int]int)
		currentNode := sList.head

		for currentNode != nil {
			_, ok := nodeDataMap[currentNode.data]
			if !ok {
				newSList.insertNodeAtTail(currentNode.data)
			}

			nodeDataMap[currentNode.data] = currentNode.data
			currentNode = currentNode.next
		}

		*sList = newSList
	}
}

func (sList singlyLinkedList) printSize() {
	fmt.Println(sList.size)
}

func (sList singlyLinkedList) printList() {
	currentNode := sList.head

	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.data)
		currentNode = currentNode.next
	}

	fmt.Println(currentNode)
}

func (sList singlyLinkedList) printListReverse() {
	currentNode := sList.head
	position := 0

	for currentNode != nil {
		if position == 0 {
			defer fmt.Println(currentNode.data)
		} else {
			defer fmt.Printf("%d <- ", currentNode.data)
		}

		currentNode = currentNode.next
		position++
	}

	if sList.size > 0 {
		fmt.Printf("%v <- ", currentNode)
	} else {
		fmt.Println(currentNode)
	}
}

func compareLists(sList1, sList2 singlyLinkedList) {
	isEqual := true

	if sList1.head != nil && sList2.head == nil || sList1.head == nil && sList2.head != nil {
		isEqual = false
	} else {
		sList1CurrentNode := sList1.head
		sList2CurrentNode := sList2.head

		for sList1CurrentNode != nil && sList2CurrentNode != nil {
			if sList1CurrentNode.data == sList2CurrentNode.data {
				sList1CurrentNode = sList1CurrentNode.next
				sList2CurrentNode = sList2CurrentNode.next
			} else {
				isEqual = false
				break
			}
		}

	}

	if isEqual {
		fmt.Println("Both lists are equal")
	} else {
		fmt.Println("Both lists are not equal")
	}
}

func mergeLists(sList1, sList2 singlyLinkedList) singlyLinkedList {
	var newSList singlyLinkedList

	sList1CurrentNode := sList1.head
	sList2CurrentNode := sList2.head

	for sList1CurrentNode != nil || sList2CurrentNode != nil {
		if sList1CurrentNode == nil && sList2CurrentNode != nil {
			newSList.insertNodeAtTail(sList2CurrentNode.data)
			sList2CurrentNode = sList2CurrentNode.next
		} else if sList1CurrentNode != nil && sList2CurrentNode == nil {
			newSList.insertNodeAtTail(sList1CurrentNode.data)
			sList1CurrentNode = sList1CurrentNode.next
		} else {
			if sList1CurrentNode.data <= sList2CurrentNode.data {
				newSList.insertNodeAtTail(sList1CurrentNode.data)
				newSList.insertNodeAtTail(sList2CurrentNode.data)
			} else {
				newSList.insertNodeAtTail(sList2CurrentNode.data)
				newSList.insertNodeAtTail(sList1CurrentNode.data)
			}

			sList1CurrentNode = sList1CurrentNode.next
			sList2CurrentNode = sList2CurrentNode.next
		}
	}

	return newSList
}

func main() {
	linkedList := singlyLinkedList{}

	/*
		Insert node at head
	*/

	// linkedList.insertNodeAtHead(4)
	// linkedList.insertNodeAtHead(3)
	// linkedList.insertNodeAtHead(2)
	// linkedList.insertNodeAtHead(1)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)

	/*
		Insert node at tail
	*/

	// linkedList.insertNodeAtTail(5)
	// linkedList.insertNodeAtTail(6)
	// linkedList.insertNodeAtTail(7)
	// linkedList.insertNodeAtTail(8)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)

	/*
		Insert node at given position
	*/

	// linkedList.insertNodeAtPosition(1, 0)
	// linkedList.insertNodeAtTail(4)
	// linkedList.insertNodeAtPosition(1, 0)
	// linkedList.insertNodeAtPosition(2, 1)
	// linkedList.insertNodeAtPosition(3, 2)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)

	/*
		Delete node at head
	*/

	// linkedList.deleteNodeAtHead()
	// linkedList.insertNodeAtTail(1)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.deleteNodeAtHead()
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.insertNodeAtTail(2)
	// linkedList.insertNodeAtTail(3)
	// linkedList.insertNodeAtTail(4)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.deleteNodeAtHead()
	// linkedList.deleteNodeAtHead()
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)

	/*
		Delete node at tail
	*/

	// linkedList.deleteNodeAtTail()
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.insertNodeAtTail(1)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.deleteNodeAtTail()
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.insertNodeAtTail(2)
	// linkedList.insertNodeAtTail(3)
	// linkedList.insertNodeAtTail(4)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.deleteNodeAtTail()
	// linkedList.deleteNodeAtTail()
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)

	/*
		Delete node at given position
	*/

	// linkedList.deleteNodeAtPosition(0)
	// linkedList.insertNodeAtTail(1)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.deleteNodeAtPosition(5)
	// linkedList.insertNodeAtTail(2)
	// linkedList.insertNodeAtTail(3)
	// linkedList.insertNodeAtTail(4)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.deleteNodeAtPosition(0)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.deleteNodeAtPosition(2)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)

	/*
		Print the linked list in reverse order
	*/

	// linkedList.printListReverse()
	// linkedList.insertNodeAtTail(1)
	// linkedList.printListReverse()
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.insertNodeAtTail(2)
	// linkedList.insertNodeAtTail(3)
	// linkedList.insertNodeAtTail(4)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.printListReverse()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)

	/*
		Reverse the linked list
	*/

	// linkedList.reverseList()
	// linkedList.insertNodeAtTail(1)
	// linkedList.reverseList()
	// linkedList.insertNodeAtTail(2)
	// linkedList.insertNodeAtTail(3)
	// linkedList.insertNodeAtTail(4)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.reverseList()
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)

	/*
		Compare equal lists
	*/

	// linkedList2 := singlyLinkedList{}
	// linkedList.insertNodeAtTail(1)
	// linkedList.insertNodeAtTail(2)
	// linkedList.insertNodeAtTail(3)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList2.insertNodeAtTail(1)
	// linkedList2.insertNodeAtTail(2)
	// linkedList2.insertNodeAtTail(3)
	// linkedList2.printList()
	// linkedList2.printSize()
	// fmt.Println(linkedList2.head)
	// fmt.Println(linkedList2.tail)
	// compareLists(linkedList, linkedList2)

	/*
		Compare unequal lists
	*/

	// linkedList3 := singlyLinkedList{}
	// linkedList.insertNodeAtHead(1)
	// linkedList.insertNodeAtHead(2)
	// linkedList.insertNodeAtHead(3)
	// linkedList.printList()
	// linkedList.printSize()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList3.insertNodeAtHead(1)
	// linkedList3.insertNodeAtHead(2)
	// linkedList3.printList()
	// linkedList3.printSize()
	// fmt.Println(linkedList3.head)
	// fmt.Println(linkedList3.tail)
	// compareLists(linkedList, linkedList3)

	/*
		Print node data at given position
	*/

	// linkedList.printNodeAtPosition(2)
	// linkedList.insertNodeAtTail(1)
	// linkedList.printNodeAtPosition(2)
	// linkedList.insertNodeAtTail(2)
	// linkedList.insertNodeAtTail(3)
	// linkedList.printNodeAtPosition(1)

	/*
		Merge two linked list in acceding order
	*/

	// linkedList2 := singlyLinkedList{}
	// mergedLinkedList1 := mergeLists(linkedList, linkedList2)
	// mergedLinkedList1.printList()
	// mergedLinkedList1.printSize()
	// fmt.Println(mergedLinkedList1.head)
	// fmt.Println(mergedLinkedList1.tail)
	// linkedList.insertNodeAtTail(1)
	// mergedLinkedList2 := mergeLists(linkedList, linkedList2)
	// mergedLinkedList2.printList()
	// fmt.Println(mergedLinkedList2.head)
	// fmt.Println(mergedLinkedList2.tail)
	// linkedList.insertNodeAtTail(3)
	// linkedList.insertNodeAtTail(5)
	// linkedList2.insertNodeAtTail(2)
	// linkedList2.insertNodeAtTail(4)
	// mergedLinkedList3 := mergeLists(linkedList, linkedList2)
	// mergedLinkedList3.printList()
	// fmt.Println(mergedLinkedList2.head)
	// fmt.Println(mergedLinkedList2.tail)

	/*
		Remove duplicate values from linked list
	*/

	// linkedList.removeDuplicates()
	// linkedList.insertNodeAtTail(1)
	// linkedList.insertNodeAtTail(2)
	// linkedList.insertNodeAtTail(3)
	// linkedList.insertNodeAtTail(3)
	// linkedList.insertNodeAtTail(2)
	// linkedList.insertNodeAtTail(1)
	// linkedList.printList()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
	// linkedList.removeDuplicates()
	// linkedList.printList()
	// fmt.Println(linkedList.head)
	// fmt.Println(linkedList.tail)
}
