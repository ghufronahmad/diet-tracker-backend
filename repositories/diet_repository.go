package repositories

import (
	"diet-tracker-backend/models"

	"gorm.io/gorm"
)

// DietRepository bertanggung jawab penuh atas seluruh query SQL ke database
type DietRepository struct {
	db *gorm.DB
}

func NewDietRepository(db *gorm.DB) *DietRepository {
	return &DietRepository{db: db}
}

// --- KAMUS MAKANAN ---

func (r *DietRepository) GetAllKamus() ([]models.KamusMakanan, error) {
	var kamus []models.KamusMakanan
	err := r.db.Order("nama asc").Find(&kamus).Error
	return kamus, err
}

func (r *DietRepository) FindKamusByNama(nama string) (*models.KamusMakanan, error) {
	var item models.KamusMakanan
	err := r.db.Where("LOWER(nama) = LOWER(?)", nama).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *DietRepository) UpsertKamus(item *models.KamusMakanan) error {
	existing, err := r.FindKamusByNama(item.Nama)
	if err == nil && existing != nil {
		existing.KaloriPerPorsi = item.KaloriPerPorsi
		return r.db.Save(existing).Error
	}
	return r.db.Create(item).Error
}

// --- CATATAN MAKANAN ---

func (r *DietRepository) GetAllMakanan() ([]models.CatatanMakanan, error) {
	var list []models.CatatanMakanan
	err := r.db.Order("id asc").Find(&list).Error
	return list, err
}

func (r *DietRepository) GetMakananByTanggal(tanggal string) ([]models.CatatanMakanan, error) {
	var list []models.CatatanMakanan
	err := r.db.Where("tanggal = ?", tanggal).Find(&list).Error
	return list, err
}

func (r *DietRepository) CreateMakanan(m *models.CatatanMakanan) error {
	return r.db.Create(m).Error
}

func (r *DietRepository) DeleteMakanan(id uint) error {
	return r.db.Delete(&models.CatatanMakanan{}, id).Error
}

// --- CATATAN OLAHRAGA ---

func (r *DietRepository) GetAllOlahraga() ([]models.CatatanOlahraga, error) {
	var list []models.CatatanOlahraga
	err := r.db.Order("id asc").Find(&list).Error
	return list, err
}

func (r *DietRepository) CreateOlahraga(o *models.CatatanOlahraga) error {
	return r.db.Create(o).Error
}

func (r *DietRepository) DeleteOlahraga(id uint) error {
	return r.db.Delete(&models.CatatanOlahraga{}, id).Error
}

// --- TARGET DIET ---

func (r *DietRepository) GetTarget() (models.TargetDiet, error) {
	var target models.TargetDiet
	err := r.db.First(&target).Error
	return target, err
}

func (r *DietRepository) UpdateTarget(input models.TargetDiet) (models.TargetDiet, error) {
	var target models.TargetDiet
	if err := r.db.First(&target).Error; err != nil {
		return target, err
	}
	target.TargetKaloriHarian = input.TargetKaloriHarian
	target.TargetMenitHarian = input.TargetMenitHarian
	target.TargetMenitMingguan = input.TargetMenitMingguan
	target.TargetMenitBulanan = input.TargetMenitBulanan
	err := r.db.Save(&target).Error
	return target, err
}
