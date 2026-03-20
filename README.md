# Mini-Order Service

[cite_start]A production-ready microservice demonstration built in **Go** featuring a PostgreSQL backend, a non-blocking internal event bus, and graceful shutdown capabilities[cite: 84, 111, 46].

---

## 🚀 Features

* [cite_start]**RESTful API**: Endpoints for creating and retrieving orders using Go 1.22+ routing.
* [cite_start]**PostgreSQL Persistence**: Robust data storage with specific handling for unique constraint violations (duplicate IDs)[cite: 29, 111].
* [cite_start]**Asynchronous Processing**: An internal event bus using Go channels to simulate heavy background work without blocking API responses[cite: 46, 47].
* [cite_start]**Dependency Injection**: Decoupled architecture using a `Repository` interface, making the business logic easily testable[cite: 118, 160].
* [cite_start]**Middleware Logging**: Request logging that tracks method, path, and execution duration[cite: 77, 78].
* [cite_start]**Graceful Shutdown**: Properly closes the server and database connections upon receiving interrupt signals[cite: 85, 86].

---

## 🛠 Project Structure

```text
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
