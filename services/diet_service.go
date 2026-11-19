package services

import (
	"diet-tracker-backend/models"
	"diet-tracker-backend/repositories"
	"strings"
	"time"
)

// DietService menyimpan seluruh logika bisnis (kalkulasi kalori, persentase target, dll.)
type DietService struct {
	repo *repositories.DietRepository
}

func NewDietService(repo *repositories.DietRepository) *DietService {
	return &DietService{repo: repo}
}

// GetKamusMap mengembalikan map[nama]kalori agar kompatibel dengan dropdown Flutter
func (s *DietService) GetKamusMap() (map[string]int, error) {
	items, err := s.repo.GetAllKamus()
	if err != nil {
		return nil, err
	}
	result := make(map[string]int, len(items))
	for _, item := range items {
		result[item.Nama] = item.KaloriPerPorsi
	}
	return result, nil
}

func (s *DietService) TambahKamusMakanan(input models.KamusMakanan) error {
	return s.repo.UpsertKamus(&input)
}

func (s *DietService) GetDaftarMakanan() ([]models.CatatanMakanan, error) {
	return s.repo.GetAllMakanan()
}

func (s *DietService) CatatMakanan(input *models.CatatanMakanan) error {
	if input.Porsi <= 0 {
		input.Porsi = 1
	}
	if input.Tanggal == "" {
		input.Tanggal = time.Now().Format("2006-01-02")
	}
	// Jika kalori tidak diisi (0), ambil otomatis dari tabel SQL kamus_makanans
	if input.Kalori == 0 {
		if kamus, err := s.repo.FindKamusByNama(input.Nama); err == nil && kamus != nil {
			input.Nama = kamus.Nama
			input.Kalori = kamus.KaloriPerPorsi * input.Porsi
		}
	}
	return s.repo.CreateMakanan(input)
}

func (s *DietService) HapusMakanan(id uint) error {
	return s.repo.DeleteMakanan(id)
}

func (s *DietService) GetDaftarOlahraga() ([]models.CatatanOlahraga, error) {
	return s.repo.GetAllOlahraga()
}

func (s *DietService) CatatOlahraga(input *models.CatatanOlahraga) error {
	if input.Tanggal == "" {
		input.Tanggal = time.Now().Format("2006-01-02")
	}
	if input.KaloriTerbakar == 0 {
		if strings.EqualFold(input.Jenis, "Jogging") {
			input.KaloriTerbakar = input.DurasiMenit * 9 // ~9 kkal/menit
		} else {
			input.KaloriTerbakar = input.DurasiMenit * 5 // Jalan Kaki ~5 kkal/menit
		}
	}
	return s.repo.CreateOlahraga(input)
}

func (s *DietService) HapusOlahraga(id uint) error {
	return s.repo.DeleteOlahraga(id)
}

func (s *DietService) UpdateTargetDiet(input models.TargetDiet) (models.TargetDiet, error) {
	return s.repo.UpdateTarget(input)
}

// GetRekapDashboard menghitung statistik Harian, Mingguan, dan Bulanan dari database
func (s *DietService) GetRekapDashboard() (*models.RekapDashboard, error) {
	waktuSekarang := time.Now()
	hariIniStr := waktuSekarang.Format("2006-01-02")
	bulanIniPrefix := waktuSekarang.Format("2006-01")
	batas7HariLalu := waktuSekarang.AddDate(0, 0, -6).Format("2006-01-02")

	target, err := s.repo.GetTarget()
	if err != nil {
		return nil, err
	}

	makananHariIni, err := s.repo.GetMakananByTanggal(hariIniStr)
	if err != nil {
		return nil, err
	}

	kaloriHariIni := 0
	for _, m := range makananHariIni {
		kaloriHariIni += m.Kalori
	}

	semuaOlahraga, err := s.repo.GetAllOlahraga()
	if err != nil {
		return nil, err
	}

	var menitHari, kalHari int
	var jarakHari float64

	var menitMinggu, kalMinggu int
	var jarakMinggu float64

	var menitBulan, kalBulan int
	var jarakBulan float64

	for _, o := range semuaOlahraga {
		if o.Tanggal == hariIniStr {
			menitHari += o.DurasiMenit
			jarakHari += o.JarakKm
			kalHari += o.KaloriTerbakar
		}
		if o.Tanggal >= batas7HariLalu && o.Tanggal <= hariIniStr {
			menitMinggu += o.DurasiMenit
			jarakMinggu += o.JarakKm
			kalMinggu += o.KaloriTerbakar
		}
		if strings.HasPrefix(o.Tanggal, bulanIniPrefix) {
			menitBulan += o.DurasiMenit
			jarakBulan += o.JarakKm
			kalBulan += o.KaloriTerbakar
		}
	}

	rekap := &models.RekapDashboard{
		TanggalHariIni:     hariIniStr,
		KaloriMasukHariIni: kaloriHariIni,
		BatasKaloriHarian:  target.TargetKaloriHarian,
		SisaKuotaKalori:    target.TargetKaloriHarian - kaloriHariIni,
		OlahragaHarian:     hitungDetailProgres(menitHari, target.TargetMenitHarian, jarakHari, kalHari),
		OlahragaMingguan:   hitungDetailProgres(menitMinggu, target.TargetMenitMingguan, jarakMinggu, kalMinggu),
		OlahragaBulanan:    hitungDetailProgres(menitBulan, target.TargetMenitBulanan, jarakBulan, kalBulan),
	}

	return rekap, nil
}

func hitungDetailProgres(totalMenit, targetMenit int, totalJarak float64, totalKalori int) models.DetailProgresOlahraga {
	persen := 0
	if targetMenit > 0 {
		persen = (totalMenit * 100) / targetMenit
		if persen > 100 {
			persen = 100
		}
	}
	return models.DetailProgresOlahraga{
		TotalMenit:     totalMenit,
		TargetMenit:    targetMenit,
		Persentase:     persen,
		TotalJarakKm:   totalJarak,
		TotalKalori:    totalKalori,
		TargetTercapai: totalMenit >= targetMenit,
	}
}
