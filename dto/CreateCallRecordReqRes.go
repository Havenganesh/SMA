package dto

import "time"

type CreateCallRecordRequest struct {
	Caller        *Contact   `json:"caller,omitempty" bson:"caller,omitempty" validate:"required"`
	Receiver      *Contact   `json:"receiver,omitempty" bson:"receiver,omitempty" validate:"required"`
	Duration      int64      `json:"duration" bson:"duration"`
	CallType      string     `json:"callType,omitempty" bson:"callType,omitempty" validate:"required"`
	StartTime     *time.Time `json:"startTime,omitempty" bson:"startTime,omitempty" validate:"required"`
	EndTime       *time.Time `json:"endTime,omitempty" bson:"endTime,omitempty"`
	Quality       string     `json:"quality,omitempty" bson:"quality,omitempty" validate:"required"`
	NetworkIssues bool       `json:"networkIssues" bson:"networkIssues"`
}

type CreateCallRecordResponse struct {
	Message string `json:"message" bson:"message"`
}
