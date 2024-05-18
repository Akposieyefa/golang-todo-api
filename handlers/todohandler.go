package handlers

import (
	"akposieyefa/golang-todo-api/helpers"
	"akposieyefa/golang-todo-api/models"
	"akposieyefa/golang-todo-api/pkg"
	"akposieyefa/golang-todo-api/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// get all todos by auth user
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authUser := services.AuthUser(w, r)
	userID := authUser["user"].(models.User).ID

	var todos []models.Todo
	if err := pkg.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Todo list fetched successfully",
		"data":    todos,
		"success": true,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// create todos
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	err := helpers.Validate(todo)

	authUser := services.AuthUser(w, r)
	todo.UserID = authUser["user"].(models.User).ID

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	pkg.DB.Create(&todo)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo created successfully",
		"data":    todo,
		"success": true,
	})
}

// update todos
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo models.Todo
	pkg.DB.First(&todo, params["id"])
	json.NewDecoder(r.Body).Decode(&todo)

	authUser := services.AuthUser(w, r)
	userID := authUser["user"].(models.User).ID

	if todo.UserID != userID {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Sorry you are not allowed to update this todo",
			"success": false,
		})
		return
	}

	pkg.DB.Save(&todo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo details pulled successfully",
		"data":    todo,
		"success": true,
	})
}

// get single todo
func GetSingelTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo models.Todo
	pkg.DB.First(&todo, params["id"])
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo details pulled successfully",
		"data":    todo,
		"success": true,
	})
}

// delete todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var todo models.Todo
	pkg.DB.First(&todo, param["id"])

	authUser := services.AuthUser(w, r)
	userID := authUser["user"].(models.User).ID

	if todo.UserID != userID {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Sorry you are not allowed to delete this todo",
			"success": false,
		})
		return
	}
	pkg.DB.Delete(&todo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo deleted successfully",
		"data":    todo,
		"success": true,
	})
}
