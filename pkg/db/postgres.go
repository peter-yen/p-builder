package db

import (
	"database/sql"
	"fmt"
	"github.com/peter-yen/p-builder/pkg/global"
	"strings"
)

func getPostgresColumn(table string, columns []string, columnTypes []*sql.ColumnType) (arr []Column, err error) {
	for i, col := range columns {
		var comment string

		if err = r.DB.QueryRow(fmt.Sprintf(postgresCommentStmt, table, col)).
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
}
