package entities

import "time"

type Order struct {
    ID           int
    CustomerName string
    Service      Service 
    Weight       float64
    TotalPrice   int
    Status       string
    CreatedAt    time.Time
}