package linkedlist

import "testing"

func Test_LinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	if list == nil {
		t.Fatal("list was nil")
	}

	if list.String() != "[]" {
		t.Fatal("Unexpected String output:", list.String())
	}

	list.Append("hello")
	if list.String() != "[hello]" {
		t.Fatal("Expected [hello]")
	}

	check(list, 0, "hello", t)

	list.Append("world")
	if list.String() != "[hello, world]" {
		t.Fatal("Expected [hello, world]")
	}

	check(list, 0, "hello", t)
	check(list, 1, "world", t)

	v, e := list.Pop(0)
	if v != "hello" || e != nil {
		t.Fatal("Pop failed:", v, "-", e)
	}

	check(list, 0, "world", t)
}

func check(list *SinglyLinkedList, index int, value string, t *testing.T) {
	v, e := list.Get(index)
	if v != value || e != nil {
		t.Fatal("list get or append not working:", v, "-", e)
	}
}
