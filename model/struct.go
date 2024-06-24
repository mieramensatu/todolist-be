package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Task struct {
	IdTask    uint      `json:"id_task" gorm:"primaryKey"`
	Judul     string    `json:"judul"`
	Deskripsi string    `json:"deskripsi"`
	DueDate   time.Time `json:"due_date"`
	Completed bool      `json:"completed"` // Ensure this field is a boolean
	IdUser    uint      `json:"id_user"`
}

type GetJoinTask struct {
	IdTask    uint      `json:"id_task"`
	Judul     string    `json:"judul"`
	Deskripsi string    `json:"deskripsi"`
	DueDate   time.Time `json:"due_date"`
	Completed bool      `json:"completed"` // Ensure this field is a boolean
	IdUser    uint      `json:"id_user"`
	Nama      string    `json:"nama"`
}

type Users struct {
	IdUser   uint   `gorm:"primaryKey;column:id_user" json:"id_user"`
	IdRole   int    `gorm:"column:id_role" json:"id_role"`
	Nama     string `gorm:"column:nama" json:"nama"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
}

type Roles struct {
	IdRole int    `gorm:"primaryKey;column:id_role" json:"id_role"`
	Nama   string `gorm:"column:nama" json:"nama"`
}

type JWTClaims struct {
	jwt.StandardClaims
	IdUser uint `json:"id_user"`
	IdRole int  `json:"id_role"`
}
