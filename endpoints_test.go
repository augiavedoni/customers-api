package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func getMockedCustomer() customer {
	mockedCustomer := customer{
		ID:        "1",
		Name:      "Augusto",
		Surname:   "Giavedoni",
		Email:     "augusto.giavedoni@gmail.com",
		Birthdate: "2000-02-20",
	}

	return mockedCustomer
}

func getMockedCustomerResponse() gin.H {
	mockedCustomerInformation := gin.H{
		"id":        "1",
		"name":      "Augusto",
		"surname":   "Giavedoni",
		"email":     "augusto.giavedoni@gmail.com",
		"birthdate": "2000-02-20",
	}

	return mockedCustomerInformation
}

func getMockedUpdatedCustomerInformation() customer {
	mockedCustomer := customer{
		ID:        "1",
		Name:      "Augusto Patricio",
		Surname:   "Giavedoni",
		Email:     "augusto.giavedoni@outlook.com",
		Birthdate: "2000-02-20",
	}

	return mockedCustomer
}

func getMockedUpdatedCustomerInformationResponse() gin.H {
	mockedCustomerInformation := gin.H{
		"id":        "1",
		"name":      "Augusto Patricio",
		"surname":   "Giavedoni",
		"email":     "augusto.giavedoni@outlook.com",
		"birthdate": "2000-02-20",
	}

	return mockedCustomerInformation
}

func getMockedCustomers() []customer {
	var mockedCustomers = []customer{
		{
			ID:        "1",
			Name:      "Augusto",
			Surname:   "Giavedoni",
			Email:     "augusto.giavedoni@gmail.com",
			Birthdate: "2000-02-20",
		},
		{
			ID:        "2",
			Name:      "John",
			Surname:   "Wick",
			Email:     "john.wick@gmail.com",
			Birthdate: "2014-10-24",
		},
	}

	return mockedCustomers
}

func getMockedCustomersResponse() []gin.H {
	mockedCustomerInformation := []gin.H{
		{

			"id":        "1",
			"name":      "Augusto",
			"surname":   "Giavedoni",
			"email":     "augusto.giavedoni@gmail.com",
			"birthdate": "2000-02-20",
		},
		{

			"id":        "2",
			"name":      "John",
			"surname":   "Wick",
			"email":     "john.wick@gmail.com",
			"birthdate": "2014-10-24",
		},
	}

	return mockedCustomerInformation
}

func TestPostCustomer(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(getMockedCustomer())
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)

	assert.Equal(t, 201, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, getMockedCustomerResponse(), got)
}

func TestPostCustomerWithEmptyId(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(customer{
		ID:        "",
		Name:      "Augusto",
		Surname:   "Giavedoni",
		Email:     "augusto.giavedoni@gmail.com",
		Birthdate: "2000-02-20",
	})
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "ID cannot be null or empty"}, got)
}

func TestPostCustomerWithEmptyName(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(customer{
		ID:        "1",
		Name:      "",
		Surname:   "Giavedoni",
		Email:     "augusto.giavedoni@gmail.com",
		Birthdate: "2000-02-20",
	})
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Name cannot be null or empty"}, got)
}

func TestPostCustomerWithEmptySurname(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(customer{
		ID:        "1",
		Name:      "Augusto",
		Surname:   "",
		Email:     "augusto.giavedoni@gmail.com",
		Birthdate: "2000-02-20",
	})
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Surname cannot be null or empty"}, got)
}

func TestPostCustomerWithIncorrectEmail(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(customer{
		ID:        "1",
		Name:      "Augusto",
		Surname:   "Giavedoni",
		Email:     "a.a",
		Birthdate: "2000-02-20",
	})
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Email is not valid"}, got)
}

func TestPostCustomerWithEmptyEmail(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(customer{
		ID:        "1",
		Name:      "Augusto",
		Surname:   "Giavedoni",
		Email:     "",
		Birthdate: "2000-02-20",
	})
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Email cannot be null or empty"}, got)
}

func TestPostCustomerWithInvalidBirthdate(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(customer{
		ID:        "1",
		Name:      "Augusto",
		Surname:   "Giavedoni",
		Email:     "augusto.giavedoni@gmail.com",
		Birthdate: "/",
	})
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Birthdate is not valid"}, got)
}

func TestPostCustomerWithEmptyBirdthdate(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(customer{
		ID:        "1",
		Name:      "Augusto",
		Surname:   "Giavedoni",
		Email:     "augusto.giavedoni@gmail.com",
		Birthdate: "",
	})
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Birthdate cannot be null or empty"}, got)
}

