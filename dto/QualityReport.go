package dto

type QualityReport struct {
	Excellent  int `json:"excellent,omitempty" bson:"excellent,omitempty"`
	Good       int `json:"good,omitempty" bson:"good,omitempty"`
	Fair       int `json:"fair,omitempty" bson:"fair,omitempty"`
	TotalCalls int `json:"totalCalls,omitempty" bson:"totalCalls,omitempty"`
}
