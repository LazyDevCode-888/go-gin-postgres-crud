# Go Gin PostgreSQL CRUD API

RESTful API สำหรับจัดการข้อมูลสินค้า (Products) พัฒนาด้วย Go, Gin, GORM และ PostgreSQL

## Tech Stack

* Go
* Gin Framework
* GORM
* PostgreSQL
* Godotenv

---

## Features

* Create Product
* Get All Products
* Get Product By ID
* Update Product
* Delete Product

---

## Project Structure

```text
go-gin-postgres-crud/

├── cmd/
│   └── main.go
│
├── config/
│   └── database.go
│
├── controllers/
│   └── product_controller.go
│
├── models/
│   └── product.go
│
├── routes/
│   └── routes.go
│
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

## Database Schema

### Product

| Field       | Type    |
| ----------- | ------- |
| ID          | uint    |
| Name        | string  |
| Description | string  |
| Price       | float64 |
| Stock       | int     |

---

## Installation

### Clone Repository

```bash
git clone https://github.com/LazyDevCode-888/go-gin-postgres-crud.git

cd go-gin-postgres-crud
```

### Install Dependencies

```bash
go mod tidy
```

---

## Environment Variables

สร้างไฟล์ `.env`

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=123456
DB_NAME=crud_db
DB_SSLMODE=disable
```

หรือคัดลอกจาก

```bash
cp .env.example .env
```

---

## Create Database

เข้า PostgreSQL

```sql
CREATE DATABASE crud_db;
```

---

## Run Application

```bash
go run cmd/main.go
```

Server

```text
http://localhost:8080
```

---

## API Endpoints

### Create Product

POST

```http
/api/products
```

Request

```json
{
  "name": "Gaming Mouse",
  "description": "RGB Mouse",
  "price": 799,
  "stock": 20
}
```

---

### Get All Products

GET

```http
/api/products
```

---

### Get Product By ID

GET

```http
/api/products/1
```

---

### Update Product

PUT

```http
/api/products/1
```

Request

```json
{
  "name": "Gaming Mouse RGB",
  "description": "RGB Wireless Mouse",
  "price": 999,
  "stock": 50
}
```

---

### Delete Product

DELETE

```http
/api/products/1
```

---

## Dependencies

```bash
github.com/gin-gonic/gin
gorm.io/gorm
gorm.io/driver/postgres
github.com/joho/godotenv
```

---

## Author

Arthit Khabuankeo

GitHub: https://github.com/LazyDevCode-888
