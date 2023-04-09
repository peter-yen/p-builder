package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // 引入 "pq" 包
	"github.com/peter-yen/p-builder/pkg/global"
	"strings"
)

func InitDB(driverName, dir string) (db *sql.DB) {
	db, err := sql.Open(driverName, dir)
	if err != nil {
		global.Log.Println(err)
		return
	}

	if err = db.Ping(); err != nil {
		global.Log.Println("Failed to ping PostgreSQL: ", err)
		return
	}

	global.Log.Println("--- Successfully connected to PostgreSQL! ---")

	return db
}

// GetTableList 獲取表格列表
func GetTableList(db *sql.DB) (tables []Table) {

	// 查詢表格列表
	rows, err := db.Query(getTableList)
	if err != nil {
		global.Log.Println(err)
		return
	}
	defer rows.Close()

	// 遍歷結果集，獲取表格名稱
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			global.Log.Println(err)
			return
		}
		fmt.Println("Table name:", tableName)

		// MARK: 遍歷表格欄位
		tables = append(tables, Table{Name: strings.Title(tableName), Columns: iterateColumns(db, tableName)})
	}

	if err = rows.Err(); err != nil {
		global.Log.Println(err)
		return
	}

	return
}

// iterateColumns 遍歷 table 欄位 獲取 reflect type, name, comment
func iterateColumns(db *sql.DB, table string) (arr []Column) {
	rows, err := db.Query("SELECT * FROM " + table + " LIMIT 1")
	if err != nil {
		global.Log.Println(err)
		return
	}
	defer rows.Close()

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		global.Log.Println(err)
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		global.Log.Println(err)
		return
	}

	for i, col := range columns {
		var comment string

		if err = db.QueryRow(fmt.Sprintf(getCommentStmt, table, col)).
			Scan(&comment); err != nil {
			global.Log.Println(err)
			return
		}

		arr = append(arr, Column{
			Name:     strings.Title(columnTypes[i].Name()),
			JsonName: columnTypes[i].Name(),
			DataType: columnTypes[i].DatabaseTypeName(),
			GoType:   columnTypes[i].ScanType().String(),
			Comment:  comment,
		})

		fmt.Printf("Column: %s, Comment: %s\n", col, comment)
		fmt.Printf("Column Name: %s, Data Type: %s, Go Type: %v\n", columnTypes[i].Name(), columnTypes[i].DatabaseTypeName(), columnTypes[i].ScanType())
	}
	return
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

/*
// iterateColumns 遍歷表格欄位
func iterateColumns(db *sql.DB, table string) (arr []Column) {
	// 查詢表格欄位列表
	rows, err := db.Query("SELECT column_name, column_default, is_nullable, data_type FROM information_schema.columns WHERE table_name = $1", table)
	if err != nil {
		global.Log.Println(err)
		return
	}
	defer rows.Close()

	// 遍歷結果集，獲取欄位名稱
	for rows.Next() {
		var columnName string
		var columnDefault sql.NullString
		var isNullable string
		var dataType string

		err = rows.Scan(&columnName, &columnDefault, &isNullable, &dataType)
		if err != nil {
			global.Log.Println(err)
			return
		}

		arr = append(arr, Column{Name: strings.Title(columnName), DataType: dataType})

		fmt.Println("Column name:", columnName)
		fmt.Println("Column default value:", columnDefault.String)
		fmt.Println("Is nullable:", isNullable)
		fmt.Println("Data type:", dataType)
	}

	if err = rows.Err(); err != nil {
		global.Log.Println(err)
		return
	}

	// 獲取 reflect type
	getGoTypes(db, table, arr)

	return
}

*/
