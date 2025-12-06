# Backend Go - User Management API

Backend API sederhana menggunakan Go, Gin Framework, dan GORM untuk manajemen data user dengan database MySQL.

## 🚀 Tech Stack

- **Go** v1.23.0+
- **Gin** - Web Framework
- **GORM** - ORM untuk database
- **MySQL** - Database
- **godotenv** - Environment variable management

## 📁 Struktur Folder

```
backend-go/
├── app.go                 # Entry point aplikasi
├── go.mod                 # Go modules
├── .env                   # Environment variables (buat sendiri)
├── Config/
│   └── database.go        # Konfigurasi koneksi database
├── Controller/
│   └── main.controller.go # Handler untuk routes
├── Model/
│   └── user.model.go      # Model User
└── Routes/
    └── routes.go          # Definisi routes
```

## ⚙️ Setup & Installation

### 1. Clone/Download Project

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup Environment Variables

Buat file `.env` di root folder dengan isi:

```env
APP_PORT=:8000
DB_USER=root
DB_PASS=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=your_database_name
```

### 4. Setup Database

Buat database MySQL dan tabel user:

```sql
CREATE DATABASE your_database_name;

USE your_database_name;

CREATE TABLE user (
    id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255),
    name VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 5. Run Application

```bash
go run app.go
```

Server akan berjalan di `http://localhost:8000`

## 📌 API Endpoints

### 1. Test Endpoint

```http
GET /
```

**Response:**

```
Bismillah test
```

### 2. Get All Users

```http
GET /bismillah-tes
```

**Response:**

```json
[
  {
    "id": "1",
    "email": "user@example.com",
    "name": "John Doe",
    "password": "hashed_password",
    "createdAt": "2025-12-06T10:00:00Z",
    "updatedAt": "2025-12-06T10:00:00Z"
  }
]
```

### 3. Create User

```http
GET /bismillah-create
```

**Request Body:**

```json
[
  {
    "id": "1",
    "email": "user@example.com",
    "name": "John Doe",
    "password": "password123"
  }
]
```

## 📝 Notes

- Password disimpan dalam plain text (untuk production sebaiknya gunakan hashing seperti bcrypt)
- CreateUser endpoint saat ini menggunakan method GET, sebaiknya ubah ke POST
- Belum ada validation dan error handling yang lengkap

## 🔧 Development

Untuk development, gunakan:

```bash
go run app.go
```

Untuk build production:

```bash
go build -o backend-go.exe app.go
```

## 📄 License

MIT
