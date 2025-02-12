<div id="header" align="center">
  <img src="https://i.gifer.com/S9Ih.gif" width="210"/>
</div>

# ToDoList APP
- - -
### About:
- Web application that is created by design REST API
- In the application you can: 
  1. Interacting with the application via the HTTP client for API testing
  2. Register, log in to your account.
  3. Receive a token to work with your lists and tasks.
  4. perform CRUD operations.
- - -
### Technologies, frameworks and tools:

- PostgreSQL with lib/sqlx
- Framework Gin
- Implementation sign-up, sign-up, JWT.
- Launch docker, writing docker file, docker-compose
- Was also used spf13/viper

### How to use:
While in the project directory:

`docker-compose up --build todo-app`

For SQL-migrations

`migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5435/postgres?sslmode=disable' up`

Use an HTTP client to test your application, for example: Postman

- - -