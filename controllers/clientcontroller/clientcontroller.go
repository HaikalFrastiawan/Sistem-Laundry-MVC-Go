package clientcontroller

import (
    "Sistem-Laundry/entities"
    "Sistem-Laundry/models/ordermodel"
    "Sistem-Laundry/models/servicemodel"
    "github.com/gorilla/sessions"
    "html/template"
    "net/http"
    "strconv"
    "time"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func Dashboard(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "customer-session")
    name, _ := session.Values["customer_name"].(string)

    data := map[string]any{
        "customerName": name,
        "services":     servicemodel.GetAll(),
        "activeOrders": ordermodel.GetActiveOrders(name),  // Gunakan nama
        "history":      ordermodel.GetHistoryOrders(name), // Gunakan nama
    }

    tmpl, err := template.ParseFiles(
        "views/client/layout.html", 
        "views/client/dashboard.html",
        "views/client/_active_orders.html",
        "views/client/_history_orders.html",
        "views/client/_order_modal.html",
    )
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    tmpl.ExecuteTemplate(w, "client-layout", data)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        // 1. Ambil data dari session
        session, _ := store.Get(r, "customer-session")
        name, _ := session.Values["customer_name"].(string)

        // 2. Ambil data dari Form Modal (tambah delivery_method)
        serviceIDStr := r.FormValue("service_id")
        weightStr := r.FormValue("weight")
        deliveryMethod := r.FormValue("delivery_method") // Ambil data baru ini

        // 3. Konversi tipe data
        serviceID, _ := strconv.Atoi(serviceIDStr)
        weight, _ := strconv.ParseFloat(weightStr, 64)

        // 4. Bungkus ke dalam Struct Order
        newOrder := entities.Order{
            CustomerName: name,
            Weight:       weight,
            Service:      entities.Service{ID: serviceID},
            Status:       "Pending",
            CreatedAt:    time.Now(),
        }

        // 5. Panggil fungsi Create (PASTIKAN ARGUMENNYA SAMA DENGAN DI MODEL)
        // Di sini kita kirim 3 data: order struct, service id, dan metode pengiriman
        if ordermodel.Create(newOrder, serviceID, deliveryMethod) {
            http.Redirect(w, r, "/client/dashboard", http.StatusSeeOther)
        } else {
            http.Error(w, "Gagal memproses pesanan", http.StatusInternalServerError)
        }
    }
}