# Gator CLI Application

## Overview
Gator is a command-line interface (CLI) application that provides user management, feed tracking, and social interaction features.

## Prerequisites
- Go (1.20 or later)
- PostgreSQL
- Git

## Configuration
The application uses a configuration file located at `~/.gatorconfig.json` with the following structure:
```json
{
    "db_url": "<db_url>"
}
```

## Installation

### 1. Clone the Repository
```bash
git clone <repository-url>
cd gator
```

### 2. Set Up PostgreSQL Database
Create a PostgreSQL database named `gator`:
```bash
createdb gator
```

### 3. Install Dependencies
```bash
go mod download
```

### 4. Build the Application
```bash
go build
```

## Available Commands

- `login`: User login
- `register`: User registration
- `reset`: Reset user account
- `users`: List users
- `agg`: Aggregate operations
- `addfeed`: Add a new feed (requires login)
- `feeds`: List feeds
- `follow`: Follow a user (requires login)
- `unfollow`: Unfollow a user (requires login)
- `following`: List followed users (requires login)
- `browse`: Browse content (requires login)


## Configuration Notes
- The database connection is configured in `~/.gatorconfig.json`
- `current_user_name` is automatically saved upon successful login

## Development
- Ensure PostgreSQL is running
- Set up your database connection in the config file
- Use `go run .` for development

