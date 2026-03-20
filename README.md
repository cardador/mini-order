# Mini-Order Service Documentation

This project is a production-ready microservice demonstration built in **Go**. It features a PostgreSQL backend, a non-blocking internal event bus, and graceful shutdown capabilities.

---

## 🛠 Project Structure

The codebase is organized into modular packages to ensure separation of concerns:

* **`api/`**: Contains HTTP handlers for creating (`POST`) and fetching (`GET`) orders.
* **`event-bus/`**: Manages an internal worker pool using Go channels for asynchronous processing.
* **`logger/`**: Implements a custom middleware for logging request metadata and duration.
* **`model/`**: Defines the core `Order` data structure used across the service.
* **`store/`**: Contains the `Repository` interface and its PostgreSQL implementation.
* **`bootstrap.sh`**: A shell script to automate Docker startup and table initialization.

---

## 🚥 Getting Started

### 1. Setup the Environment
Run the bootstrap script to launch the PostgreSQL container via Docker Compose and create the `orders` table:
\```bash
chmod +x bootstrap.sh
./bootstrap.sh
\```

### 2. Run the Application
Start the server (defaults to port `:8080`):
\```bash
go run main.go
\```

### 3. Run Tests
The project uses table-driven unit tests with a `MockStore` to verify handler logic without requiring a live database:
\```bash
go test -v ./api
\```

---

## 📡 API Endpoints

### Create Order
* **URL**: `/order`
* **Method**: `POST`
* **Payload**: `{"id": "123", "item": "Laptop", "amount": 1500.00}`
* **Success Response**: `202 Accepted` (The order is saved to DB and sent to the event bus).

### Get Order
* **URL**: `/order/get/{id}`
* **Method**: `GET`
* **Success Response**: `200 OK` with the Order JSON object.

---

## 💡 Technical Highlights

### Robust Error Handling
The service handles PostgreSQL-specific errors. For example, it detects "Unique Constraint Violations" (Error `23505`) to inform the client if an Order ID already exists:
\```go
if pgErr, ok := err.(*pq.Error); ok {
    if pgErr.Code == "23505" {
        http.Error(w, "Order ID already exists", http.StatusConflict)
        return
    }
}
\```

### Non-Blocking Concurrency
The event bus uses a `select` statement with a `default` case. This prevents the API from hanging if the background processing buffer (capacity 100) is full:
\```go
func AddOrder(o model.Order) bool {
    select {
    case eventbus <- o:
        return true
    default:
        return false // Drops order if bus is full
    }
}
\```

### Clean Architecture
By defining a `store.Repository` interface, the API handlers are not tied to PostgreSQL. They can be tested easily using a `MockStore`:
\```go
type Repository interface {
    SaveOrder(ctx context.Context, o model.Order) error
    GetOrder(ctx context.Context, id string) (model.Order, error)
}
\```

---

## 🛡 License
Copyright (c) 2026 Marcus Cardador. Licensed under the **MIT License**.
