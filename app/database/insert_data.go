package database

import (
	models "be-api/features"
	"be-api/utils"

	"gorm.io/gorm"
)

func InitUsersData(db *gorm.DB) {
	users := []models.User{
		{
			FirstName:      "John",
			LastName:       "Smith",
			Username:       "johnsmith",
			Email:          "johnsmith@example.com",
			Password:       "password123",
			Phone:          "1234567890",
			BirthDate:      "1990-01-10",
			Address:        "123 Main St, City",
			Gender:         "Male",
			Bio:            "I'm John Smith, a software engineer.",
			HostingCount:   10,
			BookingCount:   15,
			ProfilePicture: "https://example.com/johnsmith.jpg",
		},
		{
			FirstName:      "Emily",
			LastName:       "Johnson",
			Username:       "emilyjohnson",
			Email:          "emilyjohnson@example.com",
			Password:       "password456",
			Phone:          "9876543210",
			BirthDate:      "1995-03-15",
			Address:        "456 Park Ave, Town",
			Gender:         "Female",
			Bio:            "I'm Emily Johnson, a graphic designer.",
			HostingCount:   5,
			BookingCount:   20,
			ProfilePicture: "https://example.com/emilyjohnson.jpg",
		},
		{
			FirstName:      "Michael",
			LastName:       "Davis",
			Username:       "michaeldavis",
			Email:          "michaeldavis@example.com",
			Password:       "password789",
			Phone:          "5678901234",
			BirthDate:      "1988-06-20",
			Address:        "789 Elm St, Village",
			Gender:         "Male",
			Bio:            "I'm Michael Davis, a marketing specialist.",
			HostingCount:   2,
			BookingCount:   8,
			ProfilePicture: "https://example.com/michaeldavis.jpg",
		},
		{
			FirstName:      "Jessica",
			LastName:       "Wilson",
			Username:       "jessicawilson",
			Email:          "jessicawilson@example.com",
			Password:       "passwordabc",
			Phone:          "4321098765",
			BirthDate:      "1992-09-05",
			Address:        "987 Oak Ln, County",
			Gender:         "Female",
			Bio:            "I'm Jessica Wilson, a teacher.",
			HostingCount:   3,
			BookingCount:   12,
			ProfilePicture: "https://example.com/jessicawilson.jpg",
		},
		{
			FirstName:      "Christopher",
			LastName:       "Brown",
			Username:       "christopherbrown",
			Email:          "christopherbrown@example.com",
			Password:       "passwordxyz",
			Phone:          "3456789012",
			BirthDate:      "1993-11-12",
			Address:        "654 Pine Rd, Hamlet",
			Gender:         "Male",
			Bio:            "I'm Christopher Brown, a musician.",
			HostingCount:   7,
			BookingCount:   18,
			ProfilePicture: "https://example.com/christopher.jpg",
		},
	}

	for _, user := range users {
		user.Password, _ = utils.HashPasword(user.Password)
		db.Create(&user)
	}
}

func InitHomestaysData(db *gorm.DB) {
	homestays := []models.Homestay{
		{
			HostID:      80,
			Title:       "Beachfront Villa",
			Description: "Luxurious villa with stunning ocean views",
			Location:    "Seaside Resort",
			Address:     "456 Ocean Drive",
			Price:       500.00,
			Facilities:  "Swimming Pool, Private Beach Access, Air Conditioning",
		},
		{
			HostID:      80,
			Title:       "Mountain Cabin Retreat",
			Description: "Rustic cabin surrounded by scenic mountains",
			Location:    "Mountain Village",
			Address:     "789 Mountain Road",
			Price:       200.00,
			Facilities:  "Fireplace, Hiking Trails Nearby, BBQ Area",
		},
		{
			HostID:      81,
			Title:       "Cozy City Apartment",
			Description: "Modern apartment in the heart of the city",
			Location:    "City Center",
			Address:     "123 Main Street",
			Price:       150.00,
			Facilities:  "Wi-Fi, Gym, Concierge Service",
		},
		{
			HostID:      82,
			Title:       "Lakeside Cottage",
			Description: "Charming cottage overlooking a peaceful lake",
			Location:    "Lakefront",
			Address:     "987 Lakeview Lane",
			Price:       300.00,
			Facilities:  "Boat Dock, BBQ Area, Fishing Gear",
		},
		{
			HostID:      82,
			Title:       "Countryside Farmhouse",
			Description: "Quaint farmhouse surrounded by scenic fields",
			Location:    "Rural Area",
			Address:     "456 Farm Road",
			Price:       250.00,
			Facilities:  "Garden, Animal Petting, Farm-to-Table Experience",
		},
		{
			HostID:      82,
			Title:       "Luxury Penthouse Suite",
			Description: "Opulent penthouse with breathtaking city views",
			Location:    "Upscale District",
			Address:     "789 Skyline Avenue",
			Price:       800.00,
			Facilities:  "Infinity Pool, Private Elevator, 24/7 Butler Service",
		},
		{
			HostID:      83,
			Title:       "Beachfront Bungalow",
			Description: "Quaint bungalow steps away from the sandy beach",
			Location:    "Coastal Village",
			Address:     "321 Beach Road",
			Price:       180.00,
			Facilities:  "Outdoor Seating, Hammocks, Surfing Lessons",
		},
	}

	for _, homestay := range homestays {
		db.Create(&homestay)
	}
}

