package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

type node struct {
	Data string
	Next *node
}

// SinglyLinkedList is the reference implementation of a singly linked list
type SinglyLinkedList struct {
	head *node
}

// NewSinglyLinkedList constructor
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Append adds a string to the end of the list
func (l *SinglyLinkedList) Append(value string) {
	cur := &l.head
	for {
		if *cur == nil {
			*cur = &node{Data: value}
			return
		}
		cur = &(*cur).Next
	}
}

// Get retrieves a string at the given index
func (l SinglyLinkedList) Get(index int) (string, error) {
	return l.find(index, func(cur **node, _ **node) (string, error) {
		return (*cur).Data, nil
	})
}

// Pop retrieves a string at the given index and removes it from the list
func (l *SinglyLinkedList) Pop(index int) (string, error) {
	return l.find(index, func(cur **node, prev **node) (string, error) {
		temp := (*cur).Data
		if prev != nil {
			(*prev).Next = (*cur).Next
		} else {
			l.head = (*cur).Next
		}
		return temp, nil
	})

}

func (l SinglyLinkedList) find(index int, f func(cur **node, prev **node) (string, error)) (string, error) {
	i := 0
	cur := &l.head
	var prev **node

	for {
		if *cur == nil {
			break
		}
		if i == index {
			return f(cur, prev)
		}
		cur = &(*cur).Next
		i++
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
