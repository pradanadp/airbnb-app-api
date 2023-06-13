package controller

import "be-api/features"

type ResponseUser struct {
	Id 				 uint
	UserName 		 string
	Email    		 string
	Phone    		 string
	Address  		 string
	ProfilePicture   string
	Role             string      
}


func EntityToResponse(input features.UserEntity)ResponseUser{
	return ResponseUser{
		Id: input.ID,
		UserName: input.Username,
		Email: input.Email,
		Phone: input.Phone,
		Address: input.Address,
		ProfilePicture: input.ProfilePicture,
		Role: input.Role,
	}
}