package db

import (
	"fmt"
	"github.com/peter-yen/p-builder/pkg/global"
	"strings"
)

func (r *repo) postgresDiver() (tables []Table) {
	// 查詢表格列表
	rows, err := r.DB.Query(postgresTableStmt)
	if err != nil {
		global.Log.Println(err)
		return
	}
	defer rows.Close()

	var tableName string
	// 遍歷結果集，獲取表格名稱
	for rows.Next() {
		err = rows.Scan(&tableName)
		if err != nil {
			global.Log.Println(err)
			return
		}
		fmt.Println("Table name:", tableName)

		// MARK: 遍歷表格欄位
		tables = append(tables, Table{Name: strings.Title(tableName), Columns: r.getPostgresColumns(tableName)})
	}

	if err = rows.Err(); err != nil {
		global.Log.Println(err)
		return
	}

	return
}

// iterateColumns 遍歷 table 欄位 獲取 reflect type, name, comment
func (r *repo) getPostgresColumns(table string) (arr []Column) {

	rows, err := r.DB.Query("SELECT * FROM " + table + " LIMIT 1")
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

		if err = r.DB.QueryRow(fmt.Sprintf(postgresCommentStmt, table, col)).
			Scan(&comment); err != nil {
			global.Log.Println(err)
			return
		}

		// TODO: ex:  member_id
		// TODO: 除了 首字母大寫外 把 橫線 去掉後的 第一個字母也要大寫
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
