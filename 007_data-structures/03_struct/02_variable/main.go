package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	rust := sage{
		Name:  "Rust",
		Motto: "A language empowering everyone to build reliable and efficient software",
	}

	err := tpl.Execute(os.Stdout, rust)
	if err != nil {
		log.Fatalln(err)
	}
}
