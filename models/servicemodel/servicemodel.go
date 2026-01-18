package servicemodel

import (
	"Sistem-Laundry/config"
	"Sistem-Laundry/entities"
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
