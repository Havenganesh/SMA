package dto

type BillReport struct {
	CallerName     string  `json:"callerName,omitempty" bson:"callerName,omitempty"`
	ReceiverName   string  `json:"receiverName,omitempty" bson:"receiverName,omitempty"`
	CallerNumber   string  `json:"callerNumber,omitempty" bson:"callerNumber,omitempty"`
	ReceiverNumber string  `json:"receiverNumber,omitempty" bson:"receiverNumber,omitempty"`
	Duration       int64   `json:"duration,omitempty" bson:"duration,omitempty"`
	Cost           float64 `json:"cost,omitempty" bson:"cost,omitempty"`
}
