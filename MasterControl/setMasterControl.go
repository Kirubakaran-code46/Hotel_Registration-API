package mastercontrol

import (
	"encoding/json"
	"fcs23pkg/apps/validation/apiaccess"
	"fcs23pkg/common"
	"fcs23pkg/ftdb"
	"fcs23pkg/helpers"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type masterResp struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

// Addin Router Name is Pending For Verifiy Access to the Client
func SetMasterControl(w http.ResponseWriter, r *http.Request) {
	log.Println("SetAppVersion (+)")
	origin := r.Header.Get("Origin")
	for _, allowedOrigin := range common.ABHIAllowOrigin {
		if allowedOrigin == origin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			log.Println(origin)
			break
		}
	}
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "CurrentTittle,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	if r.Method == "POST" {
		CurrentTittle := r.Header.Get("CurrentTittle")

		var lMasterReq map[string]interface{}
		var lMasterResp masterResp

		lMasterResp.Status = common.SuccessCode

		//-----------START TO GETTING CLIENT AND STAFF DETAILS--------------
		lClientId := ""
		var lErr1 error
		lClientId, lErr1 = apiaccess.VerifyApiAccess(r, common.ABHIAppName, common.ABHICookieName, "/mastercontrol")
		if lErr1 != nil {
			log.Println("MSMC01", lErr1)
			lMasterResp.Status = common.ErrorCode
			lMasterResp.ErrMsg = "MSMC01" + lErr1.Error()
			fmt.Fprintf(w, helpers.GetErrorString("MSMC01", "UserDetails not Found"))
			return
		} else {
			if lClientId == "" {
				fmt.Fprintf(w, helpers.GetErrorString("MSMC02", "UserDetails not Found"))
				return
			}
		}
		//-----------END OF GETTING CLIENT AND STAFF DETAILS----------------

		lBody, lErr3 := ioutil.ReadAll(r.Body)
		if lErr3 != nil {
			lMasterResp.Status = common.ErrorCode
			lMasterResp.ErrMsg = "MSMC03" + lErr3.Error()
			fmt.Fprintf(w, helpers.GetErrorString("MSMC03", "Unable To Read The Body"))
			return
		} else {
			lErr4 := json.Unmarshal(lBody, &lMasterReq)
			if lErr4 != nil {
				lMasterResp.Status = common.ErrorCode
				lMasterResp.ErrMsg = "MSMC04" + lErr4.Error()
				fmt.Fprintf(w, helpers.GetErrorString("MSMC04", "Unable to unmarshal the Data"))
				return
			} else {
				lErr5 := updateMasterControl(lMasterReq, CurrentTittle)
				if lErr5 != nil {
					lMasterResp.Status = common.ErrorCode
					lMasterResp.ErrMsg = "MSMC05" + lErr5.Error()
					fmt.Fprintf(w, helpers.GetErrorString("MSMC05", "Unable to update  the Data"))
					return
				}

			}
		}
		lData, lErr4 := json.Marshal(lMasterResp)
		if lErr4 != nil {
			log.Println("VGAV05", lErr4)
			fmt.Fprintf(w, helpers.GetErrorString("VGAV05", "Error Occur in Marshalling Datas.."))
			return
		} else {
			fmt.Fprintf(w, string(lData))
		}
	}
}

func updateMasterControl(pMasterData map[string]interface{}, pCurrentTittle string) error {
	log.Println("updateMasterControl (+)")
	log.Println("pMasterData['softDelete']", pMasterData["softDelete"])
	lTableName := ""
	switch {
	case strings.Contains(strings.ToLower(pCurrentTittle), "ncb"):
		lTableName = "a_ncb_master"
	case strings.Contains(strings.ToLower(pCurrentTittle), "ipo"):
		lTableName = "a_ipo_master"
	case strings.Contains(strings.ToLower(pCurrentTittle), "sgb"):
		lTableName = "a_sgb_master"
	}
	log.Println("pMasterData['softDelete']", pMasterData["softDelete"])
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("VUV01", lErr1)
		return lErr1
	}
	defer lDb.Close()

	lCoreString := `UPDATE ` + lTableName + `
        SET SoftDelete = ?
        WHERE id = ?;`

	_, lErr2 := lDb.Exec(lCoreString, pMasterData["softDelete"], pMasterData["id"])
	if lErr2 != nil {
		log.Println("VUV02", lErr2)
		return lErr2
	}

	log.Println("updateMasterControl (-)")
	return nil
}
