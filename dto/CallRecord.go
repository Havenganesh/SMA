package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Call Record Struct
type CallRecord struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Caller        *Contact           `json:"caller,omitempty" bson:"caller,omitempty"`
	Receiver      *Contact           `json:"receiver,omitempty" bson:"receiver,omitempty"`
	Duration      int64              `json:"duration" bson:"duration"`
	CallType      string             `json:"callType,omitempty" bson:"callType,omitempty"`
	Cost          float64            `json:"cost" bson:"cost"`
	StartTime     *time.Time         `json:"startTime,omitempty" bson:"startTime,omitempty"`
	EndTime       *time.Time         `json:"endTime,omitempty" bson:"endTime,omitempty"`
	Quality       string             `json:"quality,omitempty" bson:"quality,omitempty"`
	NetworkIssues bool               `json:"networkIssues" bson:"networkIssues"`
	SearchFilter  []string           `json:"searchFilter" bson:"searchFilter"`
}

type Contact struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	Prefix      string `json:"prefix,omitempty" bson:"prefix,omitempty"`
}
