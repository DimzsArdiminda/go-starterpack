# Backend Go Starter Kit

Starter kit API Go dengan struktur yang terinspirasi dari Laravel: konfigurasi database terpisah, model, controller, dan satu tempat untuk mendaftarkan route API.

## Tech stack

- Go 1.23+
- Gin
- GORM
- PostgreSQL
- godotenv

## Struktur folder

```text
backend-go/
├── app.go                 # Entry point aplikasi
├── Config/                # config/database.go versi Go
├── Controller/            # app/Http/Controllers versi Go
├── Model/                 # app/Models versi Go
├── Routes/                # routes/api.go versi Go
├── .env.example           # Template konfigurasi lokal
└── go.mod
```

## Instalasi sampai berjalan

### 1. Siapkan prasyarat

Pastikan Go dan PostgreSQL sudah terpasang, lalu cek versinya:

```bash
go version
psql --version
```

Gunakan Go 1.23 atau yang lebih baru. PostgreSQL harus sedang berjalan dan perintah `psql` perlu tersedia di terminal. Jika belum ada PostgreSQL, instal dari [postgresql.org](https://www.postgresql.org/download/).

### 2. Masuk ke folder proyek dan unduh dependensi

```bash
cd backend-go
go mod download / go mod tidy
```

### 3. Buat konfigurasi environment

Salin template `.env.example` menjadi `.env`.

PowerShell:

```powershell
Copy-Item .env.example .env
```

macOS/Linux/Git Bash:

```bash
cp .env.example .env
```

Isi `.env` sesuai akun PostgreSQL lokal Anda:

```env
APP_PORT=:8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=secret
DB_NAME=backend_go
```

`APP_PORT` harus memakai awalan titik dua, misalnya `:8080`. Jangan commit file `.env`, karena berisi kredensial lokal.

### 4. Buat database dan tabel

Login ke PostgreSQL dengan user yang sama seperti `DB_USER`:

```bash
psql -U postgres -h localhost
```

Jalankan SQL berikut di prompt `psql`:

```sql
CREATE DATABASE backend_go;
\c backend_go

CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Ketik `\q` untuk keluar dari `psql`. Bila memakai nama database, host, port, atau user lain, samakan nilainya di `.env`.

### 5. Jalankan aplikasi

```bash
go run app.go / air
```

Jika koneksi berhasil, Gin akan mendengarkan di `http://localhost:8080`. Biarkan terminal ini tetap berjalan selama API digunakan.

## CLI generator dan database

Semua command dijalankan dari root proyek dengan `go run .`:

```bash
go run . serve
go run . make:model User
go run . make:controller User
go run . make:middleware Auth
go run . make:request LoginRequest
go run . make:migration create_users
go run . make:crud User
go run . migrate
go run . db:seed
go run . route:list
```

Generator tidak menimpa berkas yang sudah ada. `make:crud` membuat model dan controller sekaligus; route tetap didaftarkan secara eksplisit di `Routes/routes.go`.

`make:migration` membuat berkas `.up.sql` di `Migrations/`. Isi SQL perubahan skema di berkas tersebut, lalu jalankan `migrate`. Setiap migration yang sukses dicatat pada tabel `schema_migrations` sehingga tidak dijalankan dua kali. Tambahkan seed spesifik aplikasi pada `Database/seed.go`.

### 6. Verifikasi API

Buka terminal baru dan panggil health check:

```bash
curl http://localhost:8080/health
```

Respons yang diharapkan:

```json
{"status":"ok"}
```

Di PowerShell, alternatifnya:

```powershell
Invoke-RestMethod http://localhost:8080/health
```

## Endpoint API

### Health check

```http
GET /health
```

### Daftar user

```http
GET /api/users
```

Respons sukses berbentuk:

```json
{
  "data": [
    {
      "id": "1",
      "email": "user@example.com",
      "name": "John Doe",
      "createdAt": "2025-12-06T10:00:00Z",
      "updatedAt": "2025-12-06T10:00:00Z"
    }
  ]
}
```

### Buat user

```http
POST /api/users
Content-Type: application/json
```

```json
{
  "id": "1",
  "email": "user@example.com",
  "name": "John Doe",
  "password": "password123"
}
```

Contoh menggunakan `curl`:

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"id":"1","email":"user@example.com","name":"John Doe","password":"password123"}'
```

Password di-hash memakai bcrypt dan tidak pernah dikembalikan API. Input yang tidak valid menghasilkan `422`; email atau ID yang sudah ada menghasilkan `409`.

## Perintah pengembangan

```bash
go test ./...
go build -o backend-go.exe app.go
```

## License

MIT
