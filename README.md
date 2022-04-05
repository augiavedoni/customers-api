# customers-api
Customer relationship management API required for the CodeSherpas coding assessment. I choose Golang as the lenguage for this project. Hope you like it :)

This API aims to represent the organization of customers information. Each customer is represented as follows:

- ID: string
- Name: string
- Surname: string
- Email: string
- Birthdate (YYYY-MM-DD): string

You're going to be able to add new customers, to retrieve a customer information by its ID, query all the customers saved on the system, delete customers and update a customer information.

## Steps to run this API on your machine:

IMPORTANT NOTE: you'll need Docker installed on your machine. Otherwise, if you have Go installed, you'll have to start the server manually.

1. Clone the repository
2. Open a new terminal window on the folder you cloned the repository
3. Run the following command: docker build --tag customers-api .
4. Once the Docker image is built, run the following command: docker run --publish 8080:8080 customers-api

If everything went well, you'll have the API running on your machine on the port 8080. Now you can open a new terminal window and start playing with it. Have fun!

## Endpoints:

Using the command `curl http://localhost:8080/{endpoint}` you can interact with the API. The avaible endpoints are the following:

- **POST /customer**: this endpoint expects you to send as the body of the request the information about a customer. The estructure of the model that represents a customer was explained earlier. It returns the customer information that was added to the system.
- **GET /customer/id**: this endpoint requires an ID as a parameter. It returns the information about a customer. For example: `curl http://localhost:8080/customer/1`
- **GET /customers**: it returns the information about all the customers that are present in the system. For example: `curl http://localhost:8080/customers`
- **PUT /customer/id**: this endpoint requires an ID as a parameter and all the updated information about the customer (all fields are required). It returns the updated information about the customer.
- **DELETE /customer/id**: this endpoint requires an ID as a parameter. It returns wheter the customer was deleted from the system or if the customer wasn't found. For example: `curl -X DELETE http://localhost:8080/customer/1`

### Things to consider:

- The ID is verified and can't be null, empty (except for the POST method to /customer) or a special character. If so, the API will return a 400 code (bad request) and a message.
- When adding a customer to the system, some validations are run prior to adding the customer. For example, all fields are required and the birthdate of the customer can't be after the actual date or have a different format that the one indicated before. Besides that, the email is verified so it won't accept invalid email addresses.
- If a customer is not found on the system, a 404 code (not found) and a message are going to be returned.
- It's a small project and it can have more and better validations. If you have one in mind, I'll be happy to hear from you :).












