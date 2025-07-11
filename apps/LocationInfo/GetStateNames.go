package locationinfo

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"

	"fmt"
	"net/http"
)

type Response struct {
	Status     string   `json:"status"`
	ErrMsg     string   `json:"errMsg"`
	StateNames []string `json:"stateNames"`
	CityNames  []string `json:"cityNames"`
}

func GetStateDropdown(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetStateDropdown (+)")

	if r.Method == http.MethodGet {
		var lResponse Response
		var lCities string
		var lCitiesArr []string

		lResponse.Status = common.SUCCESSCODE

		// lCoreString := `SELECT State_Name
		// 				FROM indian_states where isActive='Y'`

		// lRows, lErr := database.Gdb.Query(lCoreString)

		// if lErr != nil {
		// 	lDebug.Log(helpers.Elog, "GSDAPI001", lErr.Error())
		// 	fmt.Fprint(w, helpers.GetError_String("GSDAPI001", lErr.Error()))
		// 	return
		// }

		// for lRows.Next() {
		// 	lErr = lRows.Scan(&lStates)

		// 	if lErr != nil {
		// 		lDebug.Log(helpers.Elog, "GSDAPI002", lErr.Error())
		// 		fmt.Fprint(w, helpers.GetError_String("GSDAPI002", lErr.Error()))
		// 		return
		// 	}
		// 	lStatesArr = append(lStatesArr, lStates)
		// }
		// lResponse.StateNames = lStatesArr

		// GET CITIES
		lCoreString := `select name from city_station_mapping`

		lRows, lErr := database.Gdb.Query(lCoreString)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GSDAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GSDAPI001", lErr.Error()))
			return
		}

		for lRows.Next() {
			lErr = lRows.Scan(&lCities)

			if lErr != nil {
				lDebug.Log(helpers.Elog, "GSDAPI002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("GSDAPI002", lErr.Error()))
				return
			}
			lCitiesArr = append(lCitiesArr, lCities)
		}
		lResponse.CityNames = lCitiesArr

		lData, lErr := json.Marshal(lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GSDAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GSDAPI003", lErr.Error()))
			return
		}
		fmt.Fprint(w, string(lData))

	}
	lDebug.Log(helpers.Statement, "GetStateDropdown (-)")
}
