package model

type User struct {
	UserId   string `gorm:"primaryKey; type: varchar(255)" json:"-"`
	UserName string `gorm:"not null; type: varchar(20)" json:"userName"`
	UserType string `gorm:"not null; default: 'STU'; type: varchar(5)" json:"userType"`
}
