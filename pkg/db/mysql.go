package db

import (
	"fmt"
	"github.com/peter-yen/p-builder/pkg/global"
	"strings"
)

// mysqlDiver mysql 查詢表格列表
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

		name := bottomLineToUpper(tableName)

		// 遍歷表格欄位
		tables = append(tables, Table{Name: strings.Title(tableName), StructName: strings.Title(name), Columns: r.getMysqlColumns(tableName)})
	}

	if err = rows.Err(); err != nil {
		global.Log.Println(err)
		return
	}

	return
}

// getMysqlColumns iterate columns 遍歷 table 欄位 獲取 reflect type, name, comment
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

		// 讓 _ 之後的字母大寫
		name := bottomLineToUpper(columnName)

		arr = append(arr, Column{
			Name:     strings.Title(name),
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

// reflectMysqlType 獲取 mysql 欄位對應的 golang type
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
