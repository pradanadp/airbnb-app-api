package controller

import "be-api/features"

type Response struct {
	Id         uint
	HomestayID uint
	Reviews    string
	Ratings    float64
}

func EntityToResponse(input features.ReviewEntity) Response{
	return Response{
		Id: input.ID,
		HomestayID: input.HomestayID,
		Reviews: input.Reviews,
		Ratings: input.Ratings,
	}
}