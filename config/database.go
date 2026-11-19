package config

import (
	"diet-tracker-backend/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// LoadEnv memuat variabel konfigurasi dari file .env
func LoadEnv() {
	// Mencoba memuat .env dari root project maupun jika dijalankan dari dalam folder backend-api
	if err := godotenv.Load(".env"); err != nil {
		_ = godotenv.Load("../.env")
	}
}

// getEnv mengambil nilai environment variable atau nilai default jika kosong
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// ConnectDatabase membuka koneksi ke PostgreSQL menggunakan konfigurasi dari .env
func ConnectDatabase() *gorm.DB {
	LoadEnv()

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "diet_tracker")
	sslMode := getEnv("DB_SSLMODE", "disable")
	timeZone := getEnv("DB_TIMEZONE", "Asia/Jakarta")

	// 1. Pastikan database (misal: 'diet_tracker') sudah ada di PostgreSQL
	ensurePostgresDatabaseExists(host, port, user, password, dbName, sslMode, timeZone)

	// 2. Buat DSN (Data Source Name) untuk terhubung ke database PostgreSQL target
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbName, port, sslMode, timeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal terhubung ke PostgreSQL (periksa DB_USER & DB_PASSWORD di file .env): %v", err)
	}

	fmt.Printf("✅ Berhasil terhubung ke PostgreSQL (Database: %s di %s:%s)\n", dbName, host, port)

	// 3. AutoMigrate otomatis membuat/memperbarui tabel-tabel di PostgreSQL
	err = db.AutoMigrate(
		&models.KamusMakanan{},
		&models.CatatanMakanan{},
		&models.CatatanOlahraga{},
		&models.TargetDiet{},
	)
	if err != nil {
		log.Fatalf("❌ Gagal migrasi tabel PostgreSQL: %v", err)
	}

	// 4. Isi data awal (Seeding) jika tabel masih kosong
	seedInitialData(db)
	return db
}

// ensurePostgresDatabaseExists otomatis membuat database di PostgreSQL jika belum dibuat di pgAdmin
func ensurePostgresDatabaseExists(host, port, user, password, dbName, sslMode, timeZone string) {
	defaultDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s TimeZone=%s",
		host, user, password, port, sslMode, timeZone,
	)

	adminDB, err := gorm.Open(postgres.Open(defaultDSN), &gorm.Config{})
	if err != nil {
		// Biarkan gorm.Open utama yang menampilkan pesan error lengkapnya
		return
	}

	sqlDB, err := adminDB.DB()
	if err == nil {
		defer sqlDB.Close()
	}

	var exists int
	adminDB.Raw("SELECT 1 FROM pg_database WHERE datname = ?", dbName).Scan(&exists)
	if exists == 0 {
		fmt.Printf("📦 Database '%s' belum ada di PostgreSQL, membuat database secara otomatis...\n", dbName)
		if err := adminDB.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", dbName)).Error; err != nil {
			log.Printf("⚠️ Gagal membuat database otomatis (silakan buat database '%s' di pgAdmin): %v", dbName, err)
		}
	}
}

// seedInitialData mengisi data awal (Kamus Makanan & Target) jika database PostgreSQL masih kosong
func seedInitialData(db *gorm.DB) {
	var countTarget int64
	db.Model(&models.TargetDiet{}).Count(&countTarget)
	if countTarget == 0 {
		db.Create(&models.TargetDiet{
			TargetKaloriHarian:  2000,
			TargetMenitHarian:   30,
			TargetMenitMingguan: 150,
			TargetMenitBulanan:  600,
		})
	}

	var countKamus int64
	db.Model(&models.KamusMakanan{}).Count(&countKamus)
	if countKamus == 0 {
		daftarKamus := []models.KamusMakanan{
			{Nama: "Nasi Putih", KaloriPerPorsi: 175},
			{Nama: "Nasi Merah", KaloriPerPorsi: 110},
			{Nama: "Telur Rebus", KaloriPerPorsi: 78},
			{Nama: "Telur Dadar", KaloriPerPorsi: 150},
			{Nama: "Ayam Bakar Dada", KaloriPerPorsi: 220},
			{Nama: "Ayam Bakar Paha", KaloriPerPorsi: 250},
			{Nama: "Ayam Goreng", KaloriPerPorsi: 280},
			{Nama: "Tempe Goreng", KaloriPerPorsi: 80},
			{Nama: "Tahu Goreng", KaloriPerPorsi: 70},
			{Nama: "Kopi Susu", KaloriPerPorsi: 180},
			{Nama: "Pisang", KaloriPerPorsi: 90},
		}
		db.Create(&daftarKamus)
	}

	hariIni := time.Now().Format("2006-01-02")
	var countMakanan int64
	db.Model(&models.CatatanMakanan{}).Count(&countMakanan)
	if countMakanan == 0 {
		db.Create(&[]models.CatatanMakanan{
			{Nama: "Nasi Putih", Porsi: 1, Kalori: 175, WaktuMakan: "Pagi", Tanggal: hariIni},
			{Nama: "Ayam Bakar Dada", Porsi: 1, Kalori: 220, WaktuMakan: "Siang", Tanggal: hariIni},
		})
	}

	var countOlahraga int64
	db.Model(&models.CatatanOlahraga{}).Count(&countOlahraga)
	if countOlahraga == 0 {
		db.Create(&models.CatatanOlahraga{
			Jenis:          "Jalan Kaki",
			DurasiMenit:    30,
			JarakKm:        2.4,
			KaloriTerbakar: 150,
			Tanggal:        hariIni,
		})
	}
}
