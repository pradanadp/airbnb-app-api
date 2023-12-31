package controller

import (
	models "be-api/features"
)

type HomestayResponse struct {
	ID          uint             `json:"homestay_id,omitempty"`
	HostID      uint             `json:"host_id,omitempty"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	Location    string           `json:"location,omitempty"`
	Address     string           `json:"address,omitempty"`
	Price       float64          `json:"price,omitempty"`
	Facilities  string           `json:"facilities,omitempty"`
	Rating      float64          `json:"rating"`
	Images      []ImageResponse  `json:"image_link,omitempty"`
	Reviews     []ReviewResponse `json:"reviews,omitempty"`
}

type ReviewResponse struct {
	CustomerID uint    `json:"customer_id,omitempty"`
	Reviews    string  `json:"reviews,omitempty"`
	Ratings    float64 `json:"ratings,omitempty"`
}

type ImageResponse struct {
	Link string `json:"image_link,omitempty"`
}

func HomestayEntityToResponse(homestay models.HomestayEntity) HomestayResponse {
	var reviews []ReviewResponse
	for _, review := range homestay.Reviews {
		reviews = append(reviews, ReviewEntityToResponse(review))
	}

	var imageLinks []ImageResponse
	for _, link := range homestay.Images {
		imageLinks = append(imageLinks, ImageEntityToResponse(link))
	}

	return HomestayResponse{
		ID:          homestay.ID,
		HostID:      homestay.HostID,
		Title:       homestay.Title,
		Description: homestay.Description,
		Location:    homestay.Location,
		Price:       homestay.Price,
		Facilities:  homestay.Facilities,
		Rating:      homestay.Rating,
		Images:      imageLinks,
		Reviews:     reviews,
	}
}

func ReadAllHomestayEntityToResponse(homestay models.HomestayEntity) HomestayResponse {
	var imageLinks []ImageResponse
	for _, link := range homestay.Images {
		imageLinks = append(imageLinks, ImageEntityToResponse(link))
	}

	return HomestayResponse{
		ID:          homestay.ID,
		HostID:      homestay.HostID,
		Title:       homestay.Title,
		Description: homestay.Description,
		Location:    homestay.Location,
		Price:       homestay.Price,
		Facilities:  homestay.Facilities,
		Rating:      homestay.Rating,
		Images:      imageLinks,
	}
}

func ReadAllHomestayByHostIDEntityToResponse(homestay models.HomestayEntity) HomestayResponse {
	var imageLinks []ImageResponse
	for _, link := range homestay.Images {
		imageLinks = append(imageLinks, ImageEntityToResponse(link))
	}

	var reviews []ReviewResponse
	for _, review := range homestay.Reviews {
		reviews = append(reviews, ReviewEntityToResponse(review))
	}

	return HomestayResponse{
		ID:          homestay.ID,
		HostID:      homestay.HostID,
		Title:       homestay.Title,
		Description: homestay.Description,
		Location:    homestay.Location,
		Price:       homestay.Price,
		Facilities:  homestay.Facilities,
		Rating:      homestay.Rating,
		Images:      imageLinks,
		Reviews:     reviews,
	}
}

func ReviewEntityToResponse(review models.ReviewEntity) ReviewResponse {
	return ReviewResponse{
		CustomerID: review.CustomerID,
		Reviews:    review.Reviews,
		Ratings:    review.Ratings,
	}
}

func ImageEntityToResponse(image models.ImageEntity) ImageResponse {
	return ImageResponse{
		Link: image.Link,
	}
}
