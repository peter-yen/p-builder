package db

import (
	"fmt"
	"github.com/peter-yen/p-builder/pkg/global"
	"strings"
)

func (r *repo) mysqlDiver() (tables []Table) {
	// 查詢表格列表
	rows, err := r.DB.Query(mysqlTableStmt)
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
		tables = append(tables, Table{Name: strings.Title(tableName), Columns: r.getMysqlColumns(tableName)})
	}

	if err = rows.Err(); err != nil {
		global.Log.Println(err)
		return
	}

	return
}

func (r *repo) getMysqlColumns(tableName string) (arr []Column) {

	rows, err := r.DB.Query(fmt.Sprintf(mysqlCommentStmt, tableName))
	if err != nil {
		global.Log.Println(err)
		return
	}
	defer rows.Close()

	// mysql 不支援 rows.ColumnTypes() 的 scanType 操作

	for rows.Next() {
		var columnName, comment, dataType string
		if err = rows.Scan(&columnName, &comment, &dataType); err != nil {
			global.Log.Println(err)
			return
		}

		// TODO: ex:  member_id
		// TODO: 除了 首字母大寫外 把 橫線 去掉後的 第一個字母也要大寫
		arr = append(arr, Column{
			Name:     strings.Title(columnName),
			JsonName: columnName,
			DataType: dataType,
			GoType:   reflectMysqlType(dataType),
			Comment:  comment,
		})

		fmt.Printf("Column: %s, Comment: %s\n", columnName, comment)
		fmt.Printf("Column Name: %s, Data Type: %s, Go Type: %v\n", columnName, dataType, reflectMysqlType(dataType))
	}

	return
}

func reflectMysqlType(dataType string) string {
	// 統一規格 to lower
	switch strings.ToLower(dataType) {
	case "int", "mediumint", "bigint":
		return "int64"
	case "smallint", "tinyint":
		return "int8"
	case "float", "double", "decimal":
		return "float64"
	case "varchar", "char", "text", "longtext", "enum":
		return "string"
	case "date", "time", "datetime", "timestamp":
		return "time.Time"
	case "boolean":
		return "bool"
	case "blob":
		return "[]byte"
		// case "json": []byte
	}

	return "[]byte"
}
