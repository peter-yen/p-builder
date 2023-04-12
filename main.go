package main

import (
	"github.com/peter-yen/p-builder/pkg/db"
	"github.com/peter-yen/p-builder/pkg/flags"
	"github.com/peter-yen/p-builder/pkg/global"
	"github.com/peter-yen/p-builder/pkg/tmp"
	"log"
	"os"
)

// driver: driver name
// dir: database connection dir
// folderName: folder name
func main() {
	// TODO: 穩定過後 把 Lshortfile 去掉， error 應該 print 出錯誤就好
	global.Log = log.New(os.Stdout, "[p-builder] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	// postgres.go
	// postgresql://peter:123456@localhost:5432/tmpl?sslmode=disable
	// model

	// parse flags
	driver, dir, folderName := flags.ParseFlags()

	// init db
	instance := db.NewInstance(driver, dir)
	//instance := db.InitDB(driver, dir)
	defer instance.DB.Close()

	// get table list & column list
	list := instance.GetTableList()

	// generate model
	for _, v := range list {
		tmp.GenerateDB(v, folderName)
	}
}
