package locationinfo

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func InsertLocationDetailsAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "InsertLocationDetailsAPI (+)")

	if r.Method == http.MethodPost {

		var lReq common.LocationDetailsStruct

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "ILDAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("ILDAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "ILDAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("ILDAPI002", lErr.Error()))
			return
		}

		// GET UID FROM COOKIE
		lReq.Uid, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "ILDAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("ILDAPI003", lErr.Error()))
			return
		}

		lErr = InsertLocationDetails(lDebug, lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "ILDAPI004", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("ILDAPI004", lErr.Error()))
			return
		}

	}
	fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	lDebug.Log(helpers.Statement, "InsertLocationDetailsAPI (-)")
}

func InsertLocationDetails(pDebug *helpers.HelperStruct, pReq common.LocationDetailsStruct) error {
	pDebug.Log(helpers.Statement, "InsertLocationDetails (+)")

	lExist, lErr := common.CheckUidInTable(pDebug, "location_info", pReq.Uid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "ILD001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	// UPDATE LOCATION INFO
	if lExist {
		lErr = UpdateLocationDetails(pDebug, pReq)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "ILD002", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	} else {
		// INSERT LOCATION INFO
		lCoreString := `INSERT INTO location_info
	(Uid, Addr_line1, Addr_line2, State, City, Zip_Code, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
	VALUES(?, ?, ?, ?, ?, ?, 'AutoBot', now(), 'AutoBot', now(), 'Y');`

		_, lErr = database.Gdb.Exec(lCoreString, pReq.Uid, pReq.AddrLine1, pReq.AddrLine2, pReq.State, pReq.City, pReq.Zipcode)

		if lErr != nil {
			pDebug.Log(helpers.Elog, "ILD003", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "InsertLocationDetails (-)")
	return nil

}

func UpdateLocationDetails(pDebug *helpers.HelperStruct, pReq common.LocationDetailsStruct) error {
	pDebug.Log(helpers.Statement, "UpdateLocationDetails (+)")

	lCoreString := `UPDATE location_info
SET  Addr_line1= ?, Addr_line2= ?, State= ?, City= ?, Zip_Code= ?, UpdatedBy='AutoBot', UpdatedDate= now()
WHERE Uid = ?;`

	_, lErr := database.Gdb.Exec(lCoreString, pReq.AddrLine1, pReq.AddrLine2, pReq.State, pReq.City, pReq.Zipcode, pReq.Uid)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "ULD001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "UpdateLocationDetails (-)")
	return nil

}
