package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func render(w http.ResponseWriter, temp string, args interface{}) {
	fmt.Println("readDir")
	fmt.Println(os.ReadDir("static/"))
	parsed, err := template.ParseFiles("static/" + temp + ".html")
	if err != nil {
		panic(err)
	}
	if err := parsed.Execute(w, args); err != nil {
		panic(err)
	}
}
