package features

import (
	"time"
)

type UserEntity struct {
	ID             uint             `json:"user_id,omitempty" form:"user_id"`
	FirstName      string           `json:"first_name,omitempty" form:"first_name"`
	LastName       string           `json:"last_name,omitempty" form:"last_name"`
	Username       string           `json:"username,omitempty" form:"username"`
	Email          string           `json:"email,omitempty" form:"email"`
	Password       string           `json:"password,omitempty" form:"password"`
	Phone          string           `json:"phone,omitempty" form:"phone"`
	BirthDate      string           `json:"birth_date,omitempty" form:"birth_date"`
	Address        string           `json:"address,omitempty" form:"address"`
	Gender         string           `json:"gender,omitempty" form:"gender"`
	Bio            string           `json:"bio,omitempty" form:"bio"`
	HostingCount   uint             `json:"hosting_count,omitempty" form:"hosting_count"`
	BookingCount   uint             `json:"booking_count,omitempty" form:"booking_count"`
	ProfilePicture string           `json:"profile_picture,omitempty" form:"profile_picture"`
	CreatedAt      time.Time        `json:"created_at,omitempty"`
	UpdatedAt      time.Time        `json:"updated_at,omitempty"`
	DeletedAt      time.Time        `json:"deleted_at,omitempty"`
	Homestays      []HomestayEntity `json:"homestays,omitempty"`
	Bookings       []BookingEntity  `json:"bookings,omitempty"`
	Reviews        []ReviewEntity   `json:"reviews,omitempty"`
}

type HomestayEntity struct {
	ID          uint            `json:"homestay_id,omitempty" form:"homestay_id"`
	HostID      uint            `json:"host_id,omitempty" form:"host_id"`
	Host        UserEntity      `json:"user,omitempty"`
	Title       string          `json:"title,omitempty" form:"title"`
	Description string          `json:"description,omitempty" form:"description"`
	Location    string          `json:"location,omitempty" form:"location"`
	Address     string          `json:"address,omitempty" form:"address"`
	Price       float64         `json:"price,omitempty" form:"price"`
	Facilities  string          `json:"facilities,omitempty" form:"facilities"`
	Rating      float64         `json:"rating,omitempty" form:"rating"`
	CreatedAt   time.Time       `json:"created_at,omitempty"`
	UpdatedAt   time.Time       `json:"updated_at,omitempty"`
	DeletedAt   time.Time       `json:"deleted_at,omitempty"`
	Bookings    []BookingEntity `json:"bookings,omitempty"`
	Reviews     []ReviewEntity  `json:"reviews,omitempty"`
	Images      []ImageEntity   `json:"images,omitempty"`
}

type ImageEntity struct {
	ID         uint           `json:"image_id,omitempty" form:"image_id"`
	HomestayID uint           `json:"homestay_id,omitempty" form:"host_id"`
	Homestay   HomestayEntity `json:"homestay,omitempty"`
	Link       string         `json:"link,omitempty"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty"`
	DeletedAt  time.Time      `json:"deleted_at,omitempty"`
}

type BookingEntity struct {
	ID           uint           `json:"booking_id,omitempty" form:"booking_id"`
	CustomerID   uint           `json:"customer_id,omitempty" form:"customer_id"`
	Customer     UserEntity     `json:"customer,omitempty"`
	HomestayID   uint           `json:"homestay_id,omitempty" form:"homestay_id"`
	Homestay     HomestayEntity `json:"homestay,omitempty"`
	CheckInDate  string         `json:"check_in_date,omitempty" form:"check_in_date"`
	CheckOutdate string         `json:"check_out_date,omitempty" form:"check_out_date"`
	Status       string         `json:"booking_status,omitempty" form:"booking_status"`
	Duration     uint           `json:"duration,omitempty"`
	TotalPrice   float64        `json:"total_price,omitempty" form:"total_price"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	DeletedAt    time.Time      `json:"deleted_at,omitempty"`
}

type ReviewEntity struct {
	ID         uint           `json:"review_id,omitempty" form:"review_id"`
	CustomerID uint           `json:"customer_id,omitempty" form:"customer_id"`
	Customer   UserEntity     `json:"customer,omitempty"`
	HomestayID uint           `json:"homestay_id,omitempty" form:"homestay_id"`
	Homestay   HomestayEntity `json:"homestay,omitempty"`
	Reviews    string         `json:"reviews,omitempty" form:"reviews"`
	Ratings    uint           `json:"ratings,omitempty" form:"ratings"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty"`
	DeletedAt  time.Time      `json:"deleted_at,omitempty"`
}

type PaymentEntity struct {
	ID        uint          `json:"payment_id,omitempty" form:"payment_id"`
	BookingID uint          `json:"booking_id,omitempty" form:"booking_id"`
	Booking   BookingEntity `json:"booking,omitempty"`
	Name      string        `json:"payment_name,omitempty" form:"payment_name"`
	Status    string        `json:"payment_status,omitempty" form:"payment_status"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at,omitempty"`
	DeletedAt time.Time     `json:"deleted_at,omitempty"`
}

type LoginUser struct {
	Email    string `json:"email,omitempty" form:"email" validate:"required,email"`
	Password string `json:"password,omitempty" form:"password" validate:"required"`
}
