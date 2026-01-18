package authcontroller

import (
	"Sistem-Laundry/entities"
	"Sistem-Laundry/models/customermodel"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("views/auth/login.html")
		temp.Execute(w, nil)
		return
	}

	// Proses Login POST
	phone := r.FormValue("phone")
	password := r.FormValue("password")

	customer, ok := customermodel.Login(phone, password)
	if !ok {
		// Jika gagal, balikkan ke login dengan pesan error
		data := map[string]any{"Error": "Nomor WA atau Password salah!"}
		temp, _ := template.ParseFiles("views/auth/login.html")
		temp.Execute(w, data)
		return
	}

	// Jika sukses, buat session
	session, _ := store.Get(r, "customer-session")
	session.Values["customer_id"] = customer.ID
	session.Values["customer_name"] = customer.Name
	session.Save(r, w)

	http.Redirect(w, r, "/client/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "customer-session")
	session.Options.MaxAge = -1 // Hapus session
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "customer-session")
		if auth, ok := session.Values["customer_id"].(int); !ok || auth == 0 {
			http.Redirect(w, r, "/client/dashboard", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// Tambahkan "Sistem-Laundry/entities" di bagian import atas

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("views/auth/register.html")
		temp.Execute(w, nil)
		return
	}

	// Tangkap data dari form register
	newCustomer := entities.Customer{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Phone:    r.FormValue("phone"),
		Address:  r.FormValue("address"),
		Password: r.FormValue("password"),
	}

	// Panggil model Create yang sudah kita buat tadi
	if customermodel.Create(newCustomer) {
		// Jika berhasil, arahkan ke login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		// Jika gagal (misal nomor HP sudah terdaftar)
		data := map[string]any{"Error": "Gagal mendaftar. Nomor HP mungkin sudah digunakan."}
		temp, _ := template.ParseFiles("views/auth/register.html")
		temp.Execute(w, data)
	}
}
