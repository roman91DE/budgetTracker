
# BudgetTracker

BudgetTracker is a web application for managing personal finances, allowing users to track expenses and income efficiently. It's built with Go, using the Gin framework for the backend and SQLite for data storage.

## Features

- User Authentication
- Expense and Income Tracking
- Budget Planning
- Financial Reports

## Quick Start

### Prerequisites

- Go (1.16+)
- SQLite3

### Setup

Clone and navigate into the project:

```sh
git clone https://github.com/yourusername/budgetTracker.git
cd budgetTracker
```

Install dependencies and initialize the database:

```sh
go mod tidy
go run cmd/budgetTracker/main.go initdb
```

Run the server:

```sh
go run cmd/budgetTracker/main.go
```

Access the app at `http://localhost:8080`.

## Testing

Run automated tests with:

```sh
go test ./...
```

## Deployment

Refer to Go and SQLite documentation for deployment guidelines.

## Built With

- [Go](https://golang.org/)
- [Gin Framework](https://github.com/gin-gonic/gin)
- [SQLite](https://sqlite.org/index.html)

## Contributing

Contributions are welcome. Please open an issue or pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.
