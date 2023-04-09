package builder

import (
	"github.com/peter-yen/p-builder/pkg/db"
	"github.com/peter-yen/p-builder/pkg/flags"
	"github.com/peter-yen/p-builder/pkg/global"
	"github.com/peter-yen/p-builder/pkg/tmp"
	"log"
	"os"
)

// 專案起始點
func ListenFlags() {
	global.Log = log.New(os.Stdout, "[p-builder] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	// postgres
	// postgresql://peter:123456@localhost:5432/tmpl?sslmode=disable
	// model

	// parse flags
	driver, dir, folderName := flags.ParseFlags()

	// init db
	instance := db.InitDB(driver, dir)
	defer instance.Close()

	// get table list & column list
	list := db.GetTableList(instance)

	// generate model
	for _, v := range list {
		tmp.GenerateDB(v, folderName)
	}
}
