package engine

import "errors"

// data base error handler
const (
	TABLE_EXIST     string = "[DATABASE] Table already exist!"
	TABLE_NOT_EXIST string = "[DATABASE] Table not exist yet!"
)

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
