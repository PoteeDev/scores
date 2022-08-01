package models

type Score struct {
	ID         string             `bson:"id" json:"id"`
	Srv        map[string]Service `bson:"srv" json:"srv"`
	TotalScore float64            `bson:"total_score" json:"total_score"`
}
type Service struct {
	Reputation int     `bson:"reputation" json:"reputation"`
	Gained     int     `bson:"gained" json:"gained"`
	Lost       int     `bson:"lost" json:"lost"`
	Score      float64 `bson:"score" json:"score"`
	SLA        float64 `bson:"sla" json:"sla"`
	Status     int     `bson:"status" json:"status"`
}
