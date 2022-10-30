package models

type Score struct {
	ID          string             `bson:"id" json:"id"`
	Srv         map[string]Service `bson:"srv" json:"srv"`
	TotalScore  float64            `bson:"total_score" json:"total_score"`
	TotalGained int                `bson:"total_gained" json:"total_gained"`
	TotalLost   int                `bson:"total_lost" json:"total_lost"`
	Place       int                `bson:"place" json:"place"`
	LastPlace   int                `bson:"last_place" json:"last_place"`
}
type Service struct {
	Reputation int     `bson:"reputation" json:"reputation"`
	Gained     int     `bson:"gained" json:"gained"`
	Lost       int     `bson:"lost" json:"lost"`
	Score      float64 `bson:"score" json:"score"`
	SLA        float64 `bson:"sla" json:"sla"`
	Status     int     `bson:"status" json:"status"`
}
