package models

// CatatanOlahraga merepresentasikan tabel SQL 'catatan_olahragas'
type CatatanOlahraga struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	Jenis          string  `gorm:"not null" json:"jenis"`           // "Jalan Kaki" atau "Jogging"
	DurasiMenit    int     `gorm:"not null" json:"durasi_menit"`    // Durasi (menit)
	JarakKm        float64 `gorm:"not null" json:"jarak_km"`        // Jarak (km)
	KaloriTerbakar int     `gorm:"not null" json:"kalori_terbakar"` // Kalori yang dibakar
	Tanggal        string  `gorm:"index;not null" json:"tanggal"`   // Format: YYYY-MM-DD
}
