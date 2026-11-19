package models

// KamusMakanan merepresentasikan tabel SQL 'kamus_makanans'
// Menyimpan daftar makanan beserta kalori per porsi di database
type KamusMakanan struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Nama           string `gorm:"uniqueIndex;not null" json:"nama"`
	KaloriPerPorsi int    `gorm:"not null" json:"kalori_per_porsi"`
}

// CatatanMakanan merepresentasikan tabel SQL 'catatan_makanans'
type CatatanMakanan struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Nama       string `gorm:"not null" json:"nama"`
	Porsi      int    `gorm:"not null;default:1" json:"porsi"`
	Kalori     int    `gorm:"not null" json:"kalori"`
	WaktuMakan string `gorm:"not null" json:"waktu_makan"` // Pagi, Siang, Malam, Snack
	Tanggal    string `gorm:"index;not null" json:"tanggal"` // Format: YYYY-MM-DD
}