func TestPostCustomerWithBirdthdateAfterToday(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(customer{
		ID:        "1",
		Name:      "Augusto",
		Surname:   "Giavedoni",
		Email:     "augusto.giavedoni@gmail.com",
		Birthdate: "2023-05-13",
	})
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Birthdate cannot be after today"}, got)
}
func TestGetCustomerByIdSuccessfuly(t *testing.T) {
	postCustomerForTesting(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Method: "GET",
	}
	context.Request.Header.Set("Content-Type", "application/json")
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	getCustomerById(context)

	assert.Equal(t, 200, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, getMockedCustomerResponse(), got)
	clearCustomers(context)
}

func TestGetCustomerByIdWithNonExistentId(t *testing.T) {
	postCustomerForTesting(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Method: "GET",
	}
	context.Request.Header.Set("Content-Type", "application/json")
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "2",
		},
	}

	getCustomerById(context)

	assert.Equal(t, 404, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Customer not found"}, got)
	clearCustomers(context)
}

func TestGetCustomerByIdWithWrongId(t *testing.T) {
	postCustomerForTesting(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Method: "GET",
	}
	context.Request.Header.Set("Content-Type", "application/json")
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "/",
		},
	}

	getCustomerById(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "ID is not valid"}, got)
	clearCustomers(context)
}

func TestGetCustomerByIdWithEmptyId(t *testing.T) {
	postCustomerForTesting(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Method: "GET",
	}
	context.Request.Header.Set("Content-Type", "application/json")
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "",
		},
	}

	getCustomerById(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "ID must not be empty"}, got)
	clearCustomers(context)
}

func postCustomerForTesting(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, marshalError := json.Marshal(getMockedCustomer())
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	postCustomer(context)
}

func TestGetCustomersSuccessfuly(t *testing.T) {
	postCustomersForTesting(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	getCustomers(context)

	assert.Equal(t, 200, writer.Code)

	var got []gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, getMockedCustomersResponse(), got)
	clearCustomers(context)
}

func postCustomersForTesting(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	mockedCustomers := getMockedCustomers()

	for _, mockedCustomer := range mockedCustomers {
		jsonbytes, marshalError := json.Marshal(mockedCustomer)

		if marshalError != nil {
			panic(marshalError)
		}

		context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

		postCustomer(context)
	}
}

func TestUpdateCustomerWithCorrectInformation(t *testing.T) {
	TestPostCustomer(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Method: "PUT",
	}
	context.Request.Header.Set("Content-Type", "application/json")
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	jsonbytes, marshalError := json.Marshal(getMockedUpdatedCustomerInformation())
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	updateCustomer(context)

	assert.Equal(t, 200, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, getMockedUpdatedCustomerInformationResponse(), got)
	clearCustomers(context)
}

func TestUpdateCustomerWithNonExistingId(t *testing.T) {
	TestPostCustomer(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Method: "PUT",
	}
	context.Request.Header.Set("Content-Type", "application/json")
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "3",
		},
	}

	jsonbytes, marshalError := json.Marshal(getMockedUpdatedCustomerInformation())
	if marshalError != nil {
		panic(marshalError)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	updateCustomer(context)

	assert.Equal(t, 404, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Customer not found"}, got)
	clearCustomers(context)
}

func TestDeleteCustomerSuccessfuly(t *testing.T) {
	postCustomerForTesting(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Method: "DELETE",
	}
	context.Request.Header.Set("Content-Type", "application/json")
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	deleteCustomer(context)

	assert.Equal(t, 200, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"message": "Customer deleted successfuly"}, got)
}

func TestDeleteCustomerWithNonExistentId(t *testing.T) {
	postCustomerForTesting(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Method: "DELETE",
	}
	context.Request.Header.Set("Content-Type", "application/json")
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "2",
		},
	}

	deleteCustomer(context)

	assert.Equal(t, 404, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "Customer not found"}, got)
}

func TestDeleteCustomerWithInvalidId(t *testing.T) {
	postCustomerForTesting(t)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Method: "DELETE",
	}
	context.Request.Header.Set("Content-Type", "application/json")
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "/",
		},
	}

	deleteCustomer(context)

	assert.Equal(t, 400, writer.Code)

	var got gin.H

	err := json.Unmarshal(writer.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"error": "ID is not valid"}, got)
}
