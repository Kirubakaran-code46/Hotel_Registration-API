package basicinfo

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
	PropertyTypes []string `json:"propertyTypes"`
	Countrycode   []string `json:"countryCode"`
}

func GetBasicInfoDropdown(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetBasicInfoDropdown (+)")

	if r.Method == http.MethodGet {
		var lResponse Response
		var lProperty string
		var lPropertyArr []string
		var lCountryCode string
		var lCountryCodeArr []string

		lResponse.Status = common.SUCCESSCODE

		lCoreString := `SELECT property
						FROM hotel_property_types
						WHERE isActive='Y';`

		lRows, lErr := database.Gdb.Query(lCoreString)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GPTAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GPTAPI001", lErr.Error()))
			return
		}

		for lRows.Next() {
			lErr = lRows.Scan(&lProperty)

			if lErr != nil {
				lDebug.Log(helpers.Elog, "GPTAPI002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("GPTAPI002", lErr.Error()))
				return
			}
			lPropertyArr = append(lPropertyArr, lProperty)
		}
		lResponse.PropertyTypes = lPropertyArr

		// GET COUNTRY CODES
		lCoreString = `SELECT mobile_code
					   FROM country_code`

		lRows, lErr = database.Gdb.Query(lCoreString)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GPTAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GPTAPI003", lErr.Error()))
			return
		}

		for lRows.Next() {
			lErr = lRows.Scan(&lCountryCode)

			if lErr != nil {
				lDebug.Log(helpers.Elog, "GPTAPI004", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("GPTAPI004", lErr.Error()))
				return
			}
			lCountryCodeArr = append(lCountryCodeArr, lCountryCode)
		}
		lResponse.Countrycode = lCountryCodeArr

		lData, lErr := json.Marshal(lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GPTAPI005", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GPTAPI005", lErr.Error()))
			return
		}
		fmt.Fprint(w, string(lData))

	}
	lDebug.Log(helpers.Statement, "GetBasicInfoDropdown (-)")
}
