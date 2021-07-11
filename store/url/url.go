package url

import (
	"context"
	"github.com/kingzcheung/tinyurl/core"
	"gorm.io/gorm"
	"time"
)

type urlStore struct {
	db *gorm.DB
}

func NewUrlStore(db *gorm.DB) core.UrlStore {
	return &urlStore{db: db}
}

func (s *urlStore) UpdateSlug(ctx context.Context, id int, slug string) error {
	return s.db.WithContext(ctx).
		Model(&core.Url{}).
		Where("url_id", id).
		Update("slug", slug).
		Error
}

func (s *urlStore) Create(ctx context.Context, url *core.Url) (*core.Url, error) {
	if url.CreatedAt == 0 {
		url.CreatedAt = time.Now().Unix()
	}
	if url.UpdatedAt == 0 {
		url.UpdatedAt = time.Now().Unix()
	}
	err := s.db.WithContext(ctx).Create(url).Error

	return url, err

}

func (s *urlStore) GetByCode(ctx context.Context, val string) *core.Url {
	var u core.Url
	s.db.WithContext(ctx).Where("slug = ?", val).First(&u)
	return &u
}

func (s *urlStore) List(ctx context.Context, opt core.UrlOption) ([]*core.Url, int64) {
	var urls []*core.Url
	var count int64

	tx := s.db.WithContext(ctx).Order("url_id desc")
	if opt.SearchVal != "" {
		tx = tx.Where("short_code = ?", opt.SearchVal).
			Or("name like ?", opt.SearchVal+"%").
			Or("url like ?", opt.SearchVal+"%")
	}
	var limit = 10
	if opt.Limit > 0 {
		limit = opt.Limit
	}
	tx = tx.Limit(limit)
	if opt.Page > 0 {
		tx = tx.Offset((opt.Page - 1) * limit)
	}

	tx.Find(&urls).Limit(-1).Offset(-1).Count(&count)

	return urls, count

}
