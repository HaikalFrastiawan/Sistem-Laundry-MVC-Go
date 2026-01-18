package admincontroller

import (
    "Sistem-Laundry/models/ordermodel"
    "Sistem-Laundry/models/customermodel"
    "Sistem-Laundry/models/servicemodel"
    "Sistem-Laundry/models/employeemodel"
    "html/template"
    "net/http"
    "strconv"
)

// Fungsi helper untuk memanggil template agar tidak menulis ulang ParseFiles
func render(w http.ResponseWriter, contentPath string, data map[string]any) {
    tmpl, err := template.ParseFiles("views/admin/layout.html", contentPath)
    if err != nil {
        http.Error(w, "Template Error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // "admin-layout" adalah nama yang didefinisikan di layout.html
    tmpl.ExecuteTemplate(w, "admin-layout", data)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
    data := map[string]any{
        "Title":  "Dashboard Ringkasan",
        "Active": "dash",
        "Stats": map[string]any{
            "TotalOrders":    120,
            "TotalRevenue":   "4.5M",
            "TotalCustomers": 50,
        },
    }
    render(w, "views/admin/dashboard.html", data)
}

func IndexOrder(w http.ResponseWriter, r *http.Request) {
    orders := ordermodel.GetAll()
    data := map[string]any{
        "Title":  "Kelola Pesanan",
        "Active": "order",
        "orders": orders,
    }
    render(w, "views/admin/order/index.html", data)
}

func IndexCustomer(w http.ResponseWriter, r *http.Request) {
    // 1. Ambil data dari model
    customers := customermodel.GetAll() 
    
    data := map[string]any{
        "Title":     "Data Pelanggan",
        "Active":    "cust", // Agar menu di sidebar otomatis biru (active)
        "customers": customers,
    }

    // 2. WAJIB sertakan layout.html
    tmpl := template.Must(template.ParseFiles("views/admin/layout.html", "views/admin/customer/index.html"))
    
    // 3. Eksekusi menggunakan nama define di layout
    tmpl.ExecuteTemplate(w, "admin-layout", data)
}

func IndexMaster(w http.ResponseWriter, r *http.Request) {
    services := servicemodel.GetAll()
    employees := employeemodel.GetAll()

    data := map[string]any{
        "Title":     "Data Master & Layanan",
        "Active":    "master",
        "services":  services,
        "employees": employees,
    }

    // WAJIB sertakan layout.html
    tmpl := template.Must(template.ParseFiles("views/admin/layout.html", "views/admin/master/index.html"))
    tmpl.ExecuteTemplate(w, "admin-layout", data)
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        id, _ := strconv.Atoi(r.FormValue("id"))
        status := r.FormValue("status")
        if ok := ordermodel.UpdateStatus(id, status); ok {
            http.Redirect(w, r, "/admin/order", http.StatusSeeOther)
        } else {
            http.Error(w, "Gagal update status", http.StatusInternalServerError)
        }
    }
}
func AdminDashboard(w http.ResponseWriter, r *http.Request) {
    // Ambil data dari model
    tampilOrders := ordermodel.GetAll()

    data := map[string]any{
        "orders": tampilOrders, // Pastikan namanya "orders"
    }

    tmpl, _ := template.ParseFiles("views/admin/index.html")
    tmpl.Execute(w, data)
}