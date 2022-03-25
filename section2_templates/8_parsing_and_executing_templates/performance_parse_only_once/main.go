package main

import (
	"log"
	"os"
	"text/template"
)

/*
	We want to parse all templates only once in order to use them in this package
	We don't want to parse templates in every method in package

	We will store the parsed templates to package level variable
*/
var tpl *template.Template

/*
	Function init() run only once when program starts up
	We can use init() to parse templates only once
*/
func init() {
	/*
		ParseGlob returns (*Template, error)
		In order to handle the error and store only the *Template into "var tpl",
		we need to check the error.

		The Must functions does exactly this:

		Must is a helper that wraps a call to a function returning (*Template, error)
		and panics if the error is non-nil.

		So we call ParseGlob which returns (*Template, error)
		We pass (*Template, error) to Must that accepts this two params.
		Must checks the error and if ok, it returns the template *Template
		which we can store to "tpl" variable.
	*/
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
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
