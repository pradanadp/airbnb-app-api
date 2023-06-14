package features

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string     `gorm:"column:first_name;not nul"`
	LastName       string     `gorm:"column:last_name;not nul"`
	Username       string     `gorm:"column:username;unique;not nul"`
	Email          string     `gorm:"column:email;unique;not nul"`
	Password       string     `gorm:"column:password;not nul"`
	Phone          string     `gorm:"column:phone;unique;not nul"`
	BirthDate      string     `gorm:"column:birth_date;not nul"`
	Address        string     `gorm:"column:address;not nul"`
	Gender         string     `gorm:"type:enum('male','female');default:'male';column:gender;not nul"`
	Role           string     `gorm:"type:enum('hoster','user');default:'user';column:role;not nul"`
	Bio            string     `gorm:"column:bio"`
	NamaPemilik    string     `gorm:"column:nama_pemilik"`
	KTPFile        string     `gorm:"column:ktp_file"`
	NIBFile        string     `gorm:"column:nib_file"`
	HostingCount   uint       `gorm:"column:hosting_count"`
	BookingCount   uint       `gorm:"column:booking_count"`
	HostDocument   string     `gorm:"column:host_document"`
	ProfilePicture string     `gorm:"column:profile_picture"`
	Homestays      []Homestay `gorm:"foreignKey:HostID"`
	Bookings       []Booking  `gorm:"foreignKey:CustomerID"`
	Reviews        []Review   `gorm:"foreignKey:CustomerID"`
}

type Homestay struct {
	gorm.Model
	HostID      uint      `gorm:"column:host_id;not null"`
	Host        User      `gorm:"foreignKey:HostID"`
	Title       string    `gorm:"column:title;unique;not null"`
	Description string    `gorm:"column:description;not null"`
	Location    string    `gorm:"column:location;not null"`
	Address     string    `gorm:"column:address;not null"`
	Price       float64   `gorm:"column:price;not null"`
	Facilities  string    `gorm:"column:facilities;not null"`
	Rating      float64   `gorm:"column:rating"`
	Bookings    []Booking `gorm:"foreignKey:HomestayID"`
	Reviews     []Review  `gorm:"foreignKey:HomestayID"`
	Images      []Image   `gorm:"foreignKey:HomestayID"`
}

type Image struct {
	gorm.Model
	HomestayID uint     `gorm:"column:homestay_id;not null"`
	Homestay   Homestay `gorm:"foreignKey:HomestayID"`
	Link       string   `gorm:"column:link;not null"`
}

type Booking struct {
	gorm.Model
	OrderID      string   `gorm:"column:order_id;not null"`
	CustomerID   uint     `gorm:"column:customer_id;not null"`
	Customer     User     `gorm:"foreignKey:CustomerID"`
	HomestayID   uint     `gorm:"column:homestay_id;not null"`
	Homestay     Homestay `gorm:"foreignKey:HomestayID"`
	CheckInDate  string   `gorm:"column:check_in_date;not null"`
	CheckOutDate string   `gorm:"column:check_out_date;not null"`
	Status       string   `gorm:"type:enum('reserved','booked');column:booking_status;not null"`
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
	Ratings    float64  `gorm:"column:ratings;not null"`
}

type Payment struct {
	gorm.Model
	BookingID uint    `gorm:"column:booking_id;unique;not null"`
	Booking   Booking `gorm:"foreignKey:BookingID"`
	Name      string  `gorm:"column:payment_name;not null"`
	Status    string  `gorm:"column:payment_status;not null"`
}
