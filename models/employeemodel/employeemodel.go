package employeemodel

import (
    "Sistem-Laundry/config"
    "Sistem-Laundry/entities"
)

func GetAll() []entities.Employee {
    rows, err := config.DB.Query("SELECT id, name, role, base_salary, bonus_per_order FROM employees")
    if err != nil {
        return nil
    }
    defer rows.Close()

    var employees []entities.Employee
    for rows.Next() {
        var emp entities.Employee
        rows.Scan(&emp.ID, &emp.Name, &emp.Role, &emp.BaseSalary, &emp.BonusPerOrder)
        employees = append(employees, emp)
    }
    return employees
}