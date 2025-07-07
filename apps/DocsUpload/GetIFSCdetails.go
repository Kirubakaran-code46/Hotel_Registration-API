package docsupload

import (
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type lIfscResp struct {
	Status   string              `json:"status"`
	ErrMsg   string              `json:"errMsg"`
	IFSCdata common.IFSCResponse `json:"IFSCdata"`
}

type IfscReq struct {
	IFSC string `json:"ifsc"`
}

func GetIFSCdetailsAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetIFSCdetailsAPI (+)")

	if r.Method == http.MethodPost {
		var lResponse lIfscResp

		lResponse.Status = common.SUCCESSCODE

		var lReq IfscReq

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GIFSCAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GIFSCAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GIFSCAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GIFSCAPI002", lErr.Error()))
			return
		}

		if strings.TrimSpace(lReq.IFSC) == "" {
			lDebug.Log(helpers.Elog, "GIFSCAPI003", fmt.Errorf("IFSC NotFount"))
			fmt.Fprint(w, helpers.GetError_String("GIFSCAPI003", "IFSC NotFount"))
			return
		}

		// GET UID FROM COOKIE
		lUid, lErr := common.GetCookieValue(r, common.UIDCOOKIENAME)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GIFSCAPI004", fmt.Errorf("cookie not Found"))
			fmt.Fprint(w, helpers.GetError_String("GIFSCAPI004", "Cookie Not Found"))
			return
		}

		// Fetch IFSC data
		lResponse.IFSCdata, lErr = common.GetIFSCDetails(lDebug, lReq.IFSC, lUid)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GIFSCAPI005", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GIFSCAPI005", lErr.Error()))
			return
		}

		lData, lErr := json.Marshal(lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GIFSCAPI006", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GIFSCAPI006", lErr.Error()))
			return
		}
		fmt.Fprint(w, string(lData))

	}
	lDebug.Log(helpers.Statement, "GetIFSCdetailsAPI (-)")
}
