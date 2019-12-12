package apis

import "github.com/yhsiang/review360/models"

type ResponseData interface {
	Type() models.ResponseType
}

type ResponseStatus bool

type DataResponse struct {
	Status ResponseStatus `json:"status"`
	Data   ResponseData   `json:"data"`
}

type StatusResponse struct {
	Status ResponseStatus `json:"status"`
}

type ErrorResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
}
