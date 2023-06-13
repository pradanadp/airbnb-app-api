package controller

import (
	models "be-api/features"
)

type HomestayResponse struct {
	HostID      uint             `json:"host_id,omitempty"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	Location    string           `json:"location,omitempty"`
	Address     string           `json:"address,omitempty"`
	Price       float64          `json:"price,omitempty"`
	Facilities  string           `json:"facilities,omitempty"`
	Rating      float64          `json:"rating"`
	Images      string           `json:"image_links,omitempty"`
	Reviews     []ReviewResponse `json:"reviews,omitempty"`
}

type ReviewResponse struct {
	Ratings float64 `json:"ratings,omitempty"`
}

func HomestayEntityToResponse(homestay models.HomestayEntity) HomestayResponse {
	var reviews []ReviewResponse
	for _, review := range homestay.Reviews {
		reviews = append(reviews, ReviewEntityToResponse(review))
	}

	return HomestayResponse{
		HostID:      homestay.HostID,
		Title:       homestay.Title,
		Description: homestay.Description,
		Location:    homestay.Location,
		Price:       homestay.Price,
		Facilities:  homestay.Facilities,
		Rating:      homestay.Rating,
		// Reviews:     reviews,
	}
}

func ReviewEntityToResponse(review models.ReviewEntity) ReviewResponse {
	return ReviewResponse{
		Ratings: review.Ratings,
	}
}
