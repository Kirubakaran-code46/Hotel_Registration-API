package locationinfo

// import (
// 	database "HOTEL-REGISTRY_API/Db_Setup"
// 	"HOTEL-REGISTRY_API/common"
// 	"HOTEL-REGISTRY_API/helpers"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func UpdateLocationDetailsAPI(w http.ResponseWriter, r *http.Request) {
// 	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
// 	(w).Header().Set("Access-Control-Allow-Credentials", "true")
// 	(w).Header().Set("Access-Control-Allow-Methods", "POST")
// 	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

// 	lDebug := new(helpers.HelperStruct)

// 	lDebug.SetUid(r)
// 	lDebug.Log(helpers.Statement, "UpdateLocationDetailsAPI (+)")

// 	if r.Method == http.MethodPost {

// 		var lReq common.LocationDetailsStruct

// 		lBody, lErr := io.ReadAll(r.Body)
// 		if lErr != nil {
// 			lDebug.Log(helpers.Elog, "ULDAPI001", lErr.Error())
// 			fmt.Fprint(w, helpers.GetError_String("ULDAPI001", lErr.Error()))
// 			return
// 		}

// 		lErr = json.Unmarshal(lBody, &lReq)
// 		if lErr != nil {
// 			lDebug.Log(helpers.Elog, "ULDAPI002", lErr.Error())
// 			fmt.Fprint(w, helpers.GetError_String("ULDAPI002", lErr.Error()))
// 			return
// 		}

// 		// GET UID FROM COOKIE
// 		lReq.Uid, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
// 		if lErr != nil {
// 			lDebug.Log(helpers.Elog, "ULDAPI003", lErr.Error())
// 			fmt.Fprint(w, helpers.GetError_String("ULDAPI003", lErr.Error()))
// 			return
// 		}
// 		fmt.Println(", pReq.PrimaryMobile cookie-->", lReq.Uid)

// 		lErr = UpdateBasicDetails(lDebug, lReq)
// 		if lErr != nil {
// 			lDebug.Log(helpers.Elog, "ULDAPI004", lErr.Error())
// 			fmt.Fprint(w, helpers.GetError_String("ULDAPI004", lErr.Error()))
// 			return
// 		}
// 	}
// 	fmt.Fprint(w, helpers.GetMsg_String("S", "Updated Successfully"))
// 	lDebug.Log(helpers.Statement, "UpdateLocationDetailsAPI (-)")
// }

// func UpdateBasicDetails(pDebug *helpers.HelperStruct, pReq common.LocationDetailsStruct) error {
// 	pDebug.Log(helpers.Statement, "UpdateBasicDetails (+)")

// 	lCoreString := `UPDATE location_info
// SET  Addr_line1= ?, Addr_line2= ?, State= ?, City= ?, Zip_Code= ?, UpdatedBy='AutoBot', UpdatedDate= now()
// WHERE Uid = ?;`

// 	_, lErr := database.Gdb.Exec(lCoreString, pReq.AddrLine1, pReq.AddrLine2, pReq.State, pReq.City, pReq.Zipcode, pReq.Uid)

// 	if lErr != nil {
// 		pDebug.Log(helpers.Elog, "ULD001", lErr.Error())
// 		helpers.ErrReturn(lErr)
// 	}

// 	pDebug.Log(helpers.Statement, "UpdateBasicDetails (-)")
// 	return nil

// }
