package controller

import (
	models "be-api/features"
)

type ImageResponse struct {
	ID         uint   `json:"image_id,omitempty"`
	HomestayID uint   `json:"homestay_id,omitempty"`
	Link       string `json:"image_link,omitempty"`
}

func ImageEntityToResponse(image models.ImageEntity) ImageResponse {
	return ImageResponse{
		ID:         image.ID,
		HomestayID: image.HomestayID,
		Link:       image.Link,
	}
}
