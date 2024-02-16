package main

import (
	"github.com/capomanpc/go-todo-app/app/models"
)

func main() {
	// log.Println("test")
	// fmt.Println(models.Db)

	/*
		u := &models.User{}
		u.Name = "test2"
		u.Email = "test2@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
		/*

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


			user, _ := models.GetUser(2)
			user.CreateTodo("Firt Todo")
	*/

	/*
		t, _ := models.GetUser(1)
		fmt.Println(t)
	*/

	/*
		user, _ := models.GetUser(2)
		user.CreateTodo("Second Todo")

		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		user, _ := models.GetUser(3)
		user.CreateTodo("Third Todo")

		user2, _ := models.GetUser(2)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	t, _ := models.GetTodo(3)
	t.DeleteTodo()

}
