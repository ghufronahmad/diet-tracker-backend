package models

// TargetDiet merepresentasikan tabel SQL 'target_diets'
// Menyimpan pengaturan target personal Anda di database
type TargetDiet struct {
	ID                  uint `gorm:"primaryKey" json:"id"`
	TargetKaloriHarian  int  `gorm:"not null" json:"target_kalori_harian"`
	TargetMenitHarian   int  `gorm:"not null" json:"target_menit_harian"`
	TargetMenitMingguan int  `gorm:"not null" json:"target_menit_mingguan"`
	TargetMenitBulanan  int  `gorm:"not null" json:"target_menit_bulanan"`
}

// DetailProgresOlahraga adalah DTO (Data Transfer Object) untuk response JSON rekap olahraga
type DetailProgresOlahraga struct {
	TotalMenit     int     `json:"total_menit"`
	TargetMenit    int     `json:"target_menit"`
	Persentase     int     `json:"persentase"`
	TotalJarakKm   float64 `json:"total_jarak_km"`
	TotalKalori    int     `json:"total_kalori_terbakar"`
	TargetTercapai bool    `json:"target_tercapai"`
}

// RekapDashboard adalah DTO response JSON untuk halaman utama Flutter
type RekapDashboard struct {
	TanggalHariIni     string                `json:"tanggal_hari_ini"`
	KaloriMasukHariIni int                   `json:"kalori_masuk_hari_ini"`
	BatasKaloriHarian  int                   `json:"batas_kalori_harian"`
	SisaKuotaKalori    int                   `json:"sisa_kuota_kalori"`
	OlahragaHarian     DetailProgresOlahraga `json:"olahraga_harian"`
	OlahragaMingguan   DetailProgresOlahraga `json:"olahraga_mingguan"`
	OlahragaBulanan    DetailProgresOlahraga `json:"olahraga_bulanan"`
}
