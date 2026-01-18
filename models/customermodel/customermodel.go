package customermodel

import (
    "Sistem-Laundry/config"
    "Sistem-Laundry/entities"
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