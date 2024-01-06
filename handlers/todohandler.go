package handlers

import (
	"akposieyefa/golang-todo-api/models"
	"akposieyefa/golang-todo-api/pkg"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todos []models.Todo
	pkg.DB.Find(&todos)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo list fetched successfully",
		"data":    todos,
		"success": true,
	})
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	pkg.DB.Create(&todo)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo created successfully",
		"data":    todo,
		"success": true,
	})
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo models.Todo
	pkg.DB.First(&todo, params["id"])
	json.NewDecoder(r.Body).Decode(&todo)
	pkg.DB.Save(&todo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo details pulled successfully",
		"data":    todo,
		"success": true,
	})
}

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

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var todo models.Todo
	pkg.DB.First(&todo, param["id"])
	pkg.DB.Delete(&todo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo deleted successfully",
		"data":    todo,
		"success": true,
	})
}
