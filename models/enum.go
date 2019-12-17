package models

type ResponseType int

const (
	EmployeeType ResponseType = iota + 1
	EmployeesType
	ReviewType
	ReviewsType
)
