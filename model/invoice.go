package model

import "time"

type Invoice struct {
	InvoiceId     int       `gorm:"primaryKey" json:"invoiceId"`
	State         int       `gorm:"default:1" json:"state"`
	SubmitterName string    `gorm:"not null" json:"submitterName"`
	Type          string    `gorm:"default:none" json:"type"`
	ItemName      string    `gorm:"not null" json:"itemName"`
	Amount        float64   `gorm:"default:0.00" json:"amount"`
	Usage         string    `gorm:"not null" json:"usage"`
	DeliveryDate  time.Time `gorm:"default:2004-2-18 00:00:00" json:"deliveryDate"`

	UserId string `gorm:"not null" json:"-"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
