package customer

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jademnp/gofinal/models/customer"
	"github.com/jademnp/gofinal/database"
)
func InitTable() {
	creatTb := `
	CREATE TABLE IF NOT EXISTS customers(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`
	_ , err := database.DB.Exec(creatTb)
	if err != nil {
		panic("Failed to crete customer table")
	}
}
func Create(c *gin.Context)  {
	var input customer.Model
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}
	row := database.DB.QueryRow("INSERT INTO customers (name, email,status) values ($1,$2,$3) RETURNING id",input.Name,input.Email,input.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	input.ID = id
	c.JSON(http.StatusCreated, input)

}

func GetById(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stmt, err := database.DB.Prepare("SELECT id,name,email,status FROM customers WHERE id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result := stmt.QueryRow(id)
	var customer customer.Model
	err = result.Scan(&customer.ID,&customer.Name,&customer.Email,&customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func GetAll(c *gin.Context)  {
	stmt, err := database.DB.Prepare("SELECT id,name,email,status FROM customers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	customers := []customer.Model{}
	for result.Next(){
		var id int
		var name,email,status string
		err = result.Scan(&id,&name,&email,&status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers,customer.Model{id,name,email,status})
	}
	c.JSON(http.StatusOK, customers)
}

func UpdateById(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input customer.Model
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	stmt, err := database.DB.Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id,input.Name,input.Email,input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	input.ID = id
	c.JSON(http.StatusOK, input)
}

func DeleteById(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stmt, err := database.DB.Prepare("DELETE FROM customers WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}