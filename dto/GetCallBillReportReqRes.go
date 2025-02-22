package dto

import "time"

type GetCallBillReportReq struct {
	PhoneNumber string    `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty" validate:"required"`
	StartTime   time.Time `json:"startTime,omitempty" bson:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty" bson:"endTime,omitempty"`
}

type GetCallBillReportRes struct {
	BillReport    []*BillReport `json:"billReport" bson:"billReport"`
	TotalCost     float64       `json:"totalCost,omitempty" bson:"totalCost,omitempty"`
	TotalDuration int64         `json:"totalDuration,omitempty" bson:"totalDuration,omitempty"`
}
