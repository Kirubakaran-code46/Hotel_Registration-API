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

type PoliciesInfo struct {
	Uid                      string   `json:"uid"`
	Check_in                 string   `json:"check_in"`
	Check_out                string   `json:"check_out"`
	Checkinout_policy        string   `json:"checkinout_policy"`
	CancellationPolicy       string   `json:"cancellationPolicy"`
	Allow_unmarriedCouples   string   `json:"allow_unmarriedCouples"`
	Allow_minor_guest        string   `json:"allow_minor_guest"`
	Allow_onlymale_guests    string   `json:"allow_onlymale_guests"`
	Allow_smoking            string   `json:"allow_smoking"`
	Allow_parties            string   `json:"allow_parties"`
	Allow_invite_guests      string   `json:"allow_invite_guests"`
	Wheelchar_accessible     string   `json:"wheelchar_accessible"`
	Allow_pets               string   `json:"allow_pets"`
	Accepted_proofs          []string `json:"accepted_proofs"`
	Additional_propertyrules string   `json:"additional_propertyrules"`
}
