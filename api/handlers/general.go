package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	temp, err := template.ParseFiles("./api/templates/views/dashboard.html")
	if err != nil {
		panic(err)
	}
	err = temp.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
