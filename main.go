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

	instance := db.InitDB()
	defer instance.Close()

	list := db.GetTableList(instance)

	for _, v := range list {
		tmp.GenerateDB(v)
	}

}
