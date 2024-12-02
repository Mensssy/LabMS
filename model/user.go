package model

type User struct {
	UserId   string `gorm:"primaryKey;type: varchar(255);comment: telephone number" json:"-"`
	UserName string `gorm:"not null;type: varchar(20)"`
}
