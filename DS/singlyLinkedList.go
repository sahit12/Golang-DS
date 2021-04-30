package main

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

type SLL struct {
	head *Node
}

func (s *SLL) InsertAtFront(newnode *Node) {
	newnode.Next = s.head
	s.head = newnode
}

func (s *SLL) InsertAtEnd(newnode *Node) {
	current := s.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newnode
}

func (s *SLL) Search(value interface{}) {
	current := s.head
	count := 0
	for current != nil {
		if current.Data == value {
			fmt.Printf("Found value : [%v] at location : [%#v]", value, count)
			return
		}
		count++
		current = current.Next
	}
	fmt.Println("Value not found")
}

func (s SLL) printListData() {
	dispPoint := s.head
	for dispPoint != nil {
		fmt.Printf("Data: %d --> Next: %d", dispPoint.Data, dispPoint.Next)
		dispPoint = dispPoint.Next
		fmt.Println()
	}
	fmt.Printf("\n")
}

func main() {
	mylist := SLL{}
	node1 := Node{Data: 9}
	node2 := Node{Data: 8}
	node3 := Node{Data: 5}
	node4 := Node{Data: 3.56}
	node5 := Node{Data: 11.5}
	node6 := Node{Data: 2}
	node7 := Node{Data: 4}
	node8 := Node{Data: "name"}
	mylist.InsertAtFront(&node1)
	mylist.InsertAtFront(&node2)
	mylist.InsertAtFront(&node3)
	mylist.InsertAtFront(&node4)
	mylist.InsertAtFront(&node5)
	mylist.InsertAtFront(&node6)
	mylist.InsertAtEnd(&node7)
	mylist.InsertAtEnd(&node8)
	mylist.printListData()
	mylist.Search("name")
}
