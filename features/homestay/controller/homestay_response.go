package controller

import (
	models "be-api/features"
)

func HomestayEntityToResponse(homestay models.HomestayEntity) models.HomestayEntity {
	var reviews []models.ReviewEntity
	for _, review := range homestay.Reviews {
		reviews = append(reviews, ReviewEntityToResponse(review))
	}

	return models.HomestayEntity{
		HostID:      homestay.HostID,
		Title:       homestay.Title,
		Description: homestay.Description,
		Location:    homestay.Location,
		Price:       homestay.Price,
		Facilities:  homestay.Facilities,
		Images:      homestay.Images,
		Reviews:     reviews,
	}
}

func ReviewEntityToResponse(review models.ReviewEntity) models.ReviewEntity {
	return models.ReviewEntity{
		Ratings: review.Ratings,
	}
}
