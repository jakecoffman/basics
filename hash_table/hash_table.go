package hash_table

import (
	"hash"
	"hash/adler32"
)

var digest hash.Hash32 = adler32.New()

const MAX_TABLE_SIZE = 128

type Node struct {
	Key   string
	Value string
}

type HashTable struct {
	Table [MAX_TABLE_SIZE]*Node
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func (t HashTable) Get(key string) (string, bool) {
	hash := adler32.Checksum([]byte(key)) % MAX_TABLE_SIZE
	for {
		if t.Table[hash] != nil && t.Table[hash].Key != key {
			hash = (hash + 1) % MAX_TABLE_SIZE
		} else {
			break
		}
	}
	if t.Table[hash] == nil {
		return "", false
	} else {
		return t.Table[hash].Value, true
	}
}

func (t *HashTable) Put(key, value string) {
	hash := adler32.Checksum([]byte(key)) % MAX_TABLE_SIZE
	for {
		if t.Table[hash] != nil && t.Table[hash].Key != key {
			hash = (hash + 1) % MAX_TABLE_SIZE
		} else {
			break
		}
	}
	t.Table[hash] = &Node{key, value}
}
