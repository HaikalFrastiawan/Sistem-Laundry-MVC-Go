package main

import (
	"Sistem-Laundry/config"
	"Sistem-Laundry/controllers/homecontroller"
	"Sistem-Laundry/controllers/servicecontroller"
	"Sistem-Laundry/controllers/clientcontroller"
	"Sistem-Laundry/controllers/admincontroller"
	"log"
	"net/http"
)

func main() {

	//Connect DB
	config.ConnectDb()

/// ---------------------------------------------------------
    // 1. JALUR PUBLIK (Landing Page)
    // ---------------------------------------------------------
    http.HandleFunc("/", homecontroller.Welcome)

    // ---------------------------------------------------------
    // 2. JALUR PELANGGAN (Client Area)
    // Semua yang diawali /client/ adalah milik user/Haikal
    // ---------------------------------------------------------
    http.HandleFunc("/client/dashboard", clientcontroller.Dashboard)
    http.HandleFunc("/client/order/create", clientcontroller.CreateOrder)

    // ---------------------------------------------------------
    // 3. JALUR ADMIN (Management Area)
    // Pastikan semua modul admin pakai awalan /admin/ agar rapi
    // ---------------------------------------------------------
    http.HandleFunc("/admin/dashboard", admincontroller.Dashboard)
    http.HandleFunc("/admin/order", admincontroller.IndexOrder)
    http.HandleFunc("/admin/order/update", admincontroller.UpdateStatus)
    http.HandleFunc("/admin/customer", admincontroller.IndexCustomer)
    http.HandleFunc("/admin/master", admincontroller.IndexMaster)

    // Modul Service (Kita masukkan ke jalur admin juga)
    http.HandleFunc("/admin/services", servicecontroller.Index)
    http.HandleFunc("/admin/services/add", servicecontroller.Add)
    http.HandleFunc("/admin/services/edit", servicecontroller.Edit)
    http.HandleFunc("/admin/services/delete", servicecontroller.Delete)

    log.Println("Server Running on port 8080")
    http.ListenAndServe(":8080", nil)
}
