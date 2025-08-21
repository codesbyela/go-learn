# go-learn

Go language learning project

## Structure

- **ordermanagement/**  
  Main backend project for order and product management.

  - **internal/models/**  
    Contains Go structs for database models (Product, Order, OrderProduct, etc.).

  - **internal/repository/postgres/**  
    Repository layer for database operations using GORM and PostgreSQL.

  - **internal/services/**  
    Business logic and service layer (e.g., product service).

  - **pkg/handlers/**  
    HTTP handlers for products, orders, and payment webhook.

  - **pkg/db/**  
    Database initialization and migration code.  
    - `migrations.sql`: Raw SQL migration queries for tables.

  - **main.go**  
    Application entry point, router setup.

## Features

- Product CRUD operations
- Order CRUD operations
- Payment webhook for updating order payment status
- Pagination for product and order listing
- Database migrations (GORM + raw SQL)

## TODO

- [x] Basics
- [x] Golang and Mux
- [x] Order & Product management
- [x] Payment webhook