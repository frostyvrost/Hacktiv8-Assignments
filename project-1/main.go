package main

import (
	"errors"
	"net/http"
	_ "project-1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{}

// @title	Todo
// @description	A ToDo Application API for Hacktiv8 Project

func main() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run("localhost:8080")
}

// @Summary Get all todos
// @ID get-all-todos
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// @Summary Add a new todo
// @ID add-todo
// @Produce json
// @Param todo body Todo true "Add todo"
// @Success 200 {object} Todo
// @Router /todos [post]
func addTodos(context *gin.Context) {
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodosById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}

// @Summary Get a todo by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} Todo
// @Router /todos/{id} [get]
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

// @Summary Update a todo by ID
// @ID update-todo-by-id
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} Todo
// @Router /todos/{id} [patch]
func updateTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} Todo
// @Router /todos/{id} [delete]
func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	index := -1

	for i, t := range todos {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	todos = append(todos[:index], todos[index+1:]...)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
