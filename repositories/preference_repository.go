package repositories

import (
    "github.com/connectplus/models"
    "gorm.io/gorm"
)

type PreferenceRepository interface {
    Create(preference *models.Preference) error
    FindByUserID(userID uint) (*models.Preference, error)
    Update(preference *models.Preference) error
    Delete(userID uint) error
}

type preferenceRepository struct {
    db *gorm.DB
}

func NewPreferenceRepository(db *gorm.DB) PreferenceRepository {
    return &preferenceRepository{db: db}
}

func (r *preferenceRepository) Create(preference *models.Preference) error {
    return r.db.Create(preference).Error
}

func (r *preferenceRepository) FindByUserID(userID uint) (*models.Preference, error) {
    var preference models.Preference
    err := r.db.Where("user_id = ?", userID).First(&preference).Error
    return &preference, err
}

func (r *preferenceRepository) Update(preference *models.Preference) error {
    return r.db.Save(preference).Error
}

func (r *preferenceRepository) Delete(userID uint) error {
    return r.db.Where("user_id = ?", userID).Delete(&models.Preference{}).Error
}
