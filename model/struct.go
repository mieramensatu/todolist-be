package model

import "github.com/dgrijalva/jwt-go"

type Task struct {
	IdTask    int    `gorm:"primaryKey;column:id_task" json:"id_task"`
	IdUser    uint   `gorm:"column:id_user" json:"id_user"` // Changed to uint
	Judul     string `gorm:"column:judul" json:"judul"`
	Deskripsi string `gorm:"column:deskripsi" json:"deskripsi"`
	DueDate   string `gorm:"column:due_date" json:"due_date"`
	Completed string `gorm:"column:completed" json:"completed"`
}

type GetJoinTask struct {
	IdTask    int    `gorm:"primaryKey;column:id_task" json:"id_task"`
	IdUser    uint   `gorm:"column:id_user" json:"id_user"` // Changed to uint
	Nama      string `gorm:"column:nama" json:"nama"`
	Judul     string `gorm:"column:judul" json:"judul"`
	Deskripsi string `gorm:"column:deskripsi" json:"deskripsi"`
	DueDate   string `gorm:"column:due_date" json:"due_date"`
	Completed string `gorm:"column:completed" json:"completed"`
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
