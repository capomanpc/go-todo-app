package main

import (
	"fmt"

	"github.com/capomanpc/go-todo-app/app/models"
)

func main() {
	// log.Println("test")
	// fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "testtest"
	fmt.Println(u)

	u.CreateUser()

	/*
		u, _ := models.GetUser(1)

		u.Name = "Test2"
		u.Email = "Test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	user, _ := models.GetUser(2)
	user.CreateTodo("Firt Todo")

}
