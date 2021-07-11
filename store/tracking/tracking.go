package tracking

import (
	"gorm.io/gorm"
	"github.com/kingzcheung/tinyurl/core"
)

type trackingStore struct {
	db *gorm.DB
}

func NewTrackingStore(db *gorm.DB) core.VisitStore {
	return &trackingStore{db: db}
}

func (s *trackingStore) Create(tracking *core.Tracking) (*core.Tracking, error) {
	return tracking, s.db.Create(tracking).Error
}
