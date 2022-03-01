# Go REST API with Fiber and GORM

This repository is a simple project for ordering system API, fully dedicated for me to learn Go, Fiber framework and Go Object Relational Mapping (GORM) that can be connected with SQL database (In this project I used SQLite).

To test this API locally, you need to have [Go](https://go.dev) installed on your local machine. Then, kindly download the zip file of the source code. After extract the folder, open the terminal in that extracted folder and type

```Bash
go get github.com/gofiber/fiber/v2
```

This will install all the packages for Fiber framework

```Bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

This will install all the packages for GORM

To run this API, just type

```Bash
go run main.go
```

Then, the server will be listening on `http://localhost:3005`.

or, if you want to use Air packages for developing/modifying this API, then type

```Bash
go install github.com/cosmtrek/air@latest
```

Then, run this API by typing

```Bash
air
```

Then, the server will be listening on `http://localhost:3005`.

## Available API routes

`GET` : `http://localhost:3005`

`GET` : `http://localhost:3005/api/user`  
`GET` : `http://localhost:3005/api/user/:id`  
`POST` : `http://localhost:3005/api/user`  
`PATCH` : `http://localhost:3005/api/user/:id`  
`DELETE` : `http://localhost:3005/api/user/:id`

`GET` : `http://localhost:3005/api/product`  
`GET` : `http://localhost:3005/api/product/:id`  
`POST` : `http://localhost:3005/api/product`  
`PATCH` : `http://localhost:3005/api/product/:id`  
`DELETE` : `http://localhost:3005/api/product/:id`

`GET` : `http://localhost:3005/api/order`  
`GET` : `http://localhost:3005/api/order/:id`  
`POST` : `http://localhost:3005/api/order`

## References

-   https://go.dev
-   https://gofiber.io
-   https://github.com/gofiber/fiber
-   https://gorm.io
-   https://github.com/go-gorm/gorm
-   https://github.com/cosmtrek/air
