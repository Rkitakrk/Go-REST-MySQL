# Go-REST-MySQL

**_MySQL_**
CREATE DATABASE game;

USE game;

CREATE TABLE action (

Id int NOT NULL AUTO_INCREMENT PRIMARY KEY,

userId int NOT NULL,

gameId int NOT NULL,

action TEXT NOT NULL

);

---

**_POSTMAN_**

Sent request POST

userId int

gameId int

action string

I use Postman to send request

---

**_Router_**

/ - HomePage

/all - Show all requests

/create - save to database MySQL

---

**_GO-Command_**

go get github.com/go-sql-driver/mysql

go get github.com/gorilla/mux

go run main.go
