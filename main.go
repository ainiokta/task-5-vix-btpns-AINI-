package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/ainiokta/task-5-vix-btpns-AINI-/controllers"
	
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	// http.HandleFunc("/logout", authcontroller.Logout)
	// http.HandleFunc("/register", authcontroller.register)

	fmt.Println("server jalan pada http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
