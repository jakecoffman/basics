package linkedlist

import "testing"

func Test_LinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	if list == nil {
		t.Fatal("list was nil")
	}

	list.Append("hello")

	v, e := list.Get(0)
	if v != "hello" || e != nil {
		t.Fatal("list get or append not working:", v, e)
	}
}
