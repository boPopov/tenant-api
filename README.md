# tenant-api

## Introduction
`tenant-api` is a RESTful API built in Go for managing tenant information. It uses Docker for containerization, Swagger for API documentation, GitHub OAuth for authenticating users.

---

## Build and Run the API

### **1 Dev Environment**
#### Swagger API Documentation Setup
- **Install `swag`**: [Swag Documentation](https://github.com/swaggo/swag)
- Navigate to the `./api` directory.
- Remove the existing `swaggerdocs/` folder (if any).
- Generate Swagger documentation:
```bash
cd api
swag init --dir . --output ./swaggerdocs
```

#### Docker Commands (Development):
```bash
# Build Docker image
docker build -f Dockerfile.dev -t tenantapi:dev .

# Build API service with Docker Compose
docker-compose -f docker-compose.dev.yml build api

# Start containers
docker-compose -f docker-compose.dev.yml up -d

# Stop containers
docker-compose -f docker-compose.dev.yml down
```

### **2 Production Environment**
#### Docker Commands (Production):
```bash
# Build Docker image
docker build -t tenantapi:prod .

# Start containers
docker-compose up -d

# Stop containers
docker-compose down
```

#### Linux-Specific Commands:
```bash
# Remove all database data (use with caution!)
rm -rf db-data
```

---

## API Documentation (Swagger)
#### View Swagger Docs:
- Once the API is running, access Swagger UI at:
  - **URL:** `http://localhost:3000/swagger/index.html`

#### Regenerate Swagger Docs:
```bash
cd api
swag init --dir . --output ./swaggerdocs
```

---

## 🧪 Testing
### 🧪 Unit Tests:
```bash
go test ./... -v
```

### Integration Tests:
```bash
go test -v ./tests/integration
```

---

## Folder Structure
```
tenant-api/
├── api/                    # Main API code and handlers
├── tests/                  # Unit and integration tests
│   ├── integration         # Integration tests
│   ├── mock                # Mock functions
│   └── unit                # Unit tests
├── env/                    # Environment variable files
│   ├── api/.env            # API environment variables
│   └── db/.env             # Database environment variables
├── Dockerfile              # Production Dockerfile
├── docker-compose.yml      # Docker Compose for Production
├── Dockerfile.dev          # Development Dockerfile
├── docker-compose.dev.yml  # Docker Compose for Development
└── README.md               # Project Documentation
```

---

## Security
- OAuth 2.0 authentication is integrated with GitHub.

---

## Author
- Created by **Bojan Popov**

## License
- This project is licensed under the **MIT License**.
