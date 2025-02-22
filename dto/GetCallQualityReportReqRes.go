package dto

import "time"

type GetCallQualityReportReq struct {
	PhoneNumber string    `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty" validate:"required"`
	StartTime   time.Time `json:"startTime,omitempty" bson:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty" bson:"endTime,omitempty"`
}
type GetCallQualityReportRes struct {
	QualityReport *QualityReport `json:"qualityReport,omitempty" bson:"qualityReport,omitempty"`
}
