# Example ​​hexagonal architecture in Go

The idea of ​​hexagonal architecture is to separate:

- Domain (business rules)
- Application (use cases)
- Adapters (input: HTTP / output: DB)
- Infrastructure (real implementations)
- Interface (e.g. mocks, CLI, tests)

```bash
📁 api/
├── cmd/                  # Entry point of the application
│   └── server/           # Initializes the HTTP server
│       └── main.go
├── internal/
│   ├── domain/           # Domain entities + interfaces
│   │   └── item.go
│   ├── application/      # Use cases  
│   │   └── create_item.go
│   │   └── get_item.go
│   ├── ports/            # Input/output ports (interfaces)
│   │   ├── inbound/      # Input (ex: controller)
│   │   │   └── http.go
│   │   └── outbound/     # Output (ex: repository)
│   │       └── item_repository.go
│   ├── adapters/
│   │   ├── inbound/
│   │   │   └── http/     # HTTP Handler with Quick
│   │   │       └── handler.go
│   │   └── outbound/
│   │       └── postgres/ # Fake/mock DB implementation
│   │           └── repository.go
├── go.mod
├── go.sum
```

### 🧠 Hexagonal View

- Core with domain/ and application/ (business rules)
- Ports with ports/ defines the expected interfaces
- Adapters with adapters/ connect real inputs and outputs to the core
- Framework with cmd/server/main.go calls quick, defines routes and injects dependencies


🧠 what are we going to do:

- Both use cases are built on top of the ItemRepository interface
- Neither of them knows if the data comes from PostgreSQL, memory, Redis... 💡
- This keeps the domain and application logic infrastructure agnostic


### Outbound Adapter (Mock PostgreSQL)

```bash
📁 internal/adapters/outbound/postgres/repository.go
```

- A concrete implementation of ItemRepository
- Uses sync.RWMutex for safe concurrent access
- Can easily be replaced in the future by a real version with pgx, sqlx, etc.

###  Adapter Inbound (Quick Handler)
```bash
📁 internal/adapters/inbound/http/handler.go
```
- Handler orchestrates the use cases
- Does not know the direct repository
- Uses quick to respond with .JSON(...)


### 🧠 What we did

You now have a complete API with:

✅ Hexagonal Architecture
✅ Separation between domain, application and infrastructure
✅ Clean Ports and Adapters
✅ Testable with mocks
✅ Using quick for the HTTP layer
✅ Ready for real PostgreSQL plug-and-play later

### Final structure

```bash
api/
├── cmd/server/main.go
├── internal/
│   ├── adapters/
│   │   ├── inbound/http/handler.go
│   │   └── outbound/postgres/repository.go
│   ├── application/
│   │   ├── create_item.go
│   │   └── get_item.go
│   ├── domain/item.go
│   └── ports/
│       ├── inbound/http.go
│       └── outbound/item_repository.go
├── go.mod
```

## 🧪 Running Locally

### Requirements
- Go 1.21 or later
- Git

### Steps

```bash
git clone https://github.com/jeffotoni/go-hexagonal-ddd
cd go-hexagonal-ddd
go run cmd/server/main.go

   ██████╗ ██╗   ██╗██╗ ██████╗ ██╗ ██╗
  ██╔═══██╗██║   ██║██║██╔═══   ██║ ██╔╝
  ██║   ██║██║   ██║██║██║      █████╔╝
  ██║▄▄ ██║██║   ██║██║██║      ██╔═██╗
  ╚██████╔╝╚██████╔╝██║╚██████╔ ██║  ██╗
   ╚══▀▀═╝  ╚═════╝ ╚═╝ ╚═════╝ ╚═╝  ╚═╝

 Quick v0.0.1 🚀 Fast & Minimal Web Framework
─────────────────── ───────────────────────────────
 🌎 Host : http://127.0.0.1
 📌 Port : 8080
 🔀 Routes: 2
─────────────────── ───────────────────────────────
```

### 🧠 cURL POST

```bash
curl -X POST http://localhost:8080/items \
  -H 'Content-Type: application/json' \
  -d '{"name": "Monitor 4K", "value": 1999.99}'
```

```bash
{
   "ID":"58e89dc6-1bb8-43f5-ae42-36899014334d",
   "Name":"Monitor 4K",
   "Value":1999.99
}
```

### 🧠 cURL GET

```bash
curl http://localhost:8080/items/{id}
```

```bash
curl http://localhost:8080/items/58e89dc6-1bb8-43f5-ae42-36899014334d
```

```bash
{
   "ID":"58e89dc6-1bb8-43f5-ae42-36899014334d",
   "Name":"Monitor 4K",
   "Value":1999.99
}
```

### 🔭 Potential Improvements

    🐘 Real PostgreSQL integration (pgx, sqlc, sqlx)
    🧪 Unit tests with mocks and table-driven tests
    🐳 Dockerfile and docker-compose setup
    📊 Observability (metrics, logs, tracing with OpenTelemetry)
    🧠 AI/ML integration via microservices

⸻

🙌 Credits

Developed by @jeffotoni, passionate about Go, software architecture, and clean code principles.
Powered by the Quick web framework → github.com/jeffotoni/quick

⸻

💬 Feedback or Contributions?

Feel free to open an issue, start a discussion, or send a PR! Let’s build better APIs, the clean way 🚀