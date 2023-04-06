package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {

	tmp, err := template.New("test").Parse("{{define \"T\"}}Hello, {{.}}!{{end}}")
	if err != nil {
		fmt.Print(err)
		return
	}

	if err = tmp.Execute(os.Stdout, "World"); err != nil {
		fmt.Print(err)
		return
	}
}
