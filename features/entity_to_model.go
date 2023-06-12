package features

func UserEntityToModel(user UserEntity) User {
	return User{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		Email:          user.Email,
		Password:       user.Password,
		Phone:          user.Phone,
		BirthDate:      user.BirthDate,
		Address:        user.Address,
		Gender:         user.Gender,
		Bio:            user.Bio,
		HostingCount:   user.HostingCount,
		BookingCount:   user.BookingCount,
		ProfilePicture: user.ProfilePicture,
	}
}

func HomestayEntityToModel(homestay HomestayEntity) Homestay {
	return Homestay{
		HostID:      homestay.HostID,
		Title:       homestay.Title,
		Description: homestay.Description,
		Location:    homestay.Location,
		Address:     homestay.Address,
		Price:       homestay.Price,
		Facilities:  homestay.Facilities,
		Rating:      homestay.Rating,
	}
}

func ImageEntityToModel(image ImageEntity) Image {
	return Image{
		HomestayID: image.HomestayID,
		Link:       image.Link,
	}
}

func BookingEntityToModel(booking BookingEntity) Booking {
	return Booking{
		CustomerID:   booking.CustomerID,
		HomestayID:   booking.HomestayID,
		PaymentID:    booking.PaymentID,
		CheckInDate:  booking.CheckInDate,
		CheckOutdate: booking.CheckOutdate,
		Status:       booking.Status,
		Duration:     booking.Duration,
		TotalPrice:   booking.TotalPrice,
	}
}

func ReviewEntityToModel(review ReviewEntity) Review {
	return Review{
		CustomerID: review.CustomerID,
		HomestayID: review.HomestayID,
		Reviews:    review.Reviews,
		Ratings:    review.Ratings,
	}
}

func PaymentEntityToModel(payment PaymentEntity) Payment {
	return Payment{
		Name:   payment.Name,
		Status: payment.Status,
	}
}
