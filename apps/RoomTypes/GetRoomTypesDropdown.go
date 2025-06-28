package roomtypes

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"

	"fmt"
	"net/http"
)

type Response struct {
	Status        string   `json:"status"`
	ErrMsg        string   `json:"errMsg"`
	RoomTypes     []string `json:"roomTypes"`
	RoomView      []string `json:"roomView"`
	RoomAmenities []string `json:"roomAmenities"`
}

func GetRoomTypesDropdown(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetRoomTypesDropdown (+)")

	if r.Method == http.MethodGet {
		var lResponse Response
		var lRoomTypes string
		var lRoomTypesArr []string
		var lRoomView string
		var lRoomViewArr []string
		var lRoomAmenities string
		var lRoomAmenitiesArr []string

		lResponse.Status = common.SUCCESSCODE

		// GET ROOM TYPES

		lCoreString := `SELECT Types
						FROM room_type
						WHERE isActive='Y';`

		lRows, lErr := database.Gdb.Query(lCoreString)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRTD001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GRTD001", lErr.Error()))
			return
		}

		for lRows.Next() {
			lErr = lRows.Scan(&lRoomTypes)

			if lErr != nil {
				lDebug.Log(helpers.Elog, "GRTD002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("GRTD002", lErr.Error()))
				return
			}
			lRoomTypesArr = append(lRoomTypesArr, lRoomTypes)
		}
		lResponse.RoomTypes = lRoomTypesArr

		// GET ROOM VIEW

		lCoreString = `SELECT Types
						FROM room_view
						WHERE isActive='Y';`

		lRows, lErr = database.Gdb.Query(lCoreString)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRTD003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GRTD003", lErr.Error()))
			return
		}

		for lRows.Next() {
			lErr = lRows.Scan(&lRoomView)

			if lErr != nil {
				lDebug.Log(helpers.Elog, "GRTD004", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("GRTD004", lErr.Error()))
				return
			}
			lRoomViewArr = append(lRoomViewArr, lRoomView)
		}
		lResponse.RoomView = lRoomViewArr

		// GET ROOM AMENITIES

		lCoreString = `SELECT Types
						FROM room_amenities
						WHERE isActive='Y';`

		lRows, lErr = database.Gdb.Query(lCoreString)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRTD005", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GRTD005", lErr.Error()))
			return
		}

		for lRows.Next() {
			lErr = lRows.Scan(&lRoomAmenities)

			if lErr != nil {
				lDebug.Log(helpers.Elog, "GRTD006", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("GRTD006", lErr.Error()))
				return
			}
			lRoomAmenitiesArr = append(lRoomAmenitiesArr, lRoomAmenities)
		}
		lResponse.RoomAmenities = lRoomAmenitiesArr

		lData, lErr := json.Marshal(lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRTD007", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GRTD007", lErr.Error()))
			return
		}
		fmt.Fprint(w, string(lData))

	}
	lDebug.Log(helpers.Statement, "GetRoomTypesDropdown (-)")
}
