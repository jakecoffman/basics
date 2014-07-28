package linkedlist

import (
	"errors"
	"fmt"
	"strings"
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
	cur := &l.head
	for {
		if *cur == nil {
			*cur = &Node{Data: value}
			return
		}
		cur = &(*cur).Next
	}
}

func (l SinglyLinkedList) Get(index int) (string, error) {
	return l.find(index, func(cur **Node, _ **Node) (string, error) {
		return (*cur).Data, nil
	})
}

func (l *SinglyLinkedList) Pop(index int) (string, error) {
	return l.find(index, func(cur **Node, prev **Node) (string, error) {
		temp := (*cur).Data
		if prev != nil {
			(*prev).Next = (*cur).Next
		} else {
			l.head = (*cur).Next
		}
		return temp, nil
	})

}

func (l SinglyLinkedList) find(index int, f func(cur **Node, prev **Node) (string, error)) (string, error) {
	i := 0
	cur := &l.head
	var prev **Node

	for {
		if *cur == nil {
			break
		}
		if i == index {
			return f(cur, prev)
		}
		cur = &(*cur).Next
		i += 1
	}
	return "", errors.New(fmt.Sprint("Failed to find item at index ", index))
}

func (l SinglyLinkedList) String() string {
	cur := l.head
	str := []string{}

	for {
		if cur == nil {
			break
		}
		str = append(str, cur.Data)
		cur = cur.Next
	}
	return "[" + strings.Join(str, ", ") + "]"
}
