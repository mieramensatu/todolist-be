package repository

import (
	"github.com/mieramensatu/todolist-be/model"
	"gorm.io/gorm"
)

func GetAllTasksByUserId(db *gorm.DB, userId uint, isAdmin bool) ([]model.GetJoinTask, error) {
	var tasks []model.GetJoinTask
	query := db.Table("task").Select("task.id_task, task.judul, task.deskripsi, task.due_date, task.completed, users.id_user, users.nama").Joins("JOIN users ON task.id_user = users.id_user")
	if !isAdmin {
		query = query.Where("task.id_user = ?", userId)
	}
	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetAllTask(db *gorm.DB) ([]model.GetJoinTask, error) {
	var tasks []model.GetJoinTask
	if err := db.
		Table("task").
		Select("task.id_task, task.judul, task.deskripsi, task.due_date, task.completed, users.id_user, users.nama").
		Joins("JOIN users ON task.id_user = users.id_user").
		Find(&tasks).
		Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskById(db *gorm.DB, id string) (model.Task, error) {
	var task model.Task
	if err := db.First(&task, id).Error; err != nil {
		return task, err
	}
	return task, nil
}

func GetTasksByUserId(db *gorm.DB, userID uint) ([]model.Task, error) {
	var tasks []model.Task
	if err := db.Where("id_user = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func InsertTask(db *gorm.DB, task *model.Task) error {
	if err := db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTask(db *gorm.DB, id string, updatedTask model.Task) error {
	if err := db.Model(&model.Task{}).Where("id_task = ?", id).Updates(updatedTask).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTask(db *gorm.DB, id string) error {
	if err := db.Delete(&model.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
