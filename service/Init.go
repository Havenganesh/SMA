package service

import (
	"sma/db"
	"sma/dto"
	"sma/rpcGateway"

	"go.mongodb.org/mongo-driver/bson"
)

func Init() {
	rpcGateway.RegisterService("UserService", &UserService{})
	rpcGateway.RegisterService("CallRecordService", &CallRecordService{})
	db.DB.CreateUniqueIndex(&dto.User{}, bson.D{{Key: "userName", Value: 1}})
	db.DB.CreateUniqueIndex(&dto.CallRecord{}, bson.D{{Key: "searchFilter", Value: 1}})

}
