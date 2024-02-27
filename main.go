package main

import (
	"fmt"

	"github.com/capomanpc/go-todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)
	// controllers.StartMainServer()
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)
}
