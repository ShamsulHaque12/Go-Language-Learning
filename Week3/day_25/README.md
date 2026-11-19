# Product Management REST API (Clean Architecture)

A production-ready Product Management RESTful API written in Go, adhering strictly to **Clean Architecture** principles and the **Standard Go Project Layout**. Supports complete **GET**, **PUT**, **POST**, **PATCH**, and **DELETE** HTTP operations.

---

## 🏛️ Architecture Overview

The system is decoupled into isolated layers where dependencies flow strictly inward:

```
[ Client / HTTP Request ]
            ↓
[ Presentation Layer / Handler & Middleware ]
  - Routes (GET, POST, PUT, PATCH, DELETE)
  - Gin Engine, CORS, Request Logger, Recovery, JWT Auth
            ↓
[ Business Logic Layer / Service ]
  - Product Service, Auth Service, User Service
  - Business rules, validation, DTO transformation
            ↓
[ Data Access Layer / Repository ]
  - Product Repository, User Repository
  - PostgreSQL, GORM, Transactions, Row-Locking
            ↓
[ Database / PostgreSQL ]
```

---

## 📁 Directory Structure

```
Week3/day_25/
├── cmd/
│   └── api/
│       └── main.go                 # App entrypoint, dependency injection & graceful shutdown
├── config/
│   └── config.go                   # Environment-driven configuration loader (.env)
├── internal/
│   ├── domain/                     # Entities, DTOs & Domain errors
│   │   ├── product.go              # Product model & DTOs (Create, Update, Patch, Filter)
│   │   ├── user.go                 # User model & DTOs
│   │   └── errors.go               # Custom domain errors
│   ├── repository/                 # Data access layer (PostgreSQL with GORM)
│   │   ├── product_repository.go   # Product CRUD, pagination, filtering, patching
│   │   └── user_repository.go      # User CRUD & transactional balance transfer
│   ├── service/                    # Business logic layer
│   │   ├── product_service.go      # Product business logic
│   │   ├── product_service_test.go # Product unit tests (GET, POST, PUT, PATCH, DELETE)
│   │   ├── auth_service.go
│   │   ├── user_service.go
│   │   └── mock_product_repository_test.go
│   ├── handler/                    # HTTP presentation layer (Gin)
│   │   ├── product_handler.go      # Handlers for GET, POST, PUT, PATCH, DELETE
│   │   ├── auth_handler.go
│   │   ├── user_handler.go
│   │   ├── response.go             # Standardized API response builder
│   │   └── routes.go               # Gin routing engine setup
│   └── middleware/                 # HTTP middlewares
│       ├── auth_middleware.go      # JWT Bearer token validation
│       ├── cors_middleware.go      # CORS headers
│       └── logger_middleware.go    # Latency & HTTP request logging
├── pkg/
│   ├── database/
│   │   └── postgres.go             # PostgreSQL connection pooling with GORM
│   ├── token/                      # JWT generation and validation
│   └── utils/                      # Password hashing with bcrypt
├── migrations/
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   ├── 000002_create_products_table.up.sql
│   └── 000002_create_products_table.down.sql
├── .env.example
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

## ⚡ Product Management API Specification

| Method | Endpoint | Description | Payload / Params |
| :--- | :--- | :--- | :--- |
| **POST** | `/api/v1/products` | Create a new product | Complete product JSON |
| **GET** | `/api/v1/products` | Get list of products with filters | Query params (`page`, `page_size`, `category`, `status`, `search`, `sort_by`, `order`) |
| **GET** | `/api/v1/products/:id` | Get single product by ID | Path variable `:id` |
| **PUT** | `/api/v1/products/:id` | Full replacement of existing product | All product fields required |
| **PATCH** | `/api/v1/products/:id` | Partial update of existing product | Only fields to be updated |
| **DELETE** | `/api/v1/products/:id` | Delete product by ID | Path variable `:id` |

---

## 📡 API Request & Response Examples

### 1. POST: Create Product
- **URL**: `POST http://localhost:8080/api/v1/products`
- **Request Body**:
```json
{
  "name": "Apple MacBook Pro 16",
  "description": "M3 Max, 36GB RAM, 1TB SSD Space Black",
  "sku": "MACBOOK-PRO-16-M3",
  "price": 3499.99,
  "stock": 15,
  "category": "Laptops",
  "status": "active"
}
```
- **cURL**:
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Apple MacBook Pro 16",
    "description": "M3 Max, 36GB RAM, 1TB SSD Space Black",
    "sku": "MACBOOK-PRO-16-M3",
    "price": 3499.99,
    "stock": 15,
    "category": "Laptops",
    "status": "active"
  }'
