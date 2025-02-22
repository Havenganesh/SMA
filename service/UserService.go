package service

import (
	"context"
	"fmt"
	"sma/cErrors"
	"sma/db"
	"sma/dto"
	"sma/rpcGateway"
	"sma/validate"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type UserService struct{}

func (u *UserService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	fmt.Println("Create User called : ", req)
	if err := validate.Validate.Struct(req); err != nil {
		return nil, fmt.Errorf("%s : %s", cErrors.INVALID_OR_MISSING_VALUE, err)
	}
	user := dto.User{}
	user.Name = req.Name
	user.PBNumber = req.PBNumber
	user.PhoneNumber = req.PhoneNumber
	user.UserName = req.UserName
	user.Password = req.Password
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	user.CustomerID = uuid.New().String()
	if !isValidServiceType(string(req.ServiceType)) {
		return nil, cErrors.INVALID_SERVICE_TYPE
	}
	if req.ServiceType == "" {
		user.ServiceType = dto.NORMAL
	} else {
		user.ServiceType = req.ServiceType
	}
	user.SearchFilter = append(user.SearchFilter, "PHONENUMBER:"+req.PhoneNumber)
	user.SearchFilter = append(user.SearchFilter, "PBNUMBER:"+req.PBNumber)
	err := db.DB.InsertOne(user)
	if err != nil {
		return nil, err
	}
	return &dto.CreateUserResponse{Status: "success"}, nil
}

func (u *UserService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	if err := validate.Validate.Struct(req); err != nil {
		return nil, fmt.Errorf("%s : %s", cErrors.INVALID_OR_MISSING_VALUE, err)
	}
	user, err := GetUserByUserName(req.UserName)
	if err != nil {
		return nil, err
	}
	if req.Password != user.Password {
		return nil, cErrors.INVALID_PASSWORD
	}

	token, err := rpcGateway.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{Status: "Login successfully", Token: token}, nil
}

func GetUserById(customerID string) (*dto.User, error) {
	f := bson.M{"customerID": customerID}
	user, err := getUserByFilter(f)
	if err != nil || user == nil {
		return nil, fmt.Errorf("%s : %s", cErrors.USER_NOT_FOUND, err)
	}
	return user, nil
}

func GetUserByUserName(name string) (*dto.User, error) {
	f := bson.M{"userName": name}
	user, err := getUserByFilter(f)
	if err != nil || user == nil {
		return nil, fmt.Errorf("%s : %s", cErrors.USER_NOT_FOUND, err)
	}
	return user, nil
}

func getUserByFilter(f bson.M) (*dto.User, error) {
	var user dto.User
	if err := db.DB.FindOne(f, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func isValidServiceType(str string) bool {
	if str == "" {
		return true
	}
	if str == string(dto.NORMAL) || str == string(dto.ELITE) || str == string(dto.PREMIUM) {
		return true
	}
	return false
}
