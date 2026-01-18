package servicemodel

import (
	"Sistem-Laundry/config"
	"Sistem-Laundry/entities"
	"database/sql"

)

func GetAll() ([]entities.Service) {
	rows, err := config.DB.Query(
		`SELECT id, name, price_per_kg FROM services`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var services []entities.Service

	for rows.Next() {
		var service entities.Service
		err := rows.Scan(
			&service.ID,
			&service.Name,
			&service.PricePerKg,
		)
		if err != nil {
			panic(err)
		}
		services = append(services, service)
	}
	return services
}

// 1. CREATE: Menambah data baru
func Create(service entities.Service) bool {
    result, err := config.DB.Exec(`
        INSERT INTO services (name, price_per_kg) 
        VALUES (?, ?)`,
        service.Name, service.PricePerKg,
    )

    if err != nil {
        panic(err) // Sebaiknya log error, tapi panic oke untuk dev
    }

    lastInsertId, _ := result.LastInsertId()
    return lastInsertId > 0
}

// 2. DETAIL: Mengambil 1 data berdasarkan ID (Penting untuk mengisi form Edit)
func Detail(id int) entities.Service {
    row := config.DB.QueryRow(`SELECT id, name, price_per_kg FROM services WHERE id = ?`, id)

    var service entities.Service
    err := row.Scan(&service.ID, &service.Name, &service.PricePerKg)

    if err != nil {
        if err == sql.ErrNoRows {
            return entities.Service{} // Kembalikan struct kosong jika tidak ketemu
        }
        panic(err)
    }

    return service
}

// 3. UPDATE: Mengupdate data yang sudah ada
func Update(id int, service entities.Service) bool {
    query, err := config.DB.Exec(`
        UPDATE services SET name = ?, price_per_kg = ? WHERE id = ?`,
        service.Name, service.PricePerKg, id,
    )

    if err != nil {
        panic(err)
    }

    rowCount, _ := query.RowsAffected()
    return rowCount > 0
}

// 4. DELETE: Menghapus data
func Delete(id int) bool {
    result, err := config.DB.Exec(`DELETE FROM services WHERE id = ?`, id)
    if err != nil {
        panic(err)
    }

    rowCount, _ := result.RowsAffected()
    return rowCount > 0
}