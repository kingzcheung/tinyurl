package core

type AnalyticsTyp string

const (
	Source  AnalyticsTyp = "source"
	Region  AnalyticsTyp = "region"
	OS      AnalyticsTyp = "os"
	Device  AnalyticsTyp = "device"
	Browser AnalyticsTyp = "browser"
	Day     AnalyticsTyp = "day"
	Month   AnalyticsTyp = "month"
)

type Analytics struct {
	Typ   AnalyticsTyp `json:"type"`
	Item  string       `json:"item"`
	Count int64        `json:"count"`
}

type AnalyticsStore interface {
	AnalyticsBy(typ AnalyticsTyp, urlID int, start, end int64) []*Analytics
}
