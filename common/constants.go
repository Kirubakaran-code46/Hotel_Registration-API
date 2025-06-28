package common

const (
	SUCCESSCODE   = "S"
	ERRORCODE     = "E"
	YES           = "Y"
	INVALID       = "I"
	NO            = "N"
	EMAIL         = "EMAIL"
	MOBILE        = "MOBILE"
	UIDCOOKIENAME = "client_id"
)

type BasicDetailsStruct struct {
	Uid                string `json:"uid"`
	HotelName          string `json:"hotelName"`
	PropertyType       string `json:"propertyType"`
	StarCategory       string `json:"starCategory"`
	YearOfConstruction string `json:"yearOfConstruction"`
	MobileCode         string `json:"mobileCode"`
	PrimaryMobile      string `json:"primaryMobile"`
	SecondaryMobile    string `json:"secondaryMobile"`
	Email              string `json:"email"`
	ChannelManager     string `json:"channelManager"`
}

type LocationDetailsStruct struct {
	Uid       string `json:"uid"`
	AddrLine1 string `json:"addrLine1"`
	AddrLine2 string `json:"addrLine2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zipcode   string `json:"zipcode"`
}

type RoomType struct {
	RoomType         string   `json:"roomType"`
	NoOfRooms        string   `json:"noOfRooms"`
	RoomView         string   `json:"roomView"`
	RoomSizeUnit     string   `json:"roomSizeUnit"`
	RoomSize         string   `json:"roomSize"`
	MaximumOccupancy string   `json:"maximumOccupancy"`
	ExtraBed         string   `json:"extraBed"`
	ExtraPersons     string   `json:"extraPersons"`
	SingleGuestPrice string   `json:"singleGuestPrice"`
	DoubleGuestPrice string   `json:"doubleGuestPrice"`
	TripleGuestPrice string   `json:"tripleGuestPrice"`
	ExtraAdultCharge string   `json:"extraAdultCharge"`
	ChildCharge      string   `json:"childCharge"`
	BelowChildCharge string   `json:"belowChildCharge"`
	RoomAmenities    []string `json:"roomAmenities"`
	SmokingPolicy    string   `json:"smokingPolicy"`
}

type MealsInfo struct {
	Uid                     string   `json:"uid"`
	IsOperationalRestaurant string   `json:"isOperationalRestaurant"`
	MealPackage             string   `json:"mealPackage"`
	TypesOfMeals            []string `json:"typesOfMeals"`
	MealRackPrice           string   `json:"mealRackPrice"`
}

type AvailabilityInfo struct {
	Uid       string `json:"uid"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
