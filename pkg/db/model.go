package db

import "database/sql"

type repo struct {
	DB         *sql.DB
	DriverName string
}

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name     string // 首字母大寫
	JsonName string // 小寫
	GoType   string // go reflect type
	DataType string // db type
	Comment  string // 欄位備註
}
