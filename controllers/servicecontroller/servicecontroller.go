package servicecontroller

import (
	"net/http"
	"Sistem-Laundry/models/servicemodel"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	//ambil data dari model
	services := servicemodel.GetAll()

	//tampilkan data di view
	data := map[string]any{
		"services": services,
	}

	temp,err := template.ParseFiles("views/service/index.html")
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	
	temp.Execute(w, data)
}