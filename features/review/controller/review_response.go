package controller

import "be-api/features"

type ResponseReviews struct {
	Id         uint 			`json:"id,omitempty"`
	HomestayID uint				`json:"homestay_id,omitempty"`
	Reviews    string			`json:"reviews,omitempty"`
	Ratings    float64			`json:"ratings,omitempty"`
	Homestay   ResponseHomestay
	User	   ResponseUser	
}

type ResponseUser struct {
	UserName string				`json:"username,omitempty"`
}
type ResponseHomestay struct {
	Title string				`json:"title_vila,omitempty"`
}

func EntityToResponseUser(input features.UserEntity) ResponseUser{
	return ResponseUser{
		UserName: input.Username,
	}}

func EntityToResponseHomestay(input features.HomestayEntity) ResponseHomestay{
	return ResponseHomestay{
		Title: input.Title,
	}}
func EntityToResponse(input features.ReviewEntity) ResponseReviews{
	return ResponseReviews{
		Id: input.ID,
		HomestayID: input.HomestayID,
		Reviews: input.Reviews,
		Ratings: input.Ratings,
		Homestay: EntityToResponseHomestay(input.Homestay),
		User: EntityToResponseUser(input.Customer),	
	}
}