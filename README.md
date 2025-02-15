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

#### Linux Pre-Setup Commands
If running on **Linux**, ensure the following pre-setup steps are done:
```bash
# Create directory for pgAdmin persistent storage
mkdir -p ./pgadmin-data

# Set correct ownership for pgAdmin
sudo chown -R 5050:5050 ./pgadmin-data
sudo chmod -R 770 ./pgadmin-data
```
These steps ensure **pgAdmin** has the required permissions when running inside Docker.

Use the Environment variables inside the `./env/api/.env` file or inside the `docker-compose.dev.yml` file to create the server.

#### Linux-Specific Commands:
```bash
# Remove all database data (use with caution!)
sudo rm -rf db-data
sudo rm -rf pgadmin-data
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

## Testing
### Unit Tests:
```bash
# Navigate to the unit test folder
cd tests/unit_tests

# Simple run of the tests
go test . 

# Simple run of the tests without cache
go test --count=1 .

# Verbose execution of the tests
go test -v .

# Verbose execution of the tests without cache
go test --count=1 -v .
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

## Running the API Locally
For development, you can run the API directly on your local machine without using Docker:

### **Linux/macOS**:
```bash
chmod +x local_run_linux.sh
./local_run_linux.sh
```

### **Windows (PowerShell)**:
```cmd
.\local_run_windows.bat
```

These scripts:
- Load environment variables from `env/api/.env`
- Start the Go application manually

Ensure that **PostgreSQL** is running before starting the API locally.

**Note:** Additionally, change the environment variables based on your environment.
---

## Security
- OAuth 2.0 authentication is integrated with GitHub.

---

## Environment Variables
### API Environment Variables
| Variable | Description |
|----------|------------|
| `PORT` | Port where the API runs (Default: `3000`) |
| `DB_HOST` | Database Host for Accessing the database |
| `DB_USER` | Database User for Accessing the database |
| `DB_PASSWORD` | Database Password for Accessing the database |
| `DB_NAME` | Database Name for Accessing the database |
| `DB_PORT` | Port where the Database can be accessed (Default: `5432`) |
| `JWT_SECRET` | Secret key for signing JWT tokens |
| `JWT_EXPIRE_INTERVAL` | Interval until the JWT token expires (Example: 55s, 15m, 2h, 1d) |
| `DATABASE_URL` | Connection string for PostgreSQL |
| `GITHUB_CLIENT_ID` | GitHub OAuth Client ID |
| `GITHUB_CLIENT_SECRET` | GitHub OAuth Client Secret |

### Database Environment Variables
| Variable | Description |
|----------|------------|
| `POSTGRES_USER` | PostgreSQL username (Default: `postgres`) |
| `POSTGRES_PASSWORD` | PostgreSQL password |
| `POSTGRES_DB` | Database name (Default: `tenants`) |

---


## Author
- Created by **Bojan Popov**

## License
- This project is licensed under the **MIT License**.
