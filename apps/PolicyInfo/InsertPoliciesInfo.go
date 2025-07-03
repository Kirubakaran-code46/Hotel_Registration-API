package policyinfo

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func InsertPolicyInfoAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "InsertPolicyInfoAPI (+)")

	if r.Method == http.MethodPost {

		var lReq common.PoliciesInfo

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IPIAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IPIAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IPIAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IPIAPI002", lErr.Error()))
			return
		}

		// GET UID FROM COOKIE
		lReq.Uid, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IPIAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IPIAPI003", lErr.Error()))
			return
		}

		lErr = InsertPolicyDetails(lDebug, lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IPIAPI004", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IPIAPI004", lErr.Error()))
			return
		}

	}
	fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	lDebug.Log(helpers.Statement, "InsertPolicyInfoAPI (-)")
}

func InsertPolicyDetails(pDebug *helpers.HelperStruct, pReq common.PoliciesInfo) error {
	pDebug.Log(helpers.Statement, "InsertPolicyDetails (+)")

	lExist, lErr := common.CheckUidInTable(pDebug, "property_policiesinfo", pReq.Uid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IPD001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	// UPDATE PROPERTY POLICIES INFO
	if lExist {
		lErr = UpdatePoliciesDetails(pDebug, pReq)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IPD002", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	} else {
		var lCoreString string
		lTypesOfProofsArr := strings.Join(pReq.Accepted_proofs, ",")

		// INSERT MEALS INFO

		lCoreString = `INSERT INTO property_policiesinfo
(Uid, check_in, check_out, checkinout_policy, CancellationPolicy, allow_unmarriedCouples, allow_minor_guest, allow_onlymale_guests, allow_smoking, allow_parties, allow_invite_guests, wheelchar_accessible, allow_pets, accepted_proofs, additional_propertyrules, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 'AutoBot', now(),'AutoBot', now(), 'Y');`

		_, lErr = database.Gdb.Exec(lCoreString, pReq.Uid, pReq.Check_in, pReq.Check_out, pReq.Checkinout_policy, pReq.CancellationPolicy, pReq.Allow_unmarriedCouples, pReq.Allow_minor_guest, pReq.Allow_onlymale_guests, pReq.Allow_smoking, pReq.Allow_parties, pReq.Allow_invite_guests, pReq.Wheelchar_accessible, pReq.Allow_pets, lTypesOfProofsArr, pReq.Additional_propertyrules)

		if lErr != nil {
			pDebug.Log(helpers.Elog, "IPD003", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "InsertPolicyDetails (-)")
	return nil

}

func UpdatePoliciesDetails(pDebug *helpers.HelperStruct, pReq common.PoliciesInfo) error {
	pDebug.Log(helpers.Statement, "UpdatePoliciesDetails (+)")
	lTypesOfProofsArrStr := strings.Join(pReq.Accepted_proofs, ",")
	var lCoreString string
	var lErr error

	lCoreString = `UPDATE property_policiesinfo
SET check_in= ?, check_out= ?, checkinout_policy=?, CancellationPolicy= ?, allow_unmarriedCouples=?, allow_minor_guest=?, allow_onlymale_guests=?, allow_smoking=?, allow_parties=?, allow_invite_guests=?, wheelchar_accessible=?, allow_pets=?, accepted_proofs= ?, additional_propertyrules= ?, UpdatedBy= 'AutoBot', UpdatedDate= now()
WHERE Uid=? and isActive='Y';`

	_, lErr = database.Gdb.Exec(lCoreString, pReq.Check_in, pReq.Check_out, pReq.Checkinout_policy, pReq.CancellationPolicy, pReq.Allow_unmarriedCouples, pReq.Allow_minor_guest, pReq.Allow_onlymale_guests, pReq.Allow_smoking, pReq.Allow_parties, pReq.Allow_invite_guests, pReq.Wheelchar_accessible, pReq.Allow_pets, lTypesOfProofsArrStr, pReq.Additional_propertyrules, pReq.Uid)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "UPD001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "UpdatePoliciesDetails (-)")
	return nil

}
