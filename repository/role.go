package repository

import (
	"github.com/mieramensatu/todolist-be/model"
	"gorm.io/gorm"
)

func GetAllRole(db *gorm.DB) ([]model.Roles, error) {
	var roles []model.Roles
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRoleById(db *gorm.DB, id string) (model.Roles, error) {
	var role model.Roles
	if err := db.First(&role, id).Error; err != nil {
		return role, err
	}
	return role, nil
}
