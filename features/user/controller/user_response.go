package controller

import "be-api/features"

type ResponseUser struct {
	UserName string
	Email    string
	Phone    string
	Address  string
	ProfilePicture   string
}

func EntityToResponse(input features.UserEntity)ResponseUser{
	return ResponseUser{
		UserName: input.Username,
		Email: input.Email,
		Phone: input.Phone,
		Address: input.Address,
		ProfilePicture: input.ProfilePicture,
	}
}