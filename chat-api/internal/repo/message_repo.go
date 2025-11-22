package repo

import (
	"context"
	"time"

	"github.com/chat-api/internal/models"
	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) Create(ctx context.Context, m *models.Message) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r MessageRepo) GetLastMessage(ctx context.Context, room string, limit int) ([]models.Message, error) {
	var msgs []models.Message
	err := r.db.WithContext(ctx).
		Where("room =?", room).
		Order("id DESC").
		Limit(limit).
		Find(&msgs).Error

	return msgs, err
}

func (r *MessageRepo) MarkDelivered(ctx context.Context, msgID uint, userID uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&models.Message{}).Where("id = ? AND user_id = ?", msgID, userID).Update("delivered_at", now).Error
}

func (r *MessageRepo) MarkSeen(ctx context.Context, msgID uint, userID uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&models.Message{}).Where("id = ? AND user_id = ?", msgID, userID).Update("seen_at", now).Error
}
