package main

import (
	"log"
	"os"
	"p-builder/db"
	"p-builder/global"
	"p-builder/tmp"
)

func main() {
	global.Log = log.New(os.Stdout, "[p-builder] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	// TODO: 使用 flags 來設定資料庫連線資訊 和 生成的 資料夾路徑、名稱

	// init db
	instance := db.InitDB("postgres", "postgresql://peter:123456@localhost:5432/tmpl?sslmode=disable")
	defer instance.Close()

	// get table list & column list
	list := db.GetTableList(instance)

	// generate model
	for _, v := range list {
		tmp.GenerateDB(v)
	}

}
