package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jademnp/gofinal/database"
	"github.com/jademnp/gofinal/controller/customer"
  )

func main() {
	fmt.Println("customer service")

	r := gin.Default()

	database.Connect()
	customer.InitTable()

	r.POST("/customers", customer.Create)
	r.GET("/customers/:id", customer.GetById)
	r.GET("/customers", customer.GetAll)
	r.PUT("/customers/:id", customer.UpdateById)
	r.DELETE("/customers/:id", customer.DeleteById)
	r.Run(":2009")
	//run port ":2009"
}
