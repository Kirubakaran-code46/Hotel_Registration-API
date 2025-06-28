package policyinfo

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"

	"fmt"
	"net/http"
)

type Response struct {
	Status               string   `json:"status"`
	ErrMsg               string   `json:"errMsg"`
	IdentityProofs       []string `json:"identityProofs"`
	CancellationPolicies []string `json:"cancellationPolicies"`
}

func GetPoliciesInfoDropdown(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetPoliciesInfoDropdown (+)")

	if r.Method == http.MethodGet {
		var lResponse Response
		var lIdentityProofs string
		var lIdentityProofsArr []string
		var lCancellationPolicies string
		var lCancellationPoliciesArr []string

		lResponse.Status = common.SUCCESSCODE

		lCoreString := `SELECT proof
						FROM accepted_identity_proofs
						WHERE isActive='Y';`

		lRows, lErr := database.Gdb.Query(lCoreString)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GPIDAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GPIDAPI001", lErr.Error()))
			return
		}

		for lRows.Next() {
			lErr = lRows.Scan(&lIdentityProofs)

			if lErr != nil {
				lDebug.Log(helpers.Elog, "GPIDAPI002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("GPIDAPI002", lErr.Error()))
				return
			}
			lIdentityProofsArr = append(lIdentityProofsArr, lIdentityProofs)
		}
		lResponse.IdentityProofs = lIdentityProofsArr

		// GET COUNTRY CODES
		lCoreString = `SELECT policy
					   FROM hotel_cancellation_policy WHERE isActive='Y';`

		lRows, lErr = database.Gdb.Query(lCoreString)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GPIDAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GPIDAPI003", lErr.Error()))
			return
		}

		for lRows.Next() {
			lErr = lRows.Scan(&lCancellationPolicies)

			if lErr != nil {
				lDebug.Log(helpers.Elog, "GPIDAPI004", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("GPIDAPI004", lErr.Error()))
				return
			}
			lCancellationPoliciesArr = append(lCancellationPoliciesArr, lCancellationPolicies)
		}
		lResponse.CancellationPolicies = lCancellationPoliciesArr

		lData, lErr := json.Marshal(lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GPIDAPI005", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GPIDAPI005", lErr.Error()))
			return
		}
		fmt.Fprint(w, string(lData))

	}
	lDebug.Log(helpers.Statement, "GetPoliciesInfoDropdown (-)")
}
