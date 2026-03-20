Mini-Order Service

A production-ready microservice demonstration built in Go featuring a PostgreSQL backend, a non-blocking internal event bus, and graceful shutdown capabilities.
🚀 Features

    RESTful API: Endpoints for creating and retrieving orders using Go 1.22+ routing.

    PostgreSQL Persistence: Robust data storage with specific handling for unique constraint violations (duplicate IDs).

    Asynchronous Processing: An internal event bus using Go channels to simulate heavy background work without blocking API responses.

    Dependency Injection: Decoupled architecture using a Repository interface, making the business logic easily testable.

    Middleware Logging: Request logging that tracks method, path, and execution duration.

    Graceful Shutdown: Properly closes the server and database connections upon receiving interrupt signals.

🛠 Project Structure
Plaintext

.
├── api/
│   ├── handler.go       # HTTP Handlers (POST /order, GET /order/get/{id})
│   └── handler_test.go  # Table-driven unit tests with MockStore
├── event-bus/
│   └── event_bus.go     # Internal worker pool and channel logic
├── logger/
│   └── logger.go        # HTTP middleware for request logging
├── model/
│   └── model.go         # Shared Order data structures
├── store/
│   ├── postgres.go      # PostgreSQL implementation of Repository
│   └── store.go         # Repository interface definition
├── bootstrap.sh         # Automated Docker and Database setup script
├── docker-compose.yml   # PostgreSQL container configuration
└── main.go              # Application entry point and wiring

🚥 Getting Started
Prerequisites

    Go 1.22 or higher 

    Docker and Docker Compose 

1. Setup the Environment

Use the provided bootstrap script to launch the PostgreSQL container and initialize the database schema automatically:
Bash

chmod +x bootstrap.sh
./bootstrap.sh

The script waits for the database to be healthy before creating the orders table.
2. Run the Application
Bash

go run main.go

The server will start on :8080.
3. Run Tests

Execute the test suite, including checks for race conditions:
Bash

go test -v -race ./...

📡 API Endpoints
Create Order

    URL: /order

    Method: POST

    Payload: {"id": "123", "item": "Laptop", "amount": 1500.00} 

    Success Response: 202 Accepted (Order sent to background processing).

Get Order

    URL: /order/get/{id}

    Method: GET

    Success Response: 200 OK with Order JSON.

🛡 License

This project is licensed under the MIT License.
