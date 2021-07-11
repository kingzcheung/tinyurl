package core

import "context"

type Url struct {
	UrlID     int    `json:"url_id" form:"url_id" gorm:"autoIncrement;primaryKey;column:url_id;type:int" `
	URL       string `json:"url" form:"url" gorm:"column:url;type:string" `
	Slug      string `json:"slug" form:"slug" gorm:"column:slug;type:string;unique" `
	Name      string `json:"name" form:"name" gorm:"column:name;type:string" `
	Type      int8   `json:"type" form:"type" gorm:"column:type;type:int" `
	CreatedAt int64  `json:"created_at" form:"created_at" gorm:"column:created_at;type:int;autoCreateTime" `
	UpdatedAt int64  `json:"updated_at" form:"updated_at" gorm:"column:updated_at;type:int;autoCreateTime" `
	ExpiredAt int64  `json:"expired_at" form:"expired_at" gorm:"column:expired_at;type:int" `
}

func (Url) TableName() string {
	return "urls"
}

type UrlStore interface {
	Create(ctx context.Context,url *Url) (*Url, error)
	UpdateSlug(ctx context.Context, id int,slug string) error
	GetByCode(ctx context.Context,val string) *Url
	List(ctx context.Context,option UrlOption) ([]*Url, int64)
}
