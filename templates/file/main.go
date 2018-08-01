package main

import (
	"fmt"
	"github.com/kardianos/osext"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Println("Hello main")

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

	//fmt.Println("Create file")
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