```
- **Response** (`201 Created`):
```json
{
  "success": true,
  "message": "Product created successfully",
  "data": {
    "id": 1,
    "name": "Apple MacBook Pro 16",
    "description": "M3 Max, 36GB RAM, 1TB SSD Space Black",
    "sku": "MACBOOK-PRO-16-M3",
    "price": 3499.99,
    "stock": 15,
    "category": "Laptops",
    "status": "active",
    "created_at": "2026-11-19T09:55:00Z",
    "updated_at": "2026-11-19T09:55:00Z"
  }
}
```

---

### 2. GET: Get All Products (With Filters & Pagination)
- **URL**: `GET http://localhost:8080/api/v1/products?category=Laptops&search=MacBook&page=1&page_size=10&sort_by=price&order=desc`
- **cURL**:
```bash
curl -X GET "http://localhost:8080/api/v1/products?category=Laptops&page=1&page_size=10"
```
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "Products retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Apple MacBook Pro 16",
      "description": "M3 Max, 36GB RAM, 1TB SSD Space Black",
      "sku": "MACBOOK-PRO-16-M3",
      "price": 3499.99,
      "stock": 15,
      "category": "Laptops",
      "status": "active",
      "created_at": "2026-11-19T09:55:00Z",
      "updated_at": "2026-11-19T09:55:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "page_size": 10,
    "total_items": 1,
    "total_pages": 1
  }
}
```

---

### 3. GET: Get Product by ID
- **URL**: `GET http://localhost:8080/api/v1/products/1`
- **cURL**:
```bash
curl -X GET http://localhost:8080/api/v1/products/1
```
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "Product retrieved successfully",
  "data": {
    "id": 1,
    "name": "Apple MacBook Pro 16",
    "description": "M3 Max, 36GB RAM, 1TB SSD Space Black",
    "sku": "MACBOOK-PRO-16-M3",
    "price": 3499.99,
    "stock": 15,
    "category": "Laptops",
    "status": "active",
    "created_at": "2026-11-19T09:55:00Z",
    "updated_at": "2026-11-19T09:55:00Z"
  }
}
```

---

### 4. PUT: Full Update (Complete Replacement)
Replaces all product attributes. Every required field must be supplied.

- **URL**: `PUT http://localhost:8080/api/v1/products/1`
- **Request Body**:
```json
{
  "name": "Apple MacBook Pro 16 (2026 Edition)",
  "description": "Updated description with M4 Max processor",
  "sku": "MACBOOK-PRO-16-M4",
  "price": 3799.99,
  "stock": 25,
  "category": "Laptops",
  "status": "active"
}
```
- **cURL**:
```bash
curl -X PUT http://localhost:8080/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Apple MacBook Pro 16 (2026 Edition)",
    "description": "Updated description with M4 Max processor",
    "sku": "MACBOOK-PRO-16-M4",
    "price": 3799.99,
    "stock": 25,
    "category": "Laptops",
    "status": "active"
  }'
```
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "Product updated completely (PUT)",
  "data": {
    "id": 1,
    "name": "Apple MacBook Pro 16 (2026 Edition)",
    "description": "Updated description with M4 Max processor",
    "sku": "MACBOOK-PRO-16-M4",
    "price": 3799.99,
    "stock": 25,
    "category": "Laptops",
    "status": "active",
    "created_at": "2026-11-19T09:55:00Z",
    "updated_at": "2026-11-19T10:00:00Z"
  }
}
```

---

### 5. PATCH: Partial Update (Modify specific fields only)
Updates only the specified non-nil fields without touching any other field.

- **URL**: `PATCH http://localhost:8080/api/v1/products/1`
- **Request Body** (e.g. updating price and stock only):
```json
{
  "price": 3299.99,
  "stock": 10
}
```
- **cURL**:
```bash
curl -X PATCH http://localhost:8080/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{"price": 3299.99, "stock": 10}'
```
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "Product partially updated (PATCH)",
  "data": {
    "id": 1,
    "name": "Apple MacBook Pro 16 (2026 Edition)",
    "description": "Updated description with M4 Max processor",
    "sku": "MACBOOK-PRO-16-M4",
    "price": 3299.99,
    "stock": 10,
    "category": "Laptops",
    "status": "active",
    "created_at": "2026-11-19T09:55:00Z",
    "updated_at": "2026-11-19T10:02:00Z"
  }
}
```

---

### 6. DELETE: Delete Product
- **URL**: `DELETE http://localhost:8080/api/v1/products/1`
- **cURL**:
```bash
curl -X DELETE http://localhost:8080/api/v1/products/1
```
- **Response** (`200 OK`):
```json
{
  "success": true,
  "message": "Product deleted successfully (DELETE)",
  "data": {
    "id": 1
  }
}
```

---

## 🧪 Running Unit Tests

```bash
cd /media/Data/go_language/Week3/day_25
go test -v ./...
```

## 🏃 Running the Application

```bash
cd /media/Data/go_language/Week3/day_25
go run cmd/api/main.go
```
