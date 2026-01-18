package homecontroller

import (
    "html/template"
    "net/http"
)

// Menampilkan Halaman Depan (Landing Page)
func Welcome(w http.ResponseWriter, r *http.Request) {
    temp, err := template.ParseFiles("views/home/landing.html")
    if err != nil {
        panic(err)
    }
    temp.Execute(w, nil)
}
