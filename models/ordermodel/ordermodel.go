package ordermodel

import (
	"Sistem-Laundry/config"
	"Sistem-Laundry/entities"
	"time"
)

// 1. CREATE: Simpan pesanan baru dengan Ongkir & Delivery Method
func Create(order entities.Order, serviceID int, delivery string) bool {
	var pricePerKg int
	err := config.DB.QueryRow("SELECT price_per_kg FROM services WHERE id = ?", serviceID).Scan(&pricePerKg)
	if err != nil {
		return false
	}

	// Logika Ongkir
	ongkir := 0
	if delivery == "Antar Jemput" {
		ongkir = 5000
	}

	totalPrice := int(order.Weight*float64(pricePerKg)) + ongkir

	query := `
        INSERT INTO orders (customer_name, service_id, weight, total_price, status, delivery_method, created_at) 
        VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err = config.DB.Exec(query,
		order.CustomerName,
		serviceID,
		order.Weight,
		totalPrice,
		"Pending", // Status awal selalu Pending
		delivery,
		time.Now(),
	)

	return err == nil
}

// 2. GET ALL: Untuk Dashboard ADMIN (Menampilkan semua kolom)
func GetAll() []entities.Order {
	query := `
        SELECT o.id, o.customer_name, o.weight, o.total_price, o.status, o.created_at, 
               s.name, o.delivery_method
        FROM orders o 
        JOIN services s ON o.service_id = s.id 
        ORDER BY o.id DESC`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var o entities.Order
		// SCAN 8 KOLOM (Sesuai SELECT)
		err := rows.Scan(&o.ID, &o.CustomerName, &o.Weight, &o.TotalPrice, &o.Status, &o.CreatedAt, &o.Service.Name, &o.DeliveryMethod)
		if err != nil {
			continue
		}
		orders = append(orders, o)
	}
	return orders
}

// 3. GET ACTIVE: Untuk Lacak Pesanan CLIENT
func GetActiveOrders(customerName string) []entities.Order {
	query := `
        SELECT o.id, o.customer_name, o.weight, o.total_price, o.status, o.created_at, 
               s.name, o.delivery_method 
        FROM orders o
        JOIN services s ON o.service_id = s.id
        WHERE o.customer_name = ? AND o.status IN ('Pending', 'Cuci', 'Setrika')
        ORDER BY o.created_at DESC`

	rows, err := config.DB.Query(query, customerName)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var o entities.Order
		// SCAN 8 KOLOM
		err := rows.Scan(&o.ID, &o.CustomerName, &o.Weight, &o.TotalPrice, &o.Status, &o.CreatedAt, &o.Service.Name, &o.DeliveryMethod)
		if err != nil {
			continue
		}
		orders = append(orders, o)
	}
	return orders
}

// 4. GET HISTORY: Untuk Riwayat Pesanan CLIENT
func GetHistoryOrders(customerName string) []entities.Order {
	query := `
        SELECT o.id, o.customer_name, o.weight, o.total_price, o.status, o.created_at, 
               s.name, o.delivery_method 
        FROM orders o
        JOIN services s ON o.service_id = s.id
        WHERE o.customer_name = ? AND o.status IN ('Selesai', 'Diambil')
        ORDER BY o.created_at DESC`

	rows, err := config.DB.Query(query, customerName)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var o entities.Order
		// SCAN 8 KOLOM
		err := rows.Scan(&o.ID, &o.CustomerName, &o.Weight, &o.TotalPrice, &o.Status, &o.CreatedAt, &o.Service.Name, &o.DeliveryMethod)
		if err != nil {
			continue
		}
		orders = append(orders, o)
	}
	return orders
}

// 5. UPDATE STATUS: Digunakan Admin untuk proses laundry
func UpdateStatus(id int, status string) bool {
	res, _ := config.DB.Exec("UPDATE orders SET status = ? WHERE id = ?", status, id)
	count, _ := res.RowsAffected()
	return count > 0
}