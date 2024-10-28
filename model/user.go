package model

type User struct {
	UserId   string `gorm:"primaryKey;type: varchar(255);comment: telephone number"`
	UserName string `gorm:"not null;type: varchar(10)"`
}
