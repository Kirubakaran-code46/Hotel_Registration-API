package description

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func InsertDescInfoAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "InsertDescInfoAPI (+)")

	if r.Method == http.MethodPost {

		var lReq common.Description

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IDescAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IDescAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IDescAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IDescAPI002", lErr.Error()))
			return
		}

		// GET UID FROM COOKIE
		lReq.Uid, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IDescAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IDescAPI003", lErr.Error()))
			return
		}

		lErr = InsertDescInfo(lDebug, lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IDescAPI004", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IDescAPI004", lErr.Error()))
			return
		}

	}
	fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	lDebug.Log(helpers.Statement, "InsertDescInfoAPI (-)")
}

func InsertDescInfo(pDebug *helpers.HelperStruct, pReq common.Description) error {
	pDebug.Log(helpers.Statement, "InsertDescInfo (+)")

	lExist, lErr := common.CheckUidInTable(pDebug, "basic_info", pReq.Uid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IDesc001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	// UPDATE DESC INFO
	if lExist {
		lErr = UpdateDescInfo(pDebug, pReq)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IDesc002", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	} else {
		lQueryString := `INSERT INTO basic_info
(Uid, Description, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
VALUES( ?,?, 'Autobot', now(), 'Autobot', now(), 'Y');`

		_, lErr := database.Gdb.Exec(lQueryString, pReq.Uid, pReq.Description)

		if lErr != nil {
			pDebug.Log(helpers.Elog, "IDesc003", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "InsertDescInfo (-)")
	return nil

}

func UpdateDescInfo(pDebug *helpers.HelperStruct, pReq common.Description) error {
	pDebug.Log(helpers.Statement, "UpdateDescInfo (+)")

	lQueryString := `update basic_info set Description= ? where Uid =? and isActive ='Y'`

	_, lErr := database.Gdb.Exec(lQueryString, pReq.Description, pReq.Uid)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "UDesc001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "UpdateDescInfo (-)")
	return nil

}
