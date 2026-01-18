package entities

type Employee struct {
    ID            int
    Name          string
    Role          string // Kasir, Operator, Kurir
    BaseSalary    int    // Gaji Pokok
    BonusPerOrder int    // Bonus per pengerjaan
}