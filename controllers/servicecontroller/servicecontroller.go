package servicecontroller

import (
	"net/http"
	"Sistem-Laundry/models/servicemodel"
	"html/template"
	"Sistem-Laundry/entities"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
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

// 1. HALAMAN TAMBAH (GET)
func Add(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        temp, err := template.ParseFiles("views/service/add.html")
        if err != nil {
            panic(err)
        }
        temp.Execute(w, nil)
    }
}

// 2. PROSES SIMPAN (POST)
func Store(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var service entities.Service

        service.Name = r.FormValue("name")
        priceStr := r.FormValue("price") // Pastikan name di input HTML adalah "price"
        service.PricePerKg, _ = strconv.Atoi(priceStr)

        if ok := servicemodel.Create(service); ok {
            http.Redirect(w, r, "/services", http.StatusSeeOther)
        } else {
            http.Error(w, "Gagal menambah data", http.StatusInternalServerError)
        }
    }
}

// 3. HALAMAN EDIT (GET)
func Edit(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        // Ambil ID dari query string URL (?id=1)
        idString := r.URL.Query().Get("id")
        id, _ := strconv.Atoi(idString)

        // Ambil data service berdasarkan ID
        service := servicemodel.Detail(id)

        data := map[string]any{
            "service": service,
        }

        temp, err := template.ParseFiles("views/service/edit.html")
        if err != nil {
            panic(err)
        }
        temp.Execute(w, data)
    }
}

// 4. PROSES UPDATE (POST)
func Update(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var service entities.Service

        idString := r.FormValue("id")
        id, _ := strconv.Atoi(idString)

        service.Name = r.FormValue("name")
        priceStr := r.FormValue("price")
        service.PricePerKg, _ = strconv.Atoi(priceStr)

        if ok := servicemodel.Update(id, service); ok {
            http.Redirect(w, r, "/services", http.StatusSeeOther)
        } else {
            http.Error(w, "Gagal update data", http.StatusInternalServerError)
        }
    }
}

// 5. PROSES DELETE (GET/POST)
func Delete(w http.ResponseWriter, r *http.Request) {
    idString := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idString)

    if ok := servicemodel.Delete(id); ok {
        http.Redirect(w, r, "/services", http.StatusSeeOther)
    } else {
        http.Error(w, "Gagal menghapus data", http.StatusInternalServerError)
    }
}
