# Go Project

A Go project implementing clean architecture principles for user management with database integration.

## Project Structure

```
go-project/
├── build/
├── cmd/
│   ├── config.yaml
│   └── main.go
├── internal/
│   ├── adapter/
│   │   ├── config/
│   │   ├── errs/
│   │   ├── handler/
│   │   ├── logger/
│   │   └── storage/
│   └── core/
│       ├── application/
│       ├── domain/
│       ├── port/
│       └── service/
└── logs/
```

## Features

- Clean architecture implementation
- User management functionality
- Multiple database support (PostgreSQL, Redis)
- Configurable logging system
- HTTP request handling
- Error management system

## Installation

1. Clone the repository
```bash
git clone https://github.com/ximecloud/go-project
```

2. Install dependencies
```bash
go mod download
```

## Configuration

Configure the application using `cmd/config.yaml`. Example configuration:

```yaml
app:
  name: project
  port: 8080

database:
  host: localhost
  port: 5432
  username: ""
  password: ""
  database: ""
cache:
  host: localhost
  port: 6379
  database: 0
logstash:
  host: localhost
  port: 5043
```

## Usage

Run the application:
```bash
go run cmd/main.go
```

## Project Layout

- `cmd/`: Application entry points
- `internal/`: Private application code
    - `adapter/`: External interfaces and adapters
    - `core/`: Core business logic
        - `application/`: Application services
        - `domain/`: Domain models and business logic
        - `port/`: Interface definitions
        - `service/`: Implementation of core services
- `build/`: Compilation and packaging
- `logs/`: Application logs

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

MIT License
Copyright (c) 2024 XimeCloud