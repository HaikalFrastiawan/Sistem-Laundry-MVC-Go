package customermodel

import (
	"Sistem-Laundry/config"
	"Sistem-Laundry/entities"
	"golang.org/x/crypto/bcrypt"
)

func GetAll() []entities.Customer {
	// Sebutkan nama kolom secara spesifik, jangan pakai SELECT *
	rows, err := config.DB.Query("SELECT id, name, email, phone, address FROM customers")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		var c entities.Customer
		// Pastikan urutan ini: ID, Name, Email, Phone, Address
		err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.Address)
		if err != nil {
			continue
		}
		customers = append(customers, c)
	}
	return customers
}

func Login(phone string, password string) (entities.Customer, bool) {
    var customer entities.Customer

    // Ambil data termasuk kolom password
    query := "SELECT id, name, email, phone, address, password FROM customers WHERE phone = ? LIMIT 1"
    row := config.DB.QueryRow(query, phone)

    // Scan data. Jika error, berarti nomor HP tidak ada
    err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.Address, &customer.Password)
    if err != nil {
        return customer, false
    }

    // Bandingkan password input (plain) dengan password di DB (hash)
    err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
    if err != nil {
        return customer, false // Password salah
    }

    return customer, true
}

func Create(customer entities.Customer) bool {
    // 1. Enkripsi password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
    if err != nil {
        return false
    }

    // 2. Simpan ke database
    // Pastikan urutan kolom (name, email, phone, address, password) sesuai dengan tabelmu
    query := "INSERT INTO customers (name, email, phone, address, password) VALUES (?, ?, ?, ?, ?)"
    
    // Gunakan string(hashedPassword) untuk menyimpan hasil enkripsi
    _, err = config.DB.Exec(query, 
        customer.Name, 
        customer.Email, 
        customer.Phone, 
        customer.Address, 
        string(hashedPassword),
    )

    if err != nil {
        return false
    }
    return true
}

