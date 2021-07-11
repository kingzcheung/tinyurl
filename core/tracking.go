package core

type Tracking struct {
	TrackingID int    `json:"tracking_id" form:"tracking_id" gorm:"autoIncrement;column:tracking_id;type:int" `
	UrlID      int    `json:"url_id" form:"url_id" gorm:"column:url_id;type:int" `
	IP         string `json:"ip" form:"ip" gorm:"column:ip;type:string" `
	Source     string `json:"source" form:"source" gorm:"column:source;type:string" `
	Region     string `json:"region" form:"region" gorm:"column:region;type:string" `
	OS         string `json:"os" form:"os" gorm:"column:os;type:string" `
	Device     string `json:"device" form:"device" gorm:"column:device;type:string" `
	Browser    string `json:"browser" form:"browser" gorm:"column:browser;type:string" `
	CreatedAt  int64  `json:"created_at" form:"created_at" gorm:"column:created_at;type:int;autoCreateTime" `
}

func (Tracking) TableName() string {
	return "tracking"
}

type VisitStore interface {
	Create(tracking *Tracking) (*Tracking, error)
}
