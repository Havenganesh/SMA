package dto

import "time"

type User struct {
	Name         string       `json:"name,omitempty" bson:"name,omitempty"`
	UserName     string       `json:"userName,omitempty" bson:"userName,omitempty"`
	Password     string       `json:"password,omitempty" bson:"password,omitempty"`
	PBNumber     string       `json:"personalBusinessNumber,omitempty" bson:"personalBusinessNumber,omitempty"`
	PhoneNumber  string       `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	ServiceType  SERVICE_TYPE `json:"serviceType,omitempty" bson:"serviceType,omitempty"`
	CustomerID   string       `json:"customerID,omitempty" bson:"customerID,omitempty"`
	SearchFilter []string     `json:"searchFilter,omitempty" bson:"searchFilter,omitempty"`
	CreatedTime  time.Time    `json:"createdTime,omitempty" bson:"createdTime,omitempty"`
	UpdatedTime  time.Time    `json:"updatedTime,omitempty" bson:"updatedTime,omitempty"`
}

type SERVICE_TYPE string

const (
	NORMAL  SERVICE_TYPE = "NORMAL"
	ELITE   SERVICE_TYPE = "ELITE"
	PREMIUM SERVICE_TYPE = "PREMIUM"
)
