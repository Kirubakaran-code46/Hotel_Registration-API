package getuserdetails

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Response struct {
	Status           string                       `json:"status"`
	ErrMsg           string                       `json:"errMsg"`
	BasicInfo        common.BasicDetailsStruct    `json:"basicInfo"`
	LocationInfo     common.LocationDetailsStruct `json:"locationInfo"`
	RoomTypesInfo    []common.RoomType            `json:"roomTypesInfo"`
	MealsInfo        common.MealsInfo             `json:"mealsInfo"`
	AvailabilityInfo common.AvailabilityInfo      `json:"availabilityInfo"`
	PoliciesInfo     common.PoliciesInfo          `json:"policiesInfo"`
}

type ReqStruct struct {
	ClientId string `json:"clientId"`
	Stage    string `json:"stage"`
}

func GetUserDetailsAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetUserDetailsAPI (+)")

	if r.Method == http.MethodPost {
		var lResponse Response

		lResponse.Status = common.SUCCESSCODE

		var lReq ReqStruct

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IBDAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IBDAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IBDAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IBDAPI002", lErr.Error()))
			return
		}

		if strings.EqualFold(lReq.Stage, "Basic Info") {
			lResponse.BasicInfo, lErr = GetBasicInfo(lDebug, lReq.ClientId)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "IBDAPI002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("IBDAPI002", lErr.Error()))
				return
			}
		} else if strings.EqualFold(lReq.Stage, "Location") {
			lResponse.LocationInfo, lErr = GetLocationInfo(lDebug, lReq.ClientId)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "IBDAPI002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("IBDAPI002", lErr.Error()))
				return
			}

		} else if strings.EqualFold(lReq.Stage, "Room Details") {
			lResponse.RoomTypesInfo, lErr = GetRoomDetailsInfo(lDebug, lReq.ClientId)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "IBDAPI002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("IBDAPI002", lErr.Error()))
				return
			}

		} else if strings.EqualFold(lReq.Stage, "Restaurant & Meals") {
			lResponse.MealsInfo, lErr = GetMealsInfo(lDebug, lReq.ClientId)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "IBDAPI002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("IBDAPI002", lErr.Error()))
				return
			}

		} else if strings.EqualFold(lReq.Stage, "Availability") {
			lResponse.AvailabilityInfo, lErr = GetAvailabilityInfo(lDebug, lReq.ClientId)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "IBDAPI002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("IBDAPI002", lErr.Error()))
				return
			}

		} else if strings.EqualFold(lReq.Stage, "Policies") {
			lResponse.PoliciesInfo, lErr = GetPolicieInfo(lDebug, lReq.ClientId)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "IBDAPI002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("IBDAPI002", lErr.Error()))
				return
			}

		} else if strings.EqualFold(lReq.Stage, "Docs") {

		} else if strings.EqualFold(lReq.Stage, "Property Details") {

		} else if strings.EqualFold(lReq.Stage, "Notes") {

		}

		lResp, lErr := json.Marshal(&lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GDDAPI002", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}
		fmt.Fprint(w, string(lResp))
	}
	lDebug.Log(helpers.Statement, "GetUserDetailsAPI (-)")
}

// GET BASIC INFORMATION

func GetBasicInfo(pDebug *helpers.HelperStruct, pReq string) (common.BasicDetailsStruct, error) {
	pDebug.Log(helpers.Statement, "GetBasicInfo (+)")

	var lBasicInfo common.BasicDetailsStruct

	lCoreString := `SELECT Hotel_name, Property_Type, Email, Year_Of_Construction, mobile_code, Primary_Mobile, Secondary_Mobile, Star_Category, Channel_Manageer
 	FROM basic_info where Uid = ? and isActive ='Y' order by id desc limit 1`

	lRows, lErr := database.Gdb.Query(lCoreString, pReq)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GBI001", lErr.Error())
		helpers.ErrReturn(lErr)
	}

	for lRows.Next() {
		lErr = lRows.Scan(&lBasicInfo.HotelName, &lBasicInfo.PropertyType, &lBasicInfo.Email, &lBasicInfo.YearOfConstruction, &lBasicInfo.MobileCode, &lBasicInfo.PrimaryMobile, &lBasicInfo.SecondaryMobile, &lBasicInfo.StarCategory, &lBasicInfo.ChannelManager)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GBI002", lErr.Error())
			helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "GetBasicInfo (-)")
	return lBasicInfo, nil
}

