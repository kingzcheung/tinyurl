package analytics

import (
	"gorm.io/gorm"
	"github.com/kingzcheung/tinyurl/core"
	"time"
)

type analyticsStore struct {
	db *gorm.DB
}

func NewAnalyticsStore(db *gorm.DB) *analyticsStore {
	return &analyticsStore{db: db}
}

var sqlMap = map[core.AnalyticsTyp]string{
	core.OS:      `select os as item,count(os) as count from tracking where url_id = ? and created_at between ? and ? group by os`,
	core.Source:  `select source as item,count(source) as count from tracking where url_id = ? and created_at between ? and ? group by source`,
	core.Device:  `select device as item,count(device) as count from tracking where url_id = ? and created_at between ? and ? group by device`,
	core.Browser: `select browser as item,count(browser) as count from tracking where url_id = ? and created_at between ? and ? group by browser`,
	core.Region:  `select region as item,count(region) as count from tracking where url_id = ? and created_at between ? and ? group by region`,
	core.Day:     `select FROM_UNIXTIME(created_at, '%Y-%c-%e') item, count(*) count from tracking where url_id = ? and created_at between ? and ? group by FROM_UNIXTIME(created_at, '%Y-%m-%d')`,
	core.Month:   `select FROM_UNIXTIME(created_at, '%Y-%c') item, count(*) count from tracking where url_id = ? and created_at between ? and ? group by FROM_UNIXTIME(created_at, '%Y-%m')`,
}

func (s *analyticsStore) AnalyticsBy(typ core.AnalyticsTyp, urlID int, start, end int64) []*core.Analytics {
	var anas []*core.Analytics
	if end == 0 {
		end = time.Now().Unix()
	}
	sql, ok := sqlMap[typ]
	if !ok {
		return nil
	}
	s.db.Raw(sql, urlID, start, end).Scan(&anas)
	for _, ana := range anas {
		ana.Typ = typ
	}
	return anas
}
