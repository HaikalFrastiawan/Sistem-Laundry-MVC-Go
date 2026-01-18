package entities

import "time"

type Customer struct {
    ID        int
    Name      string
    Email     string
    Phone     string
    Address   string
    CreatedAt time.Time
}