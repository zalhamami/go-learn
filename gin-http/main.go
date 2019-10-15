package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World")
	router := gin.Default()
	router.GET("/", index)
	router.POST("/", post)
	router.GET("/query", withQuery)
	router.GET("/param/:name", withParam)
	router.Run()
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello, it's gin",
	})
}

func post(c *gin.Context) {
	payload := c.Request.Body
	value, err := ioutil.ReadAll(payload)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": string(value),
	})
}

func withQuery(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func withParam(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
