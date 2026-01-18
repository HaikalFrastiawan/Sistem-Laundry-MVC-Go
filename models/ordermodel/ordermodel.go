package ordermodel

import (
    "Sistem-Laundry/config"
    "Sistem-Laundry/entities"
    "database/sql"
    "time"
)

// 1. CREATE: Simpan pesanan baru & Hitung Total Harga otomatis
func Create(order entities.Order) bool {
    // A. Cari harga per kg dulu dari tabel services
    var pricePerKg int
    err := config.DB.QueryRow("SELECT price_per_kg FROM services WHERE id = ?", order.Service.ID).Scan(&pricePerKg)
    
    if err != nil {
        return false // Service tidak ditemukan
    }

    // B. Hitung Total (Berat x Harga)
    totalPrice := int(order.Weight * float64(pricePerKg))

    // C. Masukkan ke Database
    result, err := config.DB.Exec(`
        INSERT INTO orders (customer_name, service_id, weight, total_price, status, created_at) 
        VALUES (?, ?, ?, ?, 'Pending', ?)`,
        order.CustomerName, order.Service.ID, order.Weight, totalPrice, time.Now(),
    )

    if err != nil {
        return false
    }

    lastInsertId, _ := result.LastInsertId()
    return lastInsertId > 0
}

// 2. GET DATA: Untuk Tracking (Proses) & History (Selesai)
func GetByStatusAndCustomer(customerName string, statusCategory string) []entities.Order {
    var rows *sql.Rows
    var err error

    // Query JOIN supaya kita dapat Nama Service-nya, bukan cuma ID
    query := `
        SELECT o.id, o.customer_name, o.weight, o.total_price, o.status, o.created_at, 
               s.name, s.price_per_kg 
        FROM orders o
        JOIN services s ON o.service_id = s.id
        WHERE o.customer_name = ? AND `

    if statusCategory == "Proses" {
        // Ambil yang statusnya MASIH JALAN (Bukan Selesai/Diambil)
        query += `o.status IN ('Pending', 'Cuci', 'Setrika') ORDER BY o.created_at DESC`
    } else {
        // Ambil yang SUDAH BERES
        query += `o.status IN ('Selesai', 'Diambil') ORDER BY o.created_at DESC`
    }

    rows, err = config.DB.Query(query, customerName)
    if err != nil {
        return nil
    }
    defer rows.Close()

    var orders []entities.Order

    for rows.Next() {
        var order entities.Order
        // Scan data hasil gabungan tabel
        err := rows.Scan(
            &order.ID,
            &order.CustomerName,
            &order.Weight,
            &order.TotalPrice,
            &order.Status,
            &order.CreatedAt,
            &order.Service.Name,       // Masuk ke struct Service
            &order.Service.PricePerKg, 
        )

        if err != nil {
            continue
        }
        orders = append(orders, order)
    }

    return orders
}

func GetAll() []entities.Order {
    // Gunakan JOIN agar data Service.Name tidak kosong
    rows, err := config.DB.Query(`
        SELECT o.id, o.customer_name, o.weight, o.total_price, o.status, o.created_at, s.name 
        FROM orders o 
        JOIN services s ON o.service_id = s.id 
        ORDER BY o.id DESC`)
    
    if err != nil {
        return nil
    }
    defer rows.Close()

    var orders []entities.Order
    for rows.Next() {
        var order entities.Order
        // Pastikan urutan Scan sama dengan urutan SELECT di atas
        err := rows.Scan(&order.ID, &order.CustomerName, &order.Weight, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.Service.Name)
        if err != nil {
            continue
        }
        orders = append(orders, order)
    }
    return orders
}

// Fungsi update status oleh Admin/Operator
func UpdateStatus(id int, status string) bool {
    res, _ := config.DB.Exec("UPDATE orders SET status = ? WHERE id = ?", status, id)
    count, _ := res.RowsAffected()
    return count > 0
}