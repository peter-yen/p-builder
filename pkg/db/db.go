package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 引入 MySQL 驅動套件
	_ "github.com/lib/pq"              // 引入 "pq" 包
	"github.com/peter-yen/p-builder/pkg/global"
)

type repo struct {
	DB         *sql.DB
	DriverName string
}

func NewInstance(driverName, dir string) (entity repo) {

	db, err := sql.Open(driverName, dir)
	if err != nil {
		global.Log.Println(err)
		return
	}

	if err = db.Ping(); err != nil {
		global.Log.Println("Failed to ping sql server: ", err)
		return
	}

	global.Log.Println("--- Successfully connected to Server! ---")

	entity = repo{
		DB:         db,
		DriverName: driverName,
	}

	return
}

// GetTableList 獲取表格列表
func (r *repo) GetTableList() []Table {

	switch r.DriverName {
	case "postgres":
		return r.postgresDiver()
	case "mysql":
		return r.mysqlDiver()
	}

	return nil
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
