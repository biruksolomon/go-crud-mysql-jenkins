# Go CRUD API with MySQL & Jenkins CI/CD

A simple and clean RESTful CRUD API built using **Go (Golang)** and **MySQL**, designed as a portfolio project.  
The project also demonstrates a **basic CI/CD pipeline using Jenkins** and supports **Docker** for containerization.

---

## 🚀 Features

- RESTful CRUD operations (Create, Read, Update, Delete)
- Go (Golang) backend
- MySQL database integration
- Gorilla Mux router
- Dockerized application
- Basic CI/CD pipeline using Jenkins
- Clean and beginner-friendly project structure

---

## 🧱 Tech Stack

- **Backend:** Go (Golang)
- **Database:** MySQL
- **Router:** Gorilla Mux
- **CI/CD:** Jenkins
- **Containerization:** Docker
- **Version Control:** Git & GitHub

---

## 📂 Project Structure

```text
go-crud-mysql-jenkins/
├── main.go
├── config/
│   └── db.go
├── handlers/
│   └── item_handler.go
├── models/
│   └── item.go
├── Dockerfile
├── Jenkinsfile
├── go.mod
└── README.md


Required Envionment Variables
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=go_crud_db




API EndPoints

## Create Item
### POST /items
{
  "name": "Laptop",
  "price": 75000
}

## *Get All Items
### GET /items

## *Get Item By Id
### GET /items/{id}

## *Update Item 
### PUT /items/{id}
Request Body
{
  "name": "Gaming Laptop",
  "price": 90000
}

## *Delete Item 
### DELETE /items/{id}


Clone the Repository
git clone https://github.com/biruksolomon/go-crud-mysql-jenkins.git
cd go-crud-mysql-jenkins

Install Dependencies
go mod tidy

Run the Application
go run main.go

server will avialable at : http://localhost:8381

Build Docker image
docker build -t go-crud-api .

Run Docker Container
docker run -p 8080:8080 \
-e DB_USER=root \
-e DB_PASSWORD=root \
-e DB_HOST=host.docker.internal \
-e DB_PORT=3306 \
-e DB_NAME=go_crud_db \
go-crud-api

Start All service using docker compose using network
docker-compose up --build

Stop Services docker-compose down



