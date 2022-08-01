package models

type Score struct {
	ID         string             `json:"id"`
	Srv        map[string]Service `json:"srv"`
	TotalScore float64            `json:"total_score"`
}
type Service struct {
	Reputation int     `json:"reputation"`
	Gained     int     `json:"gained"`
	Lost       int     `json:"lost"`
	Score      float64 `json:"score"`
	SLA        float64 `json:"sla"`
	Status     int     `json:"status"`
}
