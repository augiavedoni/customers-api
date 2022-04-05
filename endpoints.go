package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// postCustomer adds a customer from JSON received in the request body.
func postCustomer(context *gin.Context) {
	var newCustomer customer

	// Call BindJSON to bind the received JSON to
	// newCustomer.
	if err := context.BindJSON(&newCustomer); err != nil {
		return
	}

	isUserInformationValid := verifyCustomerInformation(newCustomer, context)

	if !isUserInformationValid {
		return
	}

	// Add the new customer to the "database".
	customers = append(customers, newCustomer)
	context.IndentedJSON(http.StatusCreated, newCustomer)
}

// getCustomerById locates the customer whose ID value matches the id
// parameter sent by the client, then returns that customer as a response.
func getCustomerById(context *gin.Context) {
	id := context.Param("id")

	isIdValid := validateId(id, context)

	if !isIdValid {
		return
	}

	customer := searchCustomer(id)

	if customer.ID != "" {
		context.IndentedJSON(http.StatusOK, customer)
		return
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
	}
}

// getCustomers responds with the list of all customers as JSON.
func getCustomers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, customers)
}

func updateCustomer(context *gin.Context) {
	id := context.Param("id")

	isIdValid := validateId(id, context)

	if !isIdValid {
		return
	}

	customerInformation := searchCustomer(id)

	if customerInformation.ID != "" {
		var newCustomer customer

		// Call BindJSON to bind the received JSON to
		// newCustomer.
		if err := context.BindJSON(&newCustomer); err != nil {
			return
		}

		isUserInformationValid := verifyCustomerInformation(newCustomer, context)

		if !isUserInformationValid {
			return
		}

		updateCustomerInformation(id, newCustomer)

		context.IndentedJSON(http.StatusOK, newCustomer)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
	}
}

func deleteCustomer(context *gin.Context) {
	id := context.Param("id")

	isIdValid := validateId(id, context)

	if !isIdValid {
		return
	}

	customer := searchCustomer(id)

	if customer.ID != "" {
		removeCustomer(id)
		context.IndentedJSON(http.StatusOK, gin.H{"message": "Customer deleted successfuly"})
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
	}
}

//Used for testing porpuses
func clearCustomers(context *gin.Context) {
	customers = []customer{}
}
