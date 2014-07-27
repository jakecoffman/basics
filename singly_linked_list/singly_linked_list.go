package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Data string
	Next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (l *SinglyLinkedList) Append(value string) {
	if l.head == nil {
		l.head = &Node{Data: value}
		return
	}
	cur := l.head
	for {
		if cur.Next == nil {
			cur.Next = &Node{Data: value}
			return
		}
		cur = cur.Next
	}
}

func (l SinglyLinkedList) Get(index int) (string, error) {
	i := 0
	cur := l.head

	for {
		if cur == nil {
			break
		}
		if i == index {
			return cur.Data, nil
		}
		cur = cur.Next
	}
	return "", errors.New(fmt.Sprint("Failed to find item at index ", index))
}
