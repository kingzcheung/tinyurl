package core

type Session struct {
	Token  string `json:"token" gorm:"primaryKey;column:token;type:string"`
	Data   string `json:"data" gorm:"column:data;type:bytes"`
	Expiry string `json:"expiry" gorm:"column:expiry;type:time"`
}

func (s Session) TableName() string {
	return "sessions"
}
