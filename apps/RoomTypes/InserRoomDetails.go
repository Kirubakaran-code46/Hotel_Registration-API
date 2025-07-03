package roomtypes

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

type RoomRequest struct {
	RoomsArr []common.RoomType `json:"roomsArr"`
}

func InsertRoomDetailsAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "InsertRoomDetailsAPI (+)")

	if r.Method == http.MethodPost {

		var lReq RoomRequest

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IRDAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IRDAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IRDAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IRDAPI002", lErr.Error()))
			return
		}
		var lCookieVal string

		// GET UID FROM COOKIE
		lCookieVal, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IRDAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IRDAPI003", lErr.Error()))
			return
		}

		lErr = InsertRoomsDetails(lDebug, lReq.RoomsArr, lCookieVal)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IRDAPI004", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IRDAPI004", lErr.Error()))
			return
		}
	}

	fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	lDebug.Log(helpers.Statement, "InsertRoomDetailsAPI (-)")
}

func InsertRoomsDetails(pDebug *helpers.HelperStruct, pReq []common.RoomType, pCookieVal string) error {
	pDebug.Log(helpers.Statement, "InsertRoomsDetails (+)")

	lUidEsist, lErr := common.CheckUidInTable(pDebug, "room_types", pCookieVal)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IRD001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	if lUidEsist {
		// SOFT DELETE OLD VALUES
		lQueryString := `update room_types set isActive='N' where Uid=?`

		_, lErr = database.Gdb.Exec(lQueryString, pCookieVal)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IRD002", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	}

	for _, Req := range pReq {
		lRoomAmenitiesStr := strings.Join(Req.RoomAmenities, ",")
		// INSERT NEW VALUES
		lQueryString := `INSERT INTO room_types
						(Uid, RoomType, NoOfRooms, RoomView, RoomSizeUnit, RoomSize, MaximumOccupancy, ExtraBed, ExtraPersons, SingleGuestPrice, DoubleGuestPrice, TripleGuestPrice, ExtraAdultCharge, ChildCharge, BelowChildCharge, RoomAmenities, SmokingPolicy, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
						VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 'AutoBot', now(), 'AutoBot', now(), 'Y');`

		_, lErr = database.Gdb.Exec(lQueryString, pCookieVal, Req.RoomType, Req.NoOfRooms, Req.RoomView, Req.RoomSizeUnit, Req.RoomSize, Req.MaximumOccupancy, Req.ExtraBed, Req.ExtraPersons, Req.SingleGuestPrice, Req.DoubleGuestPrice, Req.TripleGuestPrice, Req.ExtraAdultCharge, Req.ChildCharge, Req.BelowChildCharge, lRoomAmenitiesStr, Req.SmokingPolicy)

		if lErr != nil {
			pDebug.Log(helpers.Elog, "IRD003", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "InsertRoomsDetails (-)")
	return nil
}
