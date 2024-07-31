package models

import (
	"time"
)

type MataKuliah struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Dosen struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Nama         string     `json:"nama"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	MataKuliahID uint       `json:"mata_kuliah_id"`
	MataKuliah   MataKuliah `json:"mata_kuliah" gorm:"foreignKey:MataKuliahID"`
}

type Users struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Mahasiswa struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Nilai struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Indeks       string     `json:"indeks"`
	Skor         int        `json:"skor"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	MahasiswaID  uint       `json:"mahasiswa_id"`
	MataKuliahID uint       `json:"mata_kuliah_id"`
	UserID       uint       `json:"user_id"`
	Mahasiswa    Mahasiswa  `json:"mahasiswa" gorm:"foreignKey:MahasiswaID"`
	MataKuliah   MataKuliah `json:"mata_kuliah" gorm:"foreignKey:MataKuliahID"`
	User         User       `json:"user" gorm:"foreignKey:UserID"`
}

type JadwalKuliah struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	DosenID     uint      `json:"dosen_id"`
	MahasiswaID uint      `json:"mahasiswa_id"`
	Hari        string    `json:"hari"`
	JamMulai    time.Time `json:"jam_mulai"`
	JamSelesai  time.Time `json:"jam_selesai"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Dosen       Dosen     `json:"dosen" gorm:"foreignKey:DosenID"`
	Mahasiswa   Mahasiswa `json:"mahasiswa" gorm:"foreignKey:MahasiswaID"`
}
