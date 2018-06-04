package main

import (
	"fmt"
	"github.com/kardianos/osext"
	"log"
	"os"
	"text/template"
)

func main() {

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(folderPath)

	tpl, err := template.ParseFiles(folderPath + 
		string(os.PathSeparator) + "templates" + string(os.PathSeparator) + "httpd.conf.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
Do not use the above code in production
 We will learn about efficiency improvements soon!
*/
