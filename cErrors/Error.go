package cErrors

import "fmt"

// This is the Custom Error to add New Error
// Extend make New Error in this type
type CError struct {
	message string
}

func (e CError) Error() string {
	return fmt.Sprintf("%s : ", e.message)
}

var DATABASE_CONNECTION_FAILED CError = CError{message: "DATABASE_CONNECTION_FAILED"}
var DATABASE_INSERT_ONE_FAILED CError = CError{message: "DATABASE_INSERT_ONE_FAILED"}
var DATABASE_INSERT_MANY_FAILED CError = CError{message: "DATABASE_INSERT_MANY_FAILED"}
var DATABASE_UPDATE_MANY_FAILED CError = CError{message: "DATABASE_UPDATE_MANY_FAILED"}
var DATABASE_FIND_ALL_FAILED CError = CError{message: "DATABASE_FIND_ALL_FAILED"}
var DATABASE_FIND_ONE_FAILED CError = CError{message: "DATABASE_FIND_ONE_FAILED"}
var DATABASE_AGGREAGATION_FAILED CError = CError{message: "DATABASE_AGGREAGATION_FAILED"}

var USER_NOT_FOUND CError = CError{message: "USER_NOT_FOUND"}

var DOCUMENT_CANNOT_BE_NIL CError = CError{message: "DOCUMENT_CANNOT_BE_NIL"}
var DOCUMENTS_MUST_BE_STRUCT_TYPE CError = CError{message: "DOCUMENTS_MUST_BE_STRUCT_TYPE"}
var DOCUMENTS_MUST_BE_POINTER_TYPE CError = CError{message: "DOCUMENTS_MUST_BE_POINTER_TYPE"}

var INVALID_SERVICE_NAME CError = CError{message: "INVALID_SERVICE_NAME"}
var INVALID_SERVICE_TYPE CError = CError{message: "INVALID_SERVICE_TYPE"}
var SERVICE_CANNOT_BE_NIL CError = CError{message: "INVALID_SERVICE_NAME"}
var INVALID_TOKEN CError = CError{message: "INVALID_TOKEN"}
var INVALID_PASSWORD CError = CError{message: "INVALID_PASSWORD"}
var INVALID_REQUEST CError = CError{message: "INVALID_REQUEST"}
var END_TIME_MUST_BE_AFTER_START_TIME CError = CError{message: "END_TIME_MUST_BE_AFTER_START_TIME"}

var INVALID_OR_MISSING_VALUE CError = CError{message: "INVALID_OR_MISSING_VALUE"}
