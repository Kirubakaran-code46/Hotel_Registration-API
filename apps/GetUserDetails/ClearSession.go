package getuserdetails

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"strconv"
	"strings"

	"HOTEL-REGISTRY_API/helpers"
	"fmt"
	"net/http"
)

func ClearCookieAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetBasicInfoDropdown (+)")

	if r.Method == http.MethodGet {

		lUid, lErr := common.GetCookieValue(r, common.UIDCOOKIENAME)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "CSAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("CSAPI001", lErr.Error()))
			return
		}
		// INSERT PROCESS FLOW TABLE
		lErr = InsertReqInAdminPanel(lDebug, lUid)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "CSAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("CSAPI002", lErr.Error()))
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     common.UIDCOOKIENAME,
			Value:    "",
			Path:     "/",   // Must match the Path used when setting the cookie
			MaxAge:   -1,    // Deletes the cookie
			HttpOnly: true,  // Match the original cookie settings if needed
			Secure:   false, // Set to true if using HTTPS
		})

		fmt.Fprint(w, helpers.GetMsg_String("S", "Session Cleard"))
	}

	lDebug.Log(helpers.Statement, "InsertBasicDetailsAPI (-)")
}

func InsertReqInAdminPanel(pDebug *helpers.HelperStruct, pUid string) error {
	pDebug.Log(helpers.Statement, "InsertReqInAdminPanel (+)")

	ClientId, lErr := GenerateClientId(pDebug, pUid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IRAP001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	lQueryString := `INSERT INTO requests
					(Uid, ClientId, Process, stage, isActive, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate)
					VALUES(?,?,'I','I','Y','AutoBot',now(),'AutoBot',now());`

	_, lErr = database.Gdb.Exec(lQueryString, pUid, ClientId)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IRAP002", lErr.Error())
		return helpers.ErrReturn(lErr)
	}
	pDebug.Log(helpers.Statement, "InsertReqInAdminPanel (-)")
	return nil
}

func GenerateClientId(pDebug *helpers.HelperStruct, pUid string) (string, error) {
	pDebug.Log(helpers.Statement, "GenerateClientId (+)")

	lClientId := ""
	var lCityCode string
	var lStateCode string
	var lSerialNum int
	var lCityAvailability string

	// GET STATE AND CITY IN GIVEN CODE
	lQueryString := `SELECT 
    scm.code AS state_code,
    COALESCE(csm.station_code, li.city) AS city_or_station,
    CASE 
        WHEN csm.station_code IS NOT NULL THEN 'Y'
        ELSE 'N'
    END AS city_code_available
	FROM 
	    location_info li
	JOIN 
	    state_code_mapping scm ON li.state = scm.state
	LEFT JOIN 
	    city_station_mapping csm ON li.city = csm.name
	WHERE 
	    li.Uid = ?
	    AND li.isActive = 'Y';`

	lRows, lErr := database.Gdb.Query(lQueryString, pUid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "GCID001", lErr.Error())
		return lClientId, helpers.ErrReturn(lErr)
	}

	for lRows.Next() {
		lErr = lRows.Scan(&lStateCode, &lCityCode, &lCityAvailability)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GCID002", lErr.Error())
			return lClientId, helpers.ErrReturn(lErr)
		}
	}

	// CITY NOT PRESENT INT THE CITY MAPPING TABLE -> GENERATE OWN CITY CODE

	if strings.EqualFold(lCityAvailability, "N") {
		lCityCode = strings.TrimSpace(lCityCode)
		lCityCode = strings.ToUpper(lCityCode)

		if len(lCityCode) >= 3 {
			lCityCode = lCityCode[:3]
		}
	}

	// GET LAST SEQUENTIAL NUMBER FROM OLD REQUESTS

	lQueryString = `SELECT 
  IFNULL(
    (SELECT 
       CAST(REGEXP_SUBSTR(ClientId, '[0-9]+$') AS UNSIGNED) + 1
     FROM requests
     WHERE ClientId REGEXP '[0-9]+$'
     ORDER BY id DESC
     LIMIT 1),
    1
  ) AS next_number;`

	lRows, lErr = database.Gdb.Query(lQueryString)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "GCID003", lErr.Error())
		return lClientId, helpers.ErrReturn(lErr)
	}

	for lRows.Next() {
		lErr = lRows.Scan(&lSerialNum)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GCID004", lErr.Error())
			return lClientId, helpers.ErrReturn(lErr)
		}
	}

	// GENERATE CLIENT ID => PRODUCTNAME + STATE CODE + CITY CODE + SERIAL NUMBER

	lClientId = strings.TrimSpace(common.OLLIVBRK) + strings.TrimSpace(lStateCode) + strings.TrimSpace(lCityCode) + strings.TrimSpace(strconv.Itoa(lSerialNum))

	pDebug.Log(helpers.Statement, "GenerateClientId (-)")
	return lClientId, nil
}
