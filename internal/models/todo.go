package models

type Todo struct {
	TodoID  string `gorm:"primaryKey"`
	Content string
	Done    bool `gorm:"default:false"`
	UserID  string
	User    User `gorm:"references:UserID"`
}