func InitReviewsData(db *gorm.DB) {
	reviews := []models.Review{
		{
			CustomerID: 1,
			HomestayID: 1,
			Reviews:    "Great place to stay! The host was very accommodating.",
			Ratings:    4.5,
		},
		{
			CustomerID: 2,
			HomestayID: 1,
			Reviews:    "Beautiful homestay with amazing amenities.",
			Ratings:    4.8,
		},
		{
			CustomerID: 3,
			HomestayID: 1,
			Reviews:    "Clean and comfortable rooms. Highly recommended!",
			Ratings:    4.7,
		},
		{
			CustomerID: 4,
			HomestayID: 1,
			Reviews:    "Friendly staff and excellent service. Will definitely come back.",
			Ratings:    4.6,
		},
		{
			CustomerID: 5,
			HomestayID: 1,
			Reviews:    "Lovely ambiance and peaceful surroundings. Perfect for a relaxing getaway.",
			Ratings:    4.4,
		},
		{
			CustomerID: 1,
			HomestayID: 2,
			Reviews:    "Absolutely loved our stay! The views from the homestay were breathtaking.",
			Ratings:    4.7,
		},
		{
			CustomerID: 2,
			HomestayID: 2,
			Reviews:    "Cozy and well-maintained rooms. The host was very friendly and helpful.",
			Ratings:    4.5,
		},
		{
			CustomerID: 3,
			HomestayID: 2,
			Reviews:    "Great location near the beach. Ideal for a beach vacation.",
			Ratings:    4.6,
		},
		{
			CustomerID: 4,
			HomestayID: 2,
			Reviews:    "The homestay exceeded our expectations. We had a wonderful time!",
			Ratings:    4.8,
		},
		{
			CustomerID: 5,
			HomestayID: 2,
			Reviews:    "Highly recommend this homestay. The staff was amazing!",
			Ratings:    4.9,
		},
		{
			CustomerID: 1,
			HomestayID: 3,
			Reviews:    "Peaceful and serene atmosphere. Perfect for a weekend getaway.",
			Ratings:    4.4,
		},
		{
			CustomerID: 2,
			HomestayID: 3,
			Reviews:    "The homestay had all the necessary amenities. We had a comfortable stay.",
			Ratings:    4.6,
		},
		{
			CustomerID: 3,
			HomestayID: 3,
			Reviews:    "The host was very accommodating and made us feel welcome.",
			Ratings:    4.5,
		},
		{
			CustomerID: 4,
			HomestayID: 3,
			Reviews:    "Beautiful surroundings and well-maintained property.",
			Ratings:    4.7,
		},
		{
			CustomerID: 5,
			HomestayID: 3,
			Reviews:    "We thoroughly enjoyed our stay. Would love to visit again.",
			Ratings:    4.3,
		},
		{
			CustomerID: 1,
			HomestayID: 4,
			Reviews:    "Stunning views and luxurious accommodations. Worth every penny!",
			Ratings:    4.9,
		},
		{
			CustomerID: 2,
			HomestayID: 4,
			Reviews:    "The homestay was spacious and elegantly designed. Highly recommended.",
			Ratings:    4.8,
		},
		{
			CustomerID: 3,
			HomestayID: 4,
			Reviews:    "The host provided exceptional service throughout our stay.",
			Ratings:    4.7,
		},
		{
			CustomerID: 4,
			HomestayID: 4,
			Reviews:    "We had a memorable experience at this homestay. Can't wait to return.",
			Ratings:    4.6,
		},
		{
			CustomerID: 5,
			HomestayID: 4,
			Reviews:    "The amenities exceeded our expectations. 5-star experience!",
			Ratings:    4.9,
		},
		{
			CustomerID: 1,
			HomestayID: 5,
			Reviews:    "Charming and cozy homestay. The host was very attentive.",
			Ratings:    4.5,
		},
		{
			CustomerID: 2,
			HomestayID: 5,
			Reviews:    "We had a wonderful time at this homestay. The location was perfect.",
			Ratings:    4.7,
		},
		{
			CustomerID: 3,
			HomestayID: 5,
			Reviews:    "The rooms were clean and comfortable. Excellent value for money.",
			Ratings:    4.6,
		},
		{
			CustomerID: 4,
			HomestayID: 5,
			Reviews:    "Friendly staff and great hospitality. Would definitely recommend.",
			Ratings:    4.4,
		},
		{
			CustomerID: 5,
			HomestayID: 5,
			Reviews:    "The homestay had a unique and artistic ambiance. Loved our stay!",
			Ratings:    4.8,
		},
	}

	for _, review := range reviews {
		db.Create(&review)
	}
}
