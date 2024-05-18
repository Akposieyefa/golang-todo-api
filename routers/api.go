package routers

import (
	"akposieyefa/golang-todo-api/handlers"
	"akposieyefa/golang-todo-api/middleware"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter()
	router.Use(middleware.Middleware)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Welcome to the basic todo api using golang and mux router",
			"success": true,
		})
	}).Methods("GET")

	router.HandleFunc("/todos", handlers.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.GetSingelTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("UPDATE")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/auth/register", handlers.CreateUserAccount).Methods("POST")
	router.HandleFunc("/auth/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/user", handlers.GetUserProfile).Methods("GET")

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
