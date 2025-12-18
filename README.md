# Guvi Project (Go + React)

## Prerequisites
Ensure you have the following installed and running:
- **Go** (Golang) v1.21+
- **Node.js** & **npm**
- **MySQL** (Database: `guvi_db`, Table: `users`)
- **Redis** (Port: 6379, default)
- **MongoDB** (Port: 27017, default)

## Setup & Run

### 1. Database Setup
Execute the following SQL in your MySQL instance:
```sql
CREATE DATABASE IF NOT EXISTS guvi_db;
USE guvi_db;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 2. Backend (Go)
Open a terminal in the project root:
```bash
cd server
go mod tidy
go run main.go
```
The server will start on `http://localhost:8080`.

### 3. Frontend (React)
Open a **new** terminal in the project root:
```bash
cd client
npm install   # If not already installed
npm run dev
```
The frontend will start on `http://localhost:5173`.

## Features
- **Register**: Validates password strength, stores in MySQL (hashed).
- **Login**: Verifies credentials, generates session token, stores in Redis, saves in localStorage.
- **Profile**: Protected route, fetches additional details from MongoDB, allows updates.
- **Theme**: Dark/Light mode toggle persisted in localStorage.
