package engine

import (
	"errors"
	"root/index"
	btree "root/index/b_tree"
	"root/index/hash_map"
	"root/index/trie"
)

type Table struct {
	name     string
	index    []index.CustomIndexer
	allIndex map[string]index.CustomIndexer
}

func newTable(name string) *Table {
	return &Table{
		name:  name,
		index: []index.CustomIndexer{hash_map.NewHashMap()}, //default index
		allIndex: map[string]index.CustomIndexer{
			"hash_index": hash_map.NewHashMap(),
			"trie":       trie.NewTrie(),
			"b_tree":     btree.NewBTree(),
		},
	}
}

func (db *DataBase) CreateTable(name string) (*Table, error) {
	if _, ok := db.table[name]; ok {
		return nil, errors.New(TABLE_EXIST)
	}

	//CREATE NEW FILE FOR OUR TABLE

	table := newTable(name)
	db.AddTable(name, table)

	return table, nil
}

func (t *Table) AddIndex(name string) {
	t.index = append(t.index, t.allIndex[name])
}
