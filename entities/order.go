package entities

type Order struct {
	ID         int
	CustomerID int
	ServiceID  int
	Weight     float64
	TotalPrice int
	Status     string
}
