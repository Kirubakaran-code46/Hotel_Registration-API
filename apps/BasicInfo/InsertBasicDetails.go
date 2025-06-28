package basicinfo

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func InsertBasicDetailsAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "InsertBasicDetailsAPI (+)")

	if r.Method == http.MethodPost {

		var lReq common.BasicDetailsStruct

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
		// GET UID FROM COOKIE
		lReq.Uid, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "ILDAPI003", lErr.Error())
		}

		if lReq.Uid != "" {

			lErr = UpdateBasicDetails(lDebug, lReq)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "ILDAPI004", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("ILDAPI004", lErr.Error()))
				return
			}
		} else {
			lclientId, lErr := AddBasicDetails(lDebug, lReq)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "IBDAPI005", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("IBDAPI005", lErr.Error()))
				return
			}

			//  Set cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "client_id",
				Value:    lclientId,
				Path:     "/",
				HttpOnly: false,
				Secure:   false,             // change to true if using HTTPS
				MaxAge:   60 * 60 * 24 * 30, // 30 days
			})
		}

	}
	fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	lDebug.Log(helpers.Statement, "InsertBasicDetailsAPI (-)")
}

func AddBasicDetails(pDebug *helpers.HelperStruct, pReq common.BasicDetailsStruct) (string, error) {
	pDebug.Log(helpers.Statement, "AddBasicDetails (+)")

	clientID := uuid.New().String()

	lCoreString := `  INSERT INTO basic_info
(Uid, Hotel_name, Property_Type, Email, Year_Of_Construction, mobile_code, Primary_Mobile, Secondary_Mobile, Star_Category, Channel_Manageer, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
VALUES( ?,?, ?, ?, ?, ?, ?, ?, ?,?, 'Autobot', now(), 'Autobot', now(), 'Y');`

	_, lErr := database.Gdb.Exec(lCoreString, clientID, pReq.HotelName, pReq.PropertyType, pReq.Email, pReq.YearOfConstruction, pReq.MobileCode, pReq.PrimaryMobile, pReq.SecondaryMobile, pReq.StarCategory, pReq.ChannelManager)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "ABD001", lErr.Error())
		helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "AddBasicDetails (-)")
	return clientID, nil

}

func UpdateBasicDetails(pDebug *helpers.HelperStruct, pReq common.BasicDetailsStruct) error {
	pDebug.Log(helpers.Statement, "UpdateBasicDetails (+)")

	lCoreString := `UPDATE basic_info
SET Hotel_name= ?, Property_Type= ?, Email= ?, Year_Of_Construction= ?, mobile_code= ?, Primary_Mobile= ?, Secondary_Mobile= ?, Star_Category= ?, Channel_Manageer= ?, UpdatedBy='Autobot', UpdatedDate= now()
WHERE Uid =?;`

	_, lErr := database.Gdb.Exec(lCoreString, pReq.HotelName, pReq.PropertyType, pReq.Email, pReq.YearOfConstruction, pReq.MobileCode, pReq.PrimaryMobile, pReq.SecondaryMobile, pReq.StarCategory, pReq.ChannelManager, pReq.Uid)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "UBD001", lErr.Error())
		helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "UpdateBasicDetails (-)")
	return nil

}
