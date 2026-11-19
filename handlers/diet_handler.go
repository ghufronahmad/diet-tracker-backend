package handlers

import (
	"diet-tracker-backend/models"
	"diet-tracker-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DietHandler menangani request HTTP menggunakan framework Gin
type DietHandler struct {
	service *services.DietService
}

func NewDietHandler(service *services.DietService) *DietHandler {
	return &DietHandler{service: service}
}

// GET /api/kamus-makanan
func (h *DietHandler) GetKamusMakanan(c *gin.Context) {
	kamus, err := h.service.GetKamusMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kamus)
}

// POST /api/kamus-makanan (Menambah menu makanan baru ke database kamus)
func (h *DietHandler) TambahKamusMakanan(c *gin.Context) {
	var input models.KamusMakanan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data kamus tidak valid"})
		return
	}
	if err := h.service.TambahKamusMakanan(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// GET /api/makanan
func (h *DietHandler) GetMakanan(c *gin.Context) {
	data, err := h.service.GetDaftarMakanan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// POST /api/makanan
func (h *DietHandler) CreateMakanan(c *gin.Context) {
	var input models.CatatanMakanan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON makanan tidak valid"})
		return
	}
	if err := h.service.CatatMakanan(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// DELETE /api/makanan/:id
func (h *DietHandler) DeleteMakanan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := h.service.HapusMakanan(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Catatan makanan berhasil dihapus"})
}

// GET /api/olahraga
func (h *DietHandler) GetOlahraga(c *gin.Context) {
	data, err := h.service.GetDaftarOlahraga()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// POST /api/olahraga
func (h *DietHandler) CreateOlahraga(c *gin.Context) {
	var input models.CatatanOlahraga
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON olahraga tidak valid"})
		return
	}
	if err := h.service.CatatOlahraga(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// DELETE /api/olahraga/:id
func (h *DietHandler) DeleteOlahraga(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := h.service.HapusOlahraga(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Catatan olahraga berhasil dihapus"})
}

// GET /api/rekap
func (h *DietHandler) GetRekap(c *gin.Context) {
	rekap, err := h.service.GetRekapDashboard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rekap)
}

// PUT /api/target (Mengubah target kalori & olahraga harian/mingguan/bulanan)
func (h *DietHandler) UpdateTarget(c *gin.Context) {
	var input models.TargetDiet
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format target tidak valid"})
		return
	}
	target, err := h.service.UpdateTargetDiet(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, target)
}
