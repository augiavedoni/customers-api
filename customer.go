package main

import (
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Birthdate string `json:"birthdate"`
}

func searchCustomer(id string) customer {
	foundedCustomer := customer{
		ID:        "",
		Name:      "",
		Surname:   "",
		Email:     "",
		Birthdate: "",
	}

	// Loop over the list of customers, looking for
	// a customer whose ID value matches the parameter.
	for _, customer := range customers {
		if customer.ID == id {
			foundedCustomer = customer
		}
	}

	return foundedCustomer
}

func updateCustomerInformation(id string, newCustomerInformation customer) {
	index := -1

	// Loop over the list of customers, looking for
	// a customer whose ID value matches the parameter.
	for i, customer := range customers {
		if customer.ID == id {
			index = i
		}
	}

	customers[index].Name = newCustomerInformation.Name
	customers[index].Surname = newCustomerInformation.Surname
	customers[index].Email = newCustomerInformation.Email
	customers[index].Birthdate = newCustomerInformation.Birthdate
}

func removeCustomer(index string) {
	parsedIndex, err := strconv.ParseInt(index, 6, 12)

	if err != nil {
		return
	}

	if parsedIndex != 0 {
		parsedIndex = parsedIndex - 1
	}

	auxiliaryList := make([]customer, 0)
	auxiliaryList = append(auxiliaryList, customers[:parsedIndex]...)

	customers = append(auxiliaryList, customers[parsedIndex+1:]...)
}

func verifyCustomerInformation(customerInformation customer, context *gin.Context) bool {
	if customerInformation.ID == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be null or empty"})
		return false
	} else if customerInformation.Name == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Name cannot be null or empty"})
		return false
	} else if customerInformation.Surname == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Surname cannot be null or empty"})
		return false
	} else if customerInformation.Email == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Email cannot be null or empty"})
		return false
	} else if customerInformation.Birthdate == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Birthdate cannot be null or empty"})
		return false
	}

	if !validateCustomerEmail(customerInformation.Email, context) {
		return false
	}

	if !validateCustomerBirthdate(customerInformation.Birthdate, context) {
		return false
	}

	return true
}

func validateCustomerEmail(email string, context *gin.Context) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Email is not valid"})
		return false
	} else {
		return true
	}
}

func validateCustomerBirthdate(birthdate string, context *gin.Context) bool {
	customerBirthdate, err := time.Parse("2006-01-02", birthdate)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Birthdate is not valid"})
		return false
	} else if customerBirthdate.After(time.Now()) {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Birthdate cannot be after today"})
		return false
	}

	return true
}
