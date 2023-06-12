package features

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName      string     `gorm:"column:first_name;not null"`
	LastName       string     `gorm:"column:last_name;not null"`
	Username       string     `gorm:"column:username;not null"`
	Email          string     `gorm:"column:email;unique;not null"`
	Password       string     `gorm:"column:password;not null"`
	Phone          string     `gorm:"column:phone;unique;not null"`
	BirthDate      string     `gorm:"column:birth_date;not null"`
	Address        string     `gorm:"column:address;not null"`
	Gender         string     `gorm:"type:enum('male','female');default:'male';column:gender;not null"`
	Bio            string     `gorm:"column:bio;not null"`
	HostingCount   uint       `gorm:"column:hosting_count"`
	BookingCount   uint       `gorm:"column:booking_count"`
	ProfilePicture string     `gorm:"column:profile_picture"`
	Homestays      []Homestay `gorm:"foreignKey:HostID"`
	Bookings       []Booking  `gorm:"foreignKey:CustomerID"`
	Reviews        []Review   `gorm:"foreignKey:CustomerID"`
}

type Homestay struct {
	gorm.Model
	HostID      uint      `gorm:"column:host_id;not null"`
	Host        User      `gorm:"foreignKey:HostID"`
	Title       string    `gorm:"column:title;not null"`
	Description string    `gorm:"column:description;not null"`
	Location    string    `gorm:"column:location;not null"`
	Price       float64   `gorm:"column:price;not null"`
	Facilities  string    `gorm:"column:facilities;not null"`
	Images      string    `gorm:"column:images_links;not null"`
	Bookings    []Booking `gorm:"foreignKey:HomestayID"`
	Reviews     []Review  `gorm:"foreignKey:HomestayID"`
}

type Booking struct {
	gorm.Model
	CustomerID   uint     `gorm:"column:customer_id;not null"`
	Customer     User     `gorm:"foreignKey:CustomerID"`
	HomestayID   uint     `gorm:"column:homestay_id;not null"`
	Homestay     Homestay `gorm:"foreignKey:HomestayID"`
	PaymentID    uint     `gorm:"column:payment_id;not null"`
	Payment      Payment  `gorm:"foreignKey:PaymentID"`
	CheckInDate  string   `gorm:"column:check_in_date;not null"`
	CheckOutdate string   `gorm:"column:check_out_date;not null"`
	Status       string   `gorm:"type:enum('available','reserved','booked');default:'available';column:booking_status;not null"`
	Duration     uint     `gorm:"column:duration;not null"`
	TotalPrice   float64  `gorm:"column:total_price;not null"`
}

type Review struct {
	gorm.Model
	CustomerID uint     `gorm:"column:customer_id;not null"`
	Customer   User     `gorm:"foreignKey:CustomerID"`
	HomestayID uint     `gorm:"column:homestay_id;not null"`
	Homestay   Homestay `gorm:"foreignKey:HomestayID"`
	Reviews    string   `gorm:"column:reviews;not null"`
	Ratings    uint     `gorm:"column:ratings;not null"`
}

type Payment struct {
	gorm.Model
	Name     string    `gorm:"column:payment_name;not null"`
	Status   string    `gorm:"column:payment_status;not null"`
	Bookings []Booking `gorm:"foreignKey:PaymentID"`
}
