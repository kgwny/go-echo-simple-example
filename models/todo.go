package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string
}

func CreateTodo(content string) (err error) {
	todo := Todo{
		Content: content,
	}
	DB.Create(&todo)
	return err
}

func DeleteTodo(id int) (err error) {
	DB.Delete(&Todo{}, id)
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	DB.Find(&todo, id)
	return todo, err
}

func UpdateTodo(t Todo) (err error) {
	DB.Save(&t)
	return err
}
