package clientcontroller

import (
    "Sistem-Laundry/entities"
    "Sistem-Laundry/models/ordermodel"
    "Sistem-Laundry/models/servicemodel"
    "html/template"
    "net/http"
    "strconv"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
    // 1. Ambil Data Service untuk Dropdown Modal
    services := servicemodel.GetAll()

    // 2. Simulasi Session Login (Hardcode dulu)
    customerName := "Haikal"

    // 3. Tracking: Ambil order status Pending/Cuci/Setrika
    activeOrders := ordermodel.GetByStatusAndCustomer(customerName, "Proses")

    // 4. History: Ambil order status Selesai
    historyOrders := ordermodel.GetByStatusAndCustomer(customerName, "Selesai")

    data := map[string]any{
        "customerName": customerName,
        "services":     services,
        "activeOrders": activeOrders,
        "history":      historyOrders,
    }

    temp, err := template.ParseFiles("views/client/dashboard.html")
    if err != nil {
        panic(err)
    }
    temp.Execute(w, data)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var order entities.Order

        // Ambil data dari Form Modal
        order.CustomerName = "Haikal" // Sesuaikan dengan session nanti
        
        serviceId, _ := strconv.Atoi(r.FormValue("service_id"))
        order.Service.ID = serviceId
        
        weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
        order.Weight = weight
        
        // Panggil Model untuk Simpan ke Database
        if ok := ordermodel.Create(order); ok {
            http.Redirect(w, r, "/client/dashboard", http.StatusSeeOther)
        } else {
            http.Error(w, "Gagal membuat pesanan", http.StatusInternalServerError)
        }
    }
}