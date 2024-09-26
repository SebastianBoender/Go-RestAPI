package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var users = []user{
	{ID: "1", FirstName: "Sebastian", LastName: "Boender", Email: "sebastianboender@Hotmail.com"},
	{ID: "2", FirstName: "Ali", LastName: "Yilmaz", Email: "aliyilmaz@Hotmail.com"},
	{ID: "3", FirstName: "Jelle", LastName: "van Den Berg", Email: "jellevandenberg@Hotmail.com"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func newUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", newUser)
	router.Run("localhost:8080")
}
