package main

import (
	"log"
	"os"
	"p-builder/db"
	"p-builder/flags"
	"p-builder/global"
	"p-builder/tmp"
)

func main() {
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
