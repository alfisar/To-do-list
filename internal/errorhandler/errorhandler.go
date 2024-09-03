package errorhandler

import "todolist/domain"

const (
	// code series 1xx for validation data

	// Error code for invalid input ex : email format not valid, type data not valid
	ErrCodeInvalidInput int = 101

	// Error code for data required ex: phone is required but phone is empty
	ErrCodeRequired int = 102

	// code series 2xx for error DB

	// Error code for invalid connection DB
	ErrCodeConnection int = 201

	// Error code for data is empty or not found
	ErrCodeDataNotFound int = 202

	// Error code for save data
	ErrCodeInsert int = 203

	// Error code for update data
	ErrCodeUpdate int = 204

	// Error code for delete data
	ErrCodeDelete int = 205

	// Error code for get data
	ErrCodeGet int = 206

	// code series 3xx for error auth

	// Error code for authenticate is not valid ex: token not valid
	ErrCodeInvalidAuth int = 301

	// Error code for authenticate is expired ex: token is expired
	ErrCodeExpSession int = 302

	// Error code for forbiddes akses auth ex: user not have akses
	ErrCodeForbiddenAccess int = 303

	// code series 4xx for error bisnis

	// Error code for invalid logic bisnis ex: user duplicate
	ErrCodeInvalidLogicBisnis int = 401

	// code series 5xx for error internal ex: error hashing

	// Error code for hashing data
	ErrCodeHashing int = 501

	// Error code for generate token
	ErrCodeGenerateToken int = 502

	// Error message for data exist
	ErrMsgDataExist string = "Data already exist"

	// Error message connection empty
	ErrMsgConnEmpty string = "Connection is nil"

	// Error message for data Login
	ErrMsgLoginRequired string = "Username / Password Cannot Empty"

	// Error message for data Login not match
	ErrMsgLoginFailed string = "Pastikan Username / Password Benar"
)

func ErrValidation(err error) (result domain.ErrorData) {
	result.Status = "error"
	result.Code = ErrCodeInvalidInput
	result.Message = "Invalid data input"
	result.Errors = err.Error()
	return
}

func ErrRecordNotFound() (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:  "error",
		Code:    ErrCodeDataNotFound,
		Message: "Data not found",
	}

	return
}

func ErrGetData(err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:  "error",
		Code:    ErrCodeGet,
		Message: "Failed get data",
		Errors:  err.Error(),
	}

	return
}

func ErrInsertData(err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:  "error",
		Code:    ErrCodeInsert,
		Message: "Failed insert data",
		Errors:  err.Error(),
	}

	return
}

func ErrInvalidLogic(code int, message string) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:  "error",
		Code:    code,
		Message: message,
	}

	return
}

func ErrHashing(err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:  "error",
		Code:    ErrCodeHashing,
		Message: "Internal errors",
		Errors:  err.Error(),
	}

	return
}

func ErrInternal(code int, err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:  "error",
		Code:    code,
		Message: "Internal errors",
		Errors:  err.Error(),
	}

	return
}
