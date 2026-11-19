# Diet & Workout Tracker - Backend API (Golang)

REST API Backend untuk aplikasi **Personal Diet & Workout Tracker**, dibangun dengan **Golang**, **Gin Framework**, **GORM**, dan **PostgreSQL**.

## Tech Stack
* **Bahasa:** Go (Golang)
* **Web Framework:** [Gin (`gin-gonic/gin`)](https://github.com/gin-gonic/gin)
* **ORM & Database:** [GORM](https://gorm.io/) + **PostgreSQL**
* **Environment Config:** `github.com/joho/godotenv`
* **Arsitektur:** Layered Architecture (`config`, `models`, `repositories`, `services`, `handlers`, `middleware`)

## Struktur Folder
```text
.
├── main.go                       # Entrypoint server
├── config/
│   └── database.go               # Koneksi PostgreSQL, AutoMigrate, & Seeding
├── models/
│   ├── food.go                   # Skema tabel SQL kamus_makanans & catatan_makanans
│   ├── workout.go                # Skema tabel SQL catatan_olahragas
│   └── target.go                 # Skema tabel SQL target_diets & DTO RekapDashboard
├── repositories/
│   └── diet_repository.go        # Layer query database (GORM)
├── services/
│   └── diet_service.go           # Layer logika bisnis & kalkulasi kalori/target
├── handlers/
│   └── diet_handler.go           # Layer HTTP Controller (Gin)
├── middleware/
│   └── cors.go                   # CORS Middleware
├── .env.example                  # Template konfigurasi environment
└── go.mod
```

## Cara Menjalankan
1. Salin `.env.example` menjadi `.env` dan sesuaikan konfigurasi PostgreSQL Anda:
   ```env
   APP_PORT=8080
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=password_postgres_anda
   DB_NAME=diet_tracker
   DB_SSLMODE=disable
   DB_TIMEZONE=Asia/Jakarta
   ```
2. Jalankan server:
   ```bash
   go run .
   ```
   *(Server akan otomatis membuat database `diet_tracker` beserta tabel-tabelnya di PostgreSQL jika belum ada).*

## Daftar Endpoint API (`http://localhost:8080/api`)
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| `GET` | `/api/rekap` | Rekap kalori harian & progres olahraga (Harian, Mingguan, Bulanan) |
| `PUT` | `/api/target` | Memperbarui target kalori & target menit olahraga |
| `GET` | `/api/kamus-makanan` | Daftar kamus makanan & kalori per porsi |
| `POST` | `/api/kamus-makanan` | Menambah menu makanan baru ke database kamus |
| `GET` | `/api/makanan` | Daftar riwayat makanan yang masuk |
| `POST` | `/api/makanan` | Mencatat makanan masuk (kalori otomatis dihitung dari kamus) |
| `DELETE` | `/api/makanan/:id` | Menghapus catatan makanan berdasarkan ID |
| `GET` | `/api/olahraga` | Daftar riwayat olahraga (Jalan Kaki / Jogging) |
| `POST` | `/api/olahraga` | Mencatat olahraga baru (kalori terbakar otomatis dihitung) |
| `DELETE` | `/api/olahraga/:id` | Menghapus catatan olahraga berdasarkan ID |
