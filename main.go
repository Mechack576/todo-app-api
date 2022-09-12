package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

type todo struct{
	ID			string		`json:"id"`
	Item		string		`json:"title"`
	Completed	bool		`json:"completed"`
}

var todos = [] todo{
	{ID:"1", Item:"Clean Room", Completed:false},
	{ID:"2", Item:"Leetcode", Completed:false},
	{ID:"3", Item:"Grocieries Shopping", Completed:false},
}

func getTodoList (c *gin.Context){
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodoById (id string) (*todo,error){
	for i,t := range todos {
		if t.ID == id{
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func todoById (c *gin.Context){
	id := c.Param("id")
	todo,err := getTodoById(id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func addTodo(c *gin.Context){
	var newTodo todo								//create a newTodo of type todo struct
	if error := c.BindJSON(&newTodo); error != nil { //bind the json data of newTodo (pointer to newTodo var)
		return // return if we get an error
	}
	// if no error append todo to the list
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func changeTodoStatus(c *gin.Context){
	id := c.Param("id")
	todo,err := getTodoById(id)

	if err !=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found."})
		return
	}
	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos",getTodoList)
	router.GET("/todos/:id",todoById)
	router.PATCH("/todos/:id",changeTodoStatus)
	router.POST("/todos",addTodo)
	router.Run("localhost:9090")
}