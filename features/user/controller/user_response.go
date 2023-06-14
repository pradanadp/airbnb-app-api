package controller

import "be-api/features"

type ResponseUser struct {
	Id 				 uint `json:"user_id,omitempty"`
	UserName 		 string `json:"user_name,omitempty"`
	FullName		 string `json:"full_name,omitempty"`
	Email    		 string `json:"email,omitempty"`
	Phone    		 string `json:"phone,omitempty"`
	Address  		 string `json:"address,omitempty"`
	ProfilePicture   string `json:"profile_picture,omitempty"`
	Role             string `json:"role,omitempty"`     
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

func EntityToReadResponse(input features.UserEntity)ResponseUser{
	return ResponseUser{
		Id: input.ID,
		UserName: input.Username,
		FullName: input.FullName,
		Email: input.Email,
		Phone: input.Phone,
		Address: input.Address,
		ProfilePicture: input.ProfilePicture,
		Role: input.Role,
	}
}