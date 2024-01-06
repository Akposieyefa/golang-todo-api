package main

import (
	"akposieyefa/golang-todo-api/pkg"
	"akposieyefa/golang-todo-api/routers"
)

func main() {
	pkg.LoadEnv()
	pkg.ConnectToDB()
	routers.Router()
}
