package hash_table

import "testing"

func Test_Hash_Table(t *testing.T) {
	hashtable := NewHashTable()

	if hashtable.String() != "{}" {
		t.Fatal("Unexpected String output:", hashtable.String())
	}

	hashtable.Put("key", "value")

	check(*hashtable, "key", "value", true, t)
	check(*hashtable, "none", "", false, t)

	hashtable.Put("a", "b")
	check(*hashtable, "a", "b", true, t)
	check(*hashtable, "key", "value", true, t)
	check(*hashtable, "none", "", false, t)

	if hashtable.String() != "{key:value, a:b}" {
		t.Fatal("Unexpected String output:", hashtable.String())
	}
}

func check(table HashTable, key, expected string, ok bool, t *testing.T) {
	v, k := table.Get(key)
	if v != expected || k != ok {
		t.Fatal("Unexpected:", v, ok)
	}
}
