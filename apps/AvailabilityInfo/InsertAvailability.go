package availabilityinfo

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	locationinfo "HOTEL-REGISTRY_API/apps/LocationInfo"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func InsertAvailabilityDetailsAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "InsertAvailabilityDetailsAPI (+)")

	if r.Method == http.MethodPost {

		var lReq common.AvailabilityInfo

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IADAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IADAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IADAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IADAPI002", lErr.Error()))
			return
		}

		// GET UID FROM COOKIE
		lReq.Uid, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
		fmt.Println("lReq.Uid", lReq.Uid)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IADAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IADAPI003", lErr.Error()))
			return
		}

		lErr = InsertAvailabilityDetails(lDebug, lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IADAPI004", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IADAPI004", lErr.Error()))
			return
		}

	}
	fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	lDebug.Log(helpers.Statement, "InsertAvailabilityDetailsAPI (-)")
}

func InsertAvailabilityDetails(pDebug *helpers.HelperStruct, pReq common.AvailabilityInfo) error {
	pDebug.Log(helpers.Statement, "InsertAvailabilityDetails (+)")

	lExist, lErr := locationinfo.CheckUidInTable(pDebug, "location_info", pReq.Uid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IAD001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	// UPDATE LOCATION INFO
	if lExist {
		lErr = UpdateAvailability(pDebug, pReq)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IAD002", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	} else {
		// INSERT LOCATION INFO
		lCoreString := `INSERT INTO location_info
	(Uid, availability_Start_Date, availability_End_Date, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
	VALUES(?, ?, ?,'AutoBot', now(), 'AutoBot', now(), 'Y');`

		_, lErr = database.Gdb.Exec(lCoreString, pReq.Uid, pReq.StartDate, pReq.EndDate)

		if lErr != nil {
			pDebug.Log(helpers.Elog, "IAD003", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "InsertAvailabilityDetails (-)")
	return nil

}

func UpdateAvailability(pDebug *helpers.HelperStruct, pReq common.AvailabilityInfo) error {
	pDebug.Log(helpers.Statement, "UpdateAvailability (+)")

	lCoreString := `UPDATE location_info
SET  availability_Start_Date = ?, availability_End_Date = ?, UpdatedBy='AutoBot', UpdatedDate= now()
WHERE Uid = ?;`

	_, lErr := database.Gdb.Exec(lCoreString, pReq.StartDate, pReq.EndDate, pReq.Uid)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "UAD001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "UpdateAvailability (-)")
	return nil

}