// GET LOCATION INFORMATION

func GetLocationInfo(pDebug *helpers.HelperStruct, pReq string) (common.LocationDetailsStruct, error) {
	pDebug.Log(helpers.Statement, "GetLocationInfo (+)")

	var lLocationInfo common.LocationDetailsStruct

	lCoreString := `  SELECT Addr_line1, Addr_line2, State, City, Zip_Code
					  FROM location_info where Uid = ? and isActive ='Y' order by id desc limit 1`

	lRows, lErr := database.Gdb.Query(lCoreString, pReq)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GLI001", lErr.Error())
		helpers.ErrReturn(lErr)
	}

	for lRows.Next() {
		lErr = lRows.Scan(&lLocationInfo.AddrLine1, &lLocationInfo.AddrLine2, &lLocationInfo.State, &lLocationInfo.City, &lLocationInfo.Zipcode)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GLI002", lErr.Error())
			helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "GetLocationInfo (-)")
	return lLocationInfo, nil
}

// GET ROOM INFORMATION

func GetRoomDetailsInfo(pDebug *helpers.HelperStruct, pReq string) ([]common.RoomType, error) {
	pDebug.Log(helpers.Statement, "GetRoomDetailsInfo (+)")

	var lRoomType common.RoomType
	var lRoomTypeArr []common.RoomType

	lCoreString := `SELECT RoomType, NoOfRooms, RoomView, RoomSizeUnit, RoomSize, MaximumOccupancy, ExtraBed, ExtraPersons, SingleGuestPrice, DoubleGuestPrice, TripleGuestPrice, ExtraAdultCharge, ChildCharge, BelowChildCharge, RoomAmenities, SmokingPolicy FROM room_types WHERE Uid = ? and isActive ='Y';`

	lRows, lErr := database.Gdb.Query(lCoreString, pReq)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GLI001", lErr.Error())
		helpers.ErrReturn(lErr)
	}

	for lRows.Next() {
		var lAmenitiesArrStr string

		lErr = lRows.Scan(&lRoomType.RoomType, &lRoomType.NoOfRooms, &lRoomType.RoomView, &lRoomType.RoomSizeUnit, &lRoomType.RoomSize, &lRoomType.MaximumOccupancy, &lRoomType.ExtraBed, &lRoomType.ExtraPersons, &lRoomType.SingleGuestPrice, &lRoomType.DoubleGuestPrice, &lRoomType.TripleGuestPrice, &lRoomType.ExtraAdultCharge, &lRoomType.ChildCharge, &lRoomType.BelowChildCharge, &lAmenitiesArrStr, &lRoomType.SmokingPolicy)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GLI002", lErr.Error())
			helpers.ErrReturn(lErr)
		}

		if strings.TrimSpace(lAmenitiesArrStr) == "" {
			lRoomType.RoomAmenities = []string{}
		} else {
			lRoomType.RoomAmenities = strings.Split(lAmenitiesArrStr, ",")
		}

		lRoomTypeArr = append(lRoomTypeArr, lRoomType)
	}

	pDebug.Log(helpers.Statement, "GetRoomDetailsInfo (-)")
	return lRoomTypeArr, nil
}

// GET MEALS INFORMATION

