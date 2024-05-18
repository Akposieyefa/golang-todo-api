package pkg

import "akposieyefa/golang-todo-api/models"

func MigrateTables() {
	ConnectToDB()
	DB.AutoMigrate(&models.User{}, &models.Todo{})
}
