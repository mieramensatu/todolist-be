package repository

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mieramensatu/todolist-be/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetAllUsers(db *gorm.DB) ([]model.Users, error) {
	var users []model.Users
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(db *gorm.DB, user *model.Users) error {
	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Jika id_role tidak diisi, atur nilainya ke 2 (atau nilai default yang diinginkan)
	if user.IdRole == 0 {
		user.IdRole = 2 // atau nilai default yang diinginkan
	}

	// Simpan user ke database
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByUsername(db *gorm.DB, username string) (*model.Users, error) {
	var user model.Users
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserById(db *gorm.DB, userID uint) (*model.Users, error) {
	var user model.Users
	result := db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GenerateToken(user *model.Users) (string, error) {
	claims := &model.JWTClaims{
		IdUser: user.IdUser,
		IdRole: user.IdRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // Token berlaku selama 1 jam
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DeleteUserById(db *gorm.DB, id string) error {
	// Menghapus user dari database berdasarkan ID yang diberikan
	if err := db.Where("id_user = ?", id).Delete(&model.Users{}).Error; err != nil {
		return err
	}
	return nil
}


func PromoteUserToAdmin(db *gorm.DB, userID uint) error {
	result := db.Model(&model.Users{}).Where("id_user = ?", userID).Update("id_role", 1)
	if result.Error != nil {
		return result.Error
	}
	return nil
}