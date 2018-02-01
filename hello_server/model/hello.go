package model

type Hello struct {
	ID    int64
	Hello string `gorm:"type:varchar(5000)"`
}
