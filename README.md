# Who Is The Undercover - 谁是卧底

A backend service for the popular party game "Who's the Undercover" (谁是卧底), implemented in Go.

## Overview

This project provides a web service that manages word pairs used in the "Who's the Undercover" game. Players receive similar but different words, and they need to describe their words without revealing them directly. The goal is to find out who has the different word (the "undercover") through discussion.

## Features

- RESTful API for managing game word pairs
- Random word pair selection for gameplay
- Clean architecture design following best practices
- Dependency injection with Wire

## Tech Stack

- Go
- Gin Web Framework
- XORM for database operations
- Wire for dependency injection

## Project Structure

```
.
├── cmd/              # Application entry point
├── configs/          # Configuration files
├── docs/             # Documentation and SQL scripts
├── http/             # HTTP request examples
└── internal/         # Internal application code
    ├── base/         # Base infrastructure
    ├── controller/   # HTTP controllers
    ├── entity/       # Domain entities
    ├── repo/         # Data repositories
    ├── router/       # HTTP routing
    └── service/      # Business logic services
```

## Getting Started

### Prerequisites

- Go 1.18 or higher
- MySQL database

### Installation

1. Clone the repository:

```bash
git clone https://github.com/llmons/who-is-the-undercover.git
cd who-is-the-undercover
```

2. Install dependencies:

```bash
go mod download
```

3. Set up the database:

```bash
# Execute the SQL script in docs/init.sql to set up the database
```

4. Update the configuration in `configs/config.yaml` with your database credentials

### Running the application

```bash
wire ./cmd
go run cmd
```

The server will start at `localhost:8080`.

## API Endpoints

- `GET /api/wordpairs`: Get all word pairs
- `GET /api/wordpairs/random`: Get a random word pair for a game

## Development

### Building the binary

```bash
go build -o ./tmp/main ./cmd
```

### Regenerating Wire dependency injection

```bash
wire ./cmd
```