func GetMealsInfo(pDebug *helpers.HelperStruct, pReq string) (common.MealsInfo, error) {
	pDebug.Log(helpers.Statement, "GetMealsInfo (+)")

	var lMealsInfo common.MealsInfo

	lCoreString := `  SELECT isOptionalRestaurant, Meal_Package, Types_Of_Meals, Meal_Rack_Price
					  FROM meals_info where Uid = ? and isActive ='Y' order by id desc limit 1`

	lRows, lErr := database.Gdb.Query(lCoreString, pReq)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GMI001", lErr.Error())
		helpers.ErrReturn(lErr)
	}
	var lTypeOfMealsArrStr string

	for lRows.Next() {
		lErr = lRows.Scan(&lMealsInfo.IsOperationalRestaurant, &lMealsInfo.MealPackage, &lTypeOfMealsArrStr, &lMealsInfo.MealRackPrice)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GMI002", lErr.Error())
			helpers.ErrReturn(lErr)
		}

		if strings.TrimSpace(lTypeOfMealsArrStr) == "" {
			lMealsInfo.TypesOfMeals = []string{}
		} else {
			lMealsInfo.TypesOfMeals = strings.Split(lTypeOfMealsArrStr, ",")
		}
	}

	pDebug.Log(helpers.Statement, "GetMealsInfo (-)")
	return lMealsInfo, nil
}

// GET LOCATION INFORMATION

func GetAvailabilityInfo(pDebug *helpers.HelperStruct, pReq string) (common.AvailabilityInfo, error) {
	pDebug.Log(helpers.Statement, "GetAvailabilityInfo (+)")

	var lAvailabilityInfo common.AvailabilityInfo

	lCoreString := `  SELECT availability_Start_Date, availability_End_Date
					  FROM location_info where Uid = ? and isActive ='Y' order by id desc limit 1`

	lRows, lErr := database.Gdb.Query(lCoreString, pReq)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GAI001", lErr.Error())
		helpers.ErrReturn(lErr)
	}

	for lRows.Next() {
		lErr = lRows.Scan(&lAvailabilityInfo.StartDate, &lAvailabilityInfo.EndDate)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GAI002", lErr.Error())
			helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "GetAvailabilityInfo (-)")
	return lAvailabilityInfo, nil
}

// GET POLICIES INFORMATION

func GetPolicieInfo(pDebug *helpers.HelperStruct, pReq string) (common.PoliciesInfo, error) {
	pDebug.Log(helpers.Statement, "GetPolicieInfo (+)")

	var lPoliciesInfo common.PoliciesInfo

	lCoreString := ` SELECT check_in, check_out, checkinout_policy, CancellationPolicy, allow_unmarriedCouples, allow_minor_guest, allow_onlymale_guests, allow_smoking, allow_parties, allow_invite_guests, wheelchar_accessible, allow_pets, accepted_proofs, additional_propertyrules
	FROM property_policiesinfo where Uid =? and isActive ='Y'`

	lRows, lErr := database.Gdb.Query(lCoreString, pReq)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GPI001", lErr.Error())
		helpers.ErrReturn(lErr)
	}
	var lTypeOfProofsArrStr string

	for lRows.Next() {
		lErr = lRows.Scan(&lPoliciesInfo.Check_in, &lPoliciesInfo.Check_out, &lPoliciesInfo.Checkinout_policy, &lPoliciesInfo.CancellationPolicy, &lPoliciesInfo.Allow_unmarriedCouples, &lPoliciesInfo.Allow_minor_guest, &lPoliciesInfo.Allow_onlymale_guests, &lPoliciesInfo.Allow_smoking, &lPoliciesInfo.Allow_parties, &lPoliciesInfo.Allow_invite_guests, &lPoliciesInfo.Wheelchar_accessible, &lPoliciesInfo.Allow_pets, &lTypeOfProofsArrStr, &lPoliciesInfo.Additional_propertyrules)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GPI002", lErr.Error())
			helpers.ErrReturn(lErr)
		}

		if strings.TrimSpace(lTypeOfProofsArrStr) == "" {
			lPoliciesInfo.Accepted_proofs = []string{}
		} else {
			lPoliciesInfo.Accepted_proofs = strings.Split(lTypeOfProofsArrStr, ",")
		}
	}

	pDebug.Log(helpers.Statement, "GetPolicieInfo (-)")
	return lPoliciesInfo, nil
}
