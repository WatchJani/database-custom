package engine

import (
	"errors"
)

// data base error handler


type DataBase struct {
	name  string
	table map[string]*Table
}

func (db DataBase) SelectTable(name string) (*Table, error) {
	if table, ok := db.table[name]; ok {
		return table, nil
	}

	return nil, errors.New(TABLE_NOT_EXIST)
}

func NewDataBase(name string) *DataBase {
	return &DataBase{
		name:  name,
		table: make(map[string]*Table),
	}
}

func (db *DataBase) AddTable(name string, table *Table) {
	db.table[name] = table
}

func (db *DataBase) AddIndex(table, index string) error {
	user, err := db.SelectTable(table)
	if err != nil {
		return err
	}

	user.AddIndex(index)
	return nil
}
