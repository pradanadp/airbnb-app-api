package features

func UserModelToEntity(user User) UserEntity {
	var homestayEntities []HomestayEntity
	for _, homestay := range user.Homestays {
		homestayEntities = append(homestayEntities, HomestayModelToEntity(homestay))
	}

	var bookingEntities []BookingEntity
	for _, booking := range user.Bookings {
		bookingEntities = append(bookingEntities, BookingModelToEntity(booking))
	}

	var reviewEntities []ReviewEntity
	for _, review := range user.Reviews {
		reviewEntities = append(reviewEntities, ReviewModelToEntity(review))
	}

	return UserEntity{
		ID:             user.ID,
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
		Homestays:      homestayEntities,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt.Time,
		Bookings:       bookingEntities,
		Reviews:        reviewEntities,
	}
}

func HomestayModelToEntity(homestay Homestay) HomestayEntity {
	var bookingEntities []BookingEntity
	for _, booking := range homestay.Bookings {
		bookingEntities = append(bookingEntities, BookingModelToEntity(booking))
	}

	var reviewEntities []ReviewEntity
	for _, review := range homestay.Reviews {
		reviewEntities = append(reviewEntities, ReviewModelToEntity(review))
	}

	var imageEntities []ImageEntity
	for _, image := range homestay.Images {
		imageEntities = append(imageEntities, ImageModelToEntity(image))
	}

	return HomestayEntity{
		ID:          homestay.ID,
		HostID:      homestay.HostID,
		Host:        UserModelToEntity(homestay.Host),
		Title:       homestay.Title,
		Description: homestay.Description,
		Location:    homestay.Location,
		Address:     homestay.Address,
		Price:       homestay.Price,
		Facilities:  homestay.Facilities,
		CreatedAt:   homestay.CreatedAt,
		UpdatedAt:   homestay.UpdatedAt,
		DeletedAt:   homestay.DeletedAt.Time,
		Bookings:    bookingEntities,
		Reviews:     reviewEntities,
		Images:      imageEntities,
	}
}

func ImageModelToEntity(image Image) ImageEntity {
	return ImageEntity{
		ID:         image.ID,
		HomestayID: image.HomestayID,
		Homestay:   HomestayModelToEntity(image.Homestay),
		Link:       image.Link,
		CreatedAt:  image.CreatedAt,
		UpdatedAt:  image.UpdatedAt,
		DeletedAt:  image.DeletedAt.Time,
	}
}

func BookingModelToEntity(booking Booking) BookingEntity {
	return BookingEntity{
		ID:           booking.ID,
		CustomerID:   booking.CustomerID,
		Customer:     UserModelToEntity(booking.Customer),
		HomestayID:   booking.HomestayID,
		Homestay:     HomestayModelToEntity(booking.Homestay),
		CheckInDate:  booking.CheckInDate,
		CheckOutdate: booking.CheckOutdate,
		Status:       booking.Status,
		Duration:     booking.Duration,
		TotalPrice:   booking.TotalPrice,
		CreatedAt:    booking.CreatedAt,
		UpdatedAt:    booking.UpdatedAt,
		DeletedAt:    booking.DeletedAt.Time,
	}
}

func ReviewModelToEntity(review Review) ReviewEntity {
	return ReviewEntity{
		ID:         review.ID,
		CustomerID: review.CustomerID,
		Customer:   UserModelToEntity(review.Customer),
		HomestayID: review.HomestayID,
		Homestay:   HomestayModelToEntity(review.Homestay),
		Reviews:    review.Reviews,
		Ratings:    review.Ratings,
		CreatedAt:  review.CreatedAt,
		UpdatedAt:  review.UpdatedAt,
		DeletedAt:  review.DeletedAt.Time,
	}
}

func PaymentModelToEntity(payment Payment) PaymentEntity {
	return PaymentEntity{
		ID:        payment.ID,
		BookingID: payment.BookingID,
		Booking:   BookingModelToEntity(payment.Booking),
		Name:      payment.Name,
		Status:    payment.Status,
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
		DeletedAt: payment.DeletedAt.Time,
	}
}
