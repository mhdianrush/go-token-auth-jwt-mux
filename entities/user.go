package entities

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"varchar(100)" json:"full_name"`
	Username string `gorm:"varchar(100)" json:"username"`
	Password string `gorm:"varchar(100)" json:"password"`
}
