package basicinfo

// import (
// 	database "HOTEL-REGISTRY_API/Db_Setup"
// 	"HOTEL-REGISTRY_API/common"
// 	"HOTEL-REGISTRY_API/helpers"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func UpdateBasicDetailsAPI(w http.ResponseWriter, r *http.Request) {
// 	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
// 	(w).Header().Set("Access-Control-Allow-Credentials", "true")
// 	(w).Header().Set("Access-Control-Allow-Methods", "POST")
// 	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

// 	lDebug := new(helpers.HelperStruct)

// 	lDebug.SetUid(r)
// 	lDebug.Log(helpers.Statement, "UpdateBasicDetailsAPI (+)")

// 	if r.Method == http.MethodPost {

// 		var lReq common.BasicDetailsStruct

// 		lBody, lErr := io.ReadAll(r.Body)
// 		if lErr != nil {
// 			lDebug.Log(helpers.Elog, "UBDAPI001", lErr.Error())
// 			fmt.Fprint(w, helpers.GetError_String("UBDAPI001", lErr.Error()))
// 			return
// 		}

// 		lErr = json.Unmarshal(lBody, &lReq)
// 		if lErr != nil {
// 			lDebug.Log(helpers.Elog, "UBDAPI002", lErr.Error())
// 			fmt.Fprint(w, helpers.GetError_String("UBDAPI002", lErr.Error()))
// 			return
// 		}

// 		// GET UID FROM COOKIE
// 		lReq.Uid, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
// 		if lErr != nil {
// 			lDebug.Log(helpers.Elog, "UBDAPI003", lErr.Error())
// 			fmt.Fprint(w, helpers.GetError_String("UBDAPI003", lErr.Error()))
// 			return
// 		}
// 		fmt.Println(", pReq.PrimaryMobile cookie-->", lReq.Uid)

// 		lErr = UpdateBasicDetails(lDebug, lReq)
// 		if lErr != nil {
// 			lDebug.Log(helpers.Elog, "UBDAPI004", lErr.Error())
// 			fmt.Fprint(w, helpers.GetError_String("UBDAPI004", lErr.Error()))
// 			return
// 		}
// 	}
// 	fmt.Fprint(w, helpers.GetMsg_String("S", "Updated Successfully"))
// 	lDebug.Log(helpers.Statement, "UpdateBasicDetailsAPI (-)")
// }

// func UpdateBasicDetails(pDebug *helpers.HelperStruct, pReq common.BasicDetailsStruct) error {
// 	pDebug.Log(helpers.Statement, "UpdateBasicDetails (+)")

// 	lCoreString := `UPDATE basic_info
// SET Hotel_name= ?, Property_Type= ?, Email= ?, Year_Of_Construction= ?, mobile_code= ?, Primary_Mobile= ?, Secondary_Mobile= ?, Star_Category= ?, Channel_Manageer= ?, UpdatedBy='Autobot', UpdatedDate= now()
// WHERE Uid =?;`

// 	_, lErr := database.Gdb.Exec(lCoreString, pReq.HotelName, pReq.PropertyType, pReq.Email, pReq.YearOfConstruction, pReq.MobileCode, pReq.PrimaryMobile, pReq.SecondaryMobile, pReq.StarCategory, pReq.ChannelManager, pReq.Uid)

// 	if lErr != nil {
// 		pDebug.Log(helpers.Elog, "UBD001", lErr.Error())
// 		helpers.ErrReturn(lErr)
// 	}

// 	pDebug.Log(helpers.Statement, "UpdateBasicDetails (-)")
// 	return nil

// }
