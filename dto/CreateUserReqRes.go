package dto

type CreateUserRequest struct {
	Name        string       `json:"name"  validate:"required,min=3"`
	PBNumber    string       `json:"personalBusinessNumber"  validate:"required,len=12,numeric"`
	PhoneNumber string       `json:"phoneNumber"  validate:"required,len=10,numeric"`
	ServiceType SERVICE_TYPE `json:"serviceType"`
	UserName    string       `json:"userName"  validate:"required,len=5"`
	Password    string       `json:"password" validate:"required,len=10"`
}

type CreateUserResponse struct {
	Status     string `json:"status"`
	CustomerID string `json:"customerID"`
}
