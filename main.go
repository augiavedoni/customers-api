package main

import "github.com/gin-gonic/gin"

var customers = []customer{}

func main() {
	router := gin.Default()
	router.POST("/customer", postCustomer)
	router.GET("/customers", getCustomers)
	router.GET("/customer/:id", getCustomerById)
	router.PUT("/customer/:id", updateCustomer)
	router.DELETE("/customer/:id", deleteCustomer)

	router.Run(":8080")
}
