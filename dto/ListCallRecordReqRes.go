package dto

import "time"

type ListCallRecordRequest struct {
	CallerName   string    `json:"callerName,omitempty" bson:"callerName,omitempty"`
	ReceiverName string    `json:"receiverName,omitempty" bson:"receiverName,omitempty"`
	PhoneNumber  string    `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	StartTime    time.Time `json:"startTime" bson:"startTime"`
	EndTime      time.Time `json:"endTime" bson:"endTime"`
}

type ListCallRecordResponse struct {
	Records []*CallRecord `json:"records" bson:"records"`
	Message string        `json:"message" bson:"message"`
}
