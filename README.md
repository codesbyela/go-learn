# go-learn

Go language learning project

## Structure

- [**api/**](./api)  
  Contains sample APIs:
  - [**users**](./api/users): Manage users using in-memory JSON.
  - [**books**](./api/books): Manage books using PostgreSQL.

- [**basics/**](./basics)  
  Contains simple Go programs and learning exercises.

- [**ordermanagement/**](./ordermanagement)  
  Main backend project for order and product management.

  - [**repository/postgres/**](./ordermanagement/internal/repository/postgres)  
    Repository layer for database operations using GORM and PostgreSQL.

  - [**services/**](./ordermanagement/internal/services)  
    Business logic and service layer (e.g., product service).

  - [**models/**](./ordermanagement/internal/models)  
    Go structs for database models (Product, Order, OrderProduct, etc.).

  - [**pkg/handlers/**](./ordermanagement/pkg/handlers)  
    HTTP handlers for products, orders, and payment webhook.

  - [**pkg/db/**](./ordermanagement/pkg/db)  
    Database initialization and migration code.  
    - [**migrations.sql**](./ordermanagement/pkg/db/migrations.sql): Raw SQL migration queries for tables.

  - [**main.go**](./ordermanagement/main.go)  
    Application entry point, router setup.

## Features

- Product CRUD operations
- Order CRUD operations
- Payment webhook for updating order payment status
- Pagination for product and order listing
- Database migrations (GORM + raw SQL)
- Sample APIs for users (JSON) and books (Postgres)

## TODO

- [x] Basics
- [x] Golang and Mux
- [x] Order & Product management
- [x] Payment webhook