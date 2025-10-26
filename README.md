Go Backend User Management API
==============================

This is a simple Go backend application to manage users with CRUD operations.
It uses Fiber as the web framework, MySQL as the database, Validator for input
validation, and Zap for structured logging.

Features
--------
- Create a user
- Get all users
- Get a user by ID
- Update a user by ID
- Delete a user by ID
- Calculates age from date of birth
- Input validation using go-playground/validator
- Logging with Uber zap

Tech Stack
----------
- Go 1.25+
- Fiber web framework
- MySQL database
- Zap logger
- Go Playground Validator
- SQLC for type-safe SQL queries

Project Structure
-----------------
internal/
    handlers/      # HTTP handlers
    repository/    # Database repository layer
    routes/        # Route definitions
    service/       # Business logic
    logger/        # Logger setup
main.go            # Entry point
go.mod
go.sum
.env               # Environment variables (not committed)

Setup Instructions
------------------

1. Clone the repository
-----------------------
git clone https://github.com/Anagh3/CalculateAge-Go-.git
cd CalculateAge-Go-

2. Create .env file
------------------
Create a `.env` file in the project root with your database credentials:

DB_USER=root
DB_PASS=your_mysql_password
DB_HOST=127.0.0.1:3306
DB_NAME=userdb

> Replace the values with your MySQL credentials.

3. Install dependencies
----------------------
go mod tidy

4. Create MySQL database and table
----------------------------------
Log in to MySQL:

mysql -u root -p

Create the database and table:

CREATE DATABASE userdb;
USE userdb;

CREATE TABLE IF NOT EXISTS users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    dob DATE NOT NULL
);

5. Run the server
-----------------
go run main.go

Server should start on port 3000:

INFO    🚀 Server running {"port": "3000"}

Open your browser or Postman and check:

http://localhost:3000/users

API Endpoints
-------------
Method  Endpoint         Description
------  --------------- ---------------------------------
POST    /users           Create a new user
GET     /users           Get all users
GET     /users/:id       Get user by ID
PUT     /users/:id       Update user by ID
DELETE  /users/:id       Delete user by ID

Example: Create User
-------------------
POST /users
Content-Type: application/json

{
  "name": "John Doe",
  "dob": "1998-10-26"
}

Example: Get User by ID
-----------------------
GET /users/1

Logging
-------
All logs are handled by Uber Zap. Example:

INFO    🚀 Server running {"port": "3000"}

Validation
----------
All user inputs are validated using go-playground/validator:

- name must not be empty
- dob must be a valid date in YYYY-MM-DD format

Testing
-------
Unit tests for service logic:

go test ./internal/service -v

