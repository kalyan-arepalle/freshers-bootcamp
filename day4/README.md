DataBase : mySql

Process of implementing the project:
1. Pull the repo to the src folder of your go folder in home directory
2. Start mySql.server and create database named retailer
3. Make changes in Config folders' Database.go file corresponding your server authentication
4. Start the project in GoLand Terminal with command: go run main.go
5. Use POSTMAN to perform operations on products or making orders or adding customers.

It can be noticed that operations made are reflected in database, check using MySQL workbench

APIs Documentation:
GET: /product

Get all the Products

POST: /product

Add a new product along with price and quantity

GET: /product/:id

Get Product by its ID

PATCH: /product/:id

Update a Product by its ID

DELETE: /product-api/product/:id

Delete a Product

POST: /order

Order some product from the list using customerid

GET : /order/:id

Get the order corresponding to the id

GET : /order

get info of all the Orders

POST: /customer

Add a customer to the database

GET : /customer/:id

Get customer details from his id