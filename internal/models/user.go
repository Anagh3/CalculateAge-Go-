package models

type UserInput struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}
