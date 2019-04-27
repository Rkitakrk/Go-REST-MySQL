# Go-REST-MySQL

## **_MySQL_**

we have to create our database with action table, so run file "db.sql"

---

## **_Go_**

You have to download my file "main.go"

We have to import packages "github.com/go-sql-driver/mysql" and "github.com/gorilla/mux", so use this command belows:

`go get github.com/go-sql-driver/mysql`

`go get github.com/gorilla/mux`

You have to set `userDB, passDB` to your database game in file `main.go`

Then we have to run file "main.go" with command:

`go run main.go`

---

## ****Send POST Request****

We can use `curl` or `Postman`

**CURL**

`curl -d "userId=2&gameId=3&action=23,32" -X POST http://localhost:8000/create`

**POSTMAN**

postman is a program which help send requests and is very helpful.

`userId int`

`gameId int`

`action string`

![postman](/image/postman.png)
Format: ![program postman](url)

---

## **_Router_**

/ - HomePage

/all - Show all records from action table in JSON

/create - save data to our database MySQL
