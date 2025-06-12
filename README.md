# Hospital Management System

A Golang-based web application for managing hospital operations with separate portals for receptionists and doctors.

## Features

- Single login system for both receptionists and doctors
- Patient registration and management (CRUD operations)
- Doctor's portal for viewing and updating patient information
- JWT-based authentication
- PostgreSQL database with GORM ORM
- RESTful API architecture
- Swagger documentation

## Project Structure

```
hospital-management/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   ├── controllers/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   └── services/
├── pkg/
│   ├── database/
│   └── utils/
├── docs/
├── migrations/
├── tests/
├── .env
├── .gitignore
├── go.mod
└── README.md
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL
- Make (optional, for using Makefile commands)

## Setup and Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd hospital-management
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=hospital_db
JWT_SECRET=your_jwt_secret
```

4. Run database migrations:
```bash
go run migrations/migrate.go
```

5. Start the server:
```bash
go run cmd/main.go
```

## API Documentation

The API documentation is available at `/swagger/index.html` when the server is running.

## Testing

Run the tests using:
```bash
go test ./...
```

## License

MIT License 