package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 引入 MySQL 驅動套件
	_ "github.com/lib/pq"              // 引入 "pq" 包
	"github.com/peter-yen/p-builder/pkg/global"
)

// NewInstance 初始化資料庫連線
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

// GetTableList 取得資料庫中所有的表格
func (r *repo) GetTableList() []Table {

	switch r.DriverName {
	case "postgres":
		return r.postgresDiver()
	case "mysql":
		return r.mysqlDiver()
	}

	return nil
}
