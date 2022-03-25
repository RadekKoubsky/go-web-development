package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// This will parse *.gohtml files by default.
	// We can use any file suffix for go templates, for instance:
	// Given a file my_template.my_suffix in templates dir
	// The ParseGlob func will use that suffix in the pattern:
	// template.ParseGlob("templates/*.my_suffix")
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
