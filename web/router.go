package web

import (
	"fmt"
	"net/http"
	"os"
)

const (
	Index = "index"
)

func SetRouting(argsMap map[string]interface{}) {
	fmt.Println("argsMap", argsMap)
	http.HandleFunc("/", indexHandler(argsMap[Index]))

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd+"/static/"))))
}

func makeHandler(templateName string, args interface{}) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		render(writer, templateName, args)
	}
}

func indexHandler(arg interface{}) func(http.ResponseWriter, *http.Request) {
	return makeHandler(Index, arg)
}
