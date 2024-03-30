package repository

import (
	"github.com/Dev4w4n/e-masjid.my/api/tetapan-public-api/model"
	"gorm.io/gorm"
)

type TetapanRepository interface {
	FindAll() ([]model.Tetapan, error)
	FindByKunci(kunci string) (model.Tetapan, error)
}

type TetapanRepositoryImpl struct {
	Db *gorm.DB
}

func NewTetapanRepository(db *gorm.DB) TetapanRepository {
	db.AutoMigrate(&model.Tetapan{})

	return &TetapanRepositoryImpl{Db: db}
}

func (repo *TetapanRepositoryImpl) FindAll() ([]model.Tetapan, error) {
	var tetapanList []model.Tetapan
	result := repo.Db.Find(&tetapanList)

	if result.Error != nil {
		return nil, result.Error
	}

	return tetapanList, nil
}

func (repo *TetapanRepositoryImpl) FindByKunci(kunci string) (model.Tetapan, error) {
	var tetapan model.Tetapan
	result := repo.Db.First(&tetapan, "kunci = ?", kunci)

	if result.Error != nil {
		return model.Tetapan{}, result.Error
	}

	return tetapan, nil
}
