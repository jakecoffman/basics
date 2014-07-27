package hash_table

import "testing"

func Test_Hash_Table(t *testing.T) {
	hashtable := NewHashTable()
	hashtable.Put("key", "value")
	v, ok := hashtable.Get("key")
	if v != "value" || ok != true {
		t.Fatal("Failed to get 'key'")
	}

	v, ok = hashtable.Get("none")
	if ok != false {
		t.Fatal("Unexpectedly found key")
	}
}
