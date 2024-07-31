package controllers

import (
	"backend-tugas-reactjs/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateMataKuliah creates a new MataKuliah
// @Summary Create a new MataKuliah
// @Description Create a new MataKuliah
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Param mataKuliah body models.MataKuliah true "MataKuliah"
// @Success 200 {object} models.MataKuliah
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /mata-kuliah [post]
func CreateMataKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mataKuliah models.MataKuliah

	if err := c.ShouldBindJSON(&mataKuliah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mataKuliah})
}

// GetMataKuliah retrieves all MataKuliah
// @Summary Get all MataKuliah
// @Description Get all MataKuliah
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Success 200 {array} models.MataKuliah
// @Failure 500 {object} gin.H{"error": string}
// @Router /mata-kuliah [get]
func GetMataKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mataKuliah []models.MataKuliah

	if err := db.Find(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mataKuliah})
}

// UpdateMataKuliah updates an existing MataKuliah
// @Summary Update an existing MataKuliah
// @Description Update an existing MataKuliah
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Param id path string true "MataKuliah ID"
// @Param mataKuliah body models.MataKuliah true "MataKuliah"
// @Success 200 {object} models.MataKuliah
// @Failure 400 {object} gin.H{"error": string}
// @Failure 404 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /mata-kuliah/{id} [put]
func UpdateMataKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mataKuliah models.MataKuliah

	if err := db.Where("id = ?", c.Param("id")).First(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&mataKuliah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mataKuliah})
}

// DeleteMataKuliah deletes an existing MataKuliah
// @Summary Delete an existing MataKuliah
// @Description Delete an existing MataKuliah
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Param id path string true "MataKuliah ID"
// @Success 200 {object} gin.H{"data": true}
// @Failure 404 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /mata-kuliah/{id} [delete]
func DeleteMataKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mataKuliah models.MataKuliah

	if err := db.Where("id = ?", c.Param("id")).First(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Delete(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// CreateDosen creates a new Dosen
// @Summary Create a new Dosen
// @Description Create a new Dosen
// @Tags Dosen
// @Accept json
// @Produce json
// @Param dosen body models.Dosen true "Dosen"
// @Success 200 {object} models.Dosen
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /dosen [post]
func CreateDosen(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dosen models.Dosen

	if err := c.ShouldBindJSON(&dosen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&dosen).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dosen})
}

// GetDosen retrieves all Dosen
// @Summary Get all Dosen
// @Description Get all Dosen
// @Tags Dosen
// @Accept json
// @Produce json
// @Success 200 {array} models.Dosen
// @Failure 500 {object} gin.H{"error": string}
// @Router /dosen [get]
func GetDosen(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dosen []models.Dosen

	if err := db.Find(&dosen).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dosen})
}

// UpdateDosen updates an existing Dosen
// @Summary Update an existing Dosen
// @Description Update an existing Dosen
// @Tags Dosen
// @Accept json
// @Produce json
// @Param id path string true "Dosen ID"
// @Param dosen body models.Dosen true "Dosen"
// @Success 200 {object} models.Dosen
// @Failure 400 {object} gin.H{"error": string}
// @Failure 404 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /dosen/{id} [put]
func UpdateDosen(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dosen models.Dosen

	if err := db.Where("id = ?", c.Param("id")).First(&dosen).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&dosen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&dosen).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dosen})
}

// DeleteDosen deletes an existing Dosen
// @Summary Delete an existing Dosen
// @Description Delete an existing Dosen
// @Tags Dosen
// @Accept json
// @Produce json
// @Param id path string true "Dosen ID"
// @Success 200 {object} gin.H{"data": true}
// @Failure 404 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /dosen/{id} [delete]
func DeleteDosen(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dosen models.Dosen

	if err := db.Where("id = ?", c.Param("id")).First(&dosen).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Delete(&dosen).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// CreateMahasiswa creates a new Mahasiswa
// @Summary Create a new Mahasiswa
// @Description Create a new Mahasiswa
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param mahasiswa body models.Mahasiswa true "Mahasiswa"
// @Success 200 {object} models.Mahasiswa
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /mahasiswa [post]
func CreateMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa models.Mahasiswa

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

// GetMahasiswa retrieves all Mahasiswa
// @Summary Get all Mahasiswa
// @Description Get all Mahasiswa
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Success 200 {array} models.Mahasiswa
// @Failure 500 {object} gin.H{"error": string}
// @Router /mahasiswa [get]
func GetMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa []models.Mahasiswa

	if err := db.Find(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

// UpdateMahasiswa updates an existing Mahasiswa
// @Summary Update an existing Mahasiswa
// @Description Update an existing Mahasiswa
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Mahasiswa ID"
// @Param mahasiswa body models.Mahasiswa true "Mahasiswa"
// @Success 200 {object} models.Mahasiswa
// @Failure 400 {object} gin.H{"error": string}
// @Failure 404 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /mahasiswa/{id} [put]
func UpdateMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa models.Mahasiswa

	if err := db.Where("id = ?", c.Param("id")).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

// DeleteMahasiswa deletes an existing Mahasiswa
// @Summary Delete an existing Mahasiswa
// @Description Delete an existing Mahasiswa
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Mahasiswa ID"
// @Success 200 {object} gin.H{"data": true}
// @Failure 404 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /mahasiswa/{id} [delete]
func DeleteMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa models.Mahasiswa

	if err := db.Where("id = ?", c.Param("id")).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Delete(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// CreateNilai creates a new Nilai
func CreateNilai(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var nilai models.Nilai

	if err := c.ShouldBindJSON(&nilai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&nilai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nilai})
}

// GetNilai retrieves all Nilai
func GetNilai(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var nilai []models.Nilai

	if err := db.Find(&nilai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nilai})
}

// UpdateNilai updates an existing Nilai
func UpdateNilai(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var nilai models.Nilai

	if err := db.Where("id = ?", c.Param("id")).First(&nilai).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&nilai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&nilai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nilai})
}

// DeleteNilai deletes an existing Nilai
func DeleteNilai(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var nilai models.Nilai

	if err := db.Where("id = ?", c.Param("id")).First(&nilai).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Delete(&nilai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// CreateJadwalKuliah creates a new JadwalKuliah
func CreateJadwalKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var jadwalKuliah models.JadwalKuliah

	if err := c.ShouldBindJSON(&jadwalKuliah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&jadwalKuliah).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jadwalKuliah})
}

// GetJadwalKuliah retrieves all JadwalKuliah
func GetJadwalKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var jadwalKuliah []models.JadwalKuliah

	if err := db.Find(&jadwalKuliah).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jadwalKuliah})
}

// UpdateJadwalKuliah updates an existing JadwalKuliah
func UpdateJadwalKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var jadwalKuliah models.JadwalKuliah

	if err := db.Where("id = ?", c.Param("id")).First(&jadwalKuliah).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&jadwalKuliah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&jadwalKuliah).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jadwalKuliah})
}

// DeleteJadwalKuliah deletes an existing JadwalKuliah
func DeleteJadwalKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var jadwalKuliah models.JadwalKuliah

	if err := db.Where("id = ?", c.Param("id")).First(&jadwalKuliah).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Delete(&jadwalKuliah).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
