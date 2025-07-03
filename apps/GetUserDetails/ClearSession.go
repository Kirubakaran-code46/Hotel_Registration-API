package getuserdetails

import (
	"HOTEL-REGISTRY_API/common"
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

		_, lErr := r.Cookie(common.UIDCOOKIENAME)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "CSAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("CSAPI001", lErr.Error()))
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
