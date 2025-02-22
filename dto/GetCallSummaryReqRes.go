package dto

import "time"

type GetCallSummaryRequest struct {
	PhoneNumber string    `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty" validate:"required"`
	StartTime   time.Time `json:"startTime" bson:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime" bson:"endTime,omitempty"`
}
type GetCallSummaryResponse struct {
	CallSummary *CallSummary `json:"callSummary,omitempty" bson:"callSummary,omitempty"`
}
