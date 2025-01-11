package repositories

import (
    "github.com/connectplus/models"
    "gorm.io/gorm"
)

type MessageRepository interface {
    Create(message *models.Message) error
    GetConversation(user1ID, user2ID uint) ([]models.Message, error)
    MarkAsRead(messageID uint) error
    Delete(messageID uint) error
}

type messageRepository struct {
    db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
    return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *models.Message) error {
    return r.db.Create(message).Error
}

func (r *messageRepository) GetConversation(user1ID, user2ID uint) ([]models.Message, error) {
    var messages []models.Message
    err := r.db.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
        user1ID, user2ID, user2ID, user1ID).Order("created_at asc").Find(&messages).Error
    return messages, err
}

func (r *messageRepository) MarkAsRead(messageID uint) error {
    return r.db.Model(&models.Message{}).Where("id = ?", messageID).Update("is_read", true).Error
}

func (r *messageRepository) Delete(messageID uint) error {
    return r.db.Delete(&models.Message{}, messageID).Error
}
