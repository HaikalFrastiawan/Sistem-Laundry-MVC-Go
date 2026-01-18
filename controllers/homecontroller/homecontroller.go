package homecontroller

import (
	"net/http"
	"html/template"
)


func Index(w http.ResponseWriter, r *http.Request) {	
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}