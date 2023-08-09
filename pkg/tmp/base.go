package tmp

import (
	"github.com/peter-yen/p-builder/pkg/db"
	"github.com/peter-yen/p-builder/pkg/global"
	"os"
	"strings"
	"text/template"
)

// GenerateDB 產生 model
func GenerateDB(field db.Table, folderName string) {

	tmpl, err := template.New(field.Name).Parse(str)
	if err != nil {
		global.Log.Println(err)
		return
	}

	file := createFileAndDir(folderName, field.Name+".go")
	defer file.Close()

	// os.Stdout
	if err = tmpl.Execute(file, field); err != nil {
		global.Log.Println(err)
	}

}

// createFileAndDir 建立檔案和資料夾
func createFileAndDir(dir, fileName string) *os.File {
	if err := os.MkdirAll(dir, 0777); err != nil {
		global.Log.Println(err)
		return nil
	}

	file, err := os.OpenFile(dir+"/"+prefixToLower(fileName), os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		global.Log.Println(err)
		return nil
	}

	return file
}

func prefixToLower(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}

// MARK: . 很重要！！， column 前面要加點 .
// range .Columns , end  可以使用 []struct 作為傳入參數

const str = `
package model

type {{.Name}} struct {
{{range .Columns}}
	// {{.Comment}}
	{{.Name}} {{.GoType}} ` + "`json:\"{{.JsonName}}\"`" + `
{{end}}
}
`
