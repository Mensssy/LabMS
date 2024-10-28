package model

type UserSecurity struct {
	UserId      string `gorm:"primaryKey"`
	Password    string `gorm:"not null"`
	Salt        string `gorm:"not null"`
	TokenPC     string `gorm:"default:0"`
	TokenMobile string `gorm:"default:0"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
