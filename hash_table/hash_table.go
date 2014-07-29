package hashTable

import (
	"fmt"
	"hash/adler32"
	"strings"
)

var digest = adler32.New()

const maxTableSize = 128

type node struct {
	Key   string
	Value string
}

// HashTable is a reference implementation of a hash table
type HashTable struct {
	Table [maxTableSize]*node
}

// NewHashTable creates and returns a new hash table
func NewHashTable() *HashTable {
	return &HashTable{}
}

// Get returns a value by string, and true/false if the key exists
func (t HashTable) Get(key string) (string, bool) {
	hash := adler32.Checksum([]byte(key)) % maxTableSize
	for {
		if t.Table[hash] != nil && t.Table[hash].Key != key {
			hash = (hash + 1) % maxTableSize
		} else {
			break
		}
	}
	if t.Table[hash] == nil {
		return "", false
	}
	return t.Table[hash].Value, true
}

// Put sets a value by key in the hash map
func (t *HashTable) Put(key, value string) {
	hash := adler32.Checksum([]byte(key)) % maxTableSize
	for {
		if t.Table[hash] != nil && t.Table[hash].Key != key {
			hash = (hash + 1) % maxTableSize
		} else {
			break
		}
	}
	t.Table[hash] = &node{key, value}
}

func (t HashTable) String() string {
	str := []string{}
	for _, v := range t.Table {
		if v != nil {
			str = append(str, fmt.Sprint(v.Key, ":", v.Value))
		}
	}
	return "{" + strings.Join(str, ", ") + "}"
}
