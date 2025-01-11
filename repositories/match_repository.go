package repositories

import (
    "github.com/connectplus/models"
    "gorm.io/gorm"
)

type MatchRepository interface {
    Create(match *models.Match) error
    FindByUserID(userID uint) ([]models.Match, error)
    FindByUsers(user1ID, user2ID uint) (*models.Match, error)
    UpdateStatus(matchID uint, status models.MatchStatus) error
    Delete(matchID uint) error
}

type matchRepository struct {
    db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) MatchRepository {
    return &matchRepository{db: db}
}

func (r *matchRepository) Create(match *models.Match) error {
    return r.db.Create(match).Error
}

func (r *matchRepository) FindByUserID(userID uint) ([]models.Match, error) {
    var matches []models.Match
    err := r.db.Where("user1_id = ? OR user2_id = ?", userID, userID).Find(&matches).Error
    return matches, err
}

func (r *matchRepository) FindByUsers(user1ID, user2ID uint) (*models.Match, error) {
    var match models.Match
    err := r.db.Where("(user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?)",
        user1ID, user2ID, user2ID, user1ID).First(&match).Error
    return &match, err
}

func (r *matchRepository) UpdateStatus(matchID uint, status models.MatchStatus) error {
    return r.db.Model(&models.Match{}).Where("id = ?", matchID).Update("status", status).Error
}

func (r *matchRepository) Delete(matchID uint) error {
    return r.db.Delete(&models.Match{}, matchID).Error
}
