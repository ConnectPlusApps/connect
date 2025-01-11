package repositories

import (
    "github.com/connectplus/models"
    "gorm.io/gorm"
)

type ProfileRepository interface {
    Create(profile *models.Profile) error
    FindByUserID(userID uint) (*models.Profile, error)
    Update(profile *models.Profile) error
    Delete(userID uint) error
}

type profileRepository struct {
    db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
    return &profileRepository{db: db}
}

func (r *profileRepository) Create(profile *models.Profile) error {
    return r.db.Create(profile).Error
}

func (r *profileRepository) FindByUserID(userID uint) (*models.Profile, error) {
    var profile models.Profile
    err := r.db.Where("user_id = ?", userID).First(&profile).Error
    return &profile, err
}

func (r *profileRepository) Update(profile *models.Profile) error {
    return r.db.Save(profile).Error
}

func (r *profileRepository) Delete(userID uint) error {
    return r.db.Where("user_id = ?", userID).Delete(&models.Profile{}).Error
}
