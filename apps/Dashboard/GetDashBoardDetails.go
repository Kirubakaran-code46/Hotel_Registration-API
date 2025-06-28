package dashboard

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      string             `json:"status"`
	ErrMsg      string             `json:"errMsg"`
	DashDetails []DashboardDetails `json:"dashDetails"`
}

type DashboardDetails struct {
	Id          int    `json:"id"`
	Role        string `json:"role"`
	Description string `json:"description"`
	Router      string `json:"router"`
}

func GetDashBoardDetailsAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetDashBoardDetailsAPI (+)")

	if r.Method == http.MethodGet {
		var lResponse Response

		lResponse.Status = common.SUCCESSCODE

		lDetails, lErr := GetDashDetails(lDebug)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GDDAPI001", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}
		lResponse.DashDetails = lDetails

		lResp, lErr := json.Marshal(&lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GDDAPI002", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}
		fmt.Fprint(w, string(lResp))
	}
	lDebug.Log(helpers.Statement, "GetDashBoardDetailsAPI (-)")
}

func GetDashDetails(pDebug *helpers.HelperStruct) ([]DashboardDetails, error) {
	pDebug.Log(helpers.Statement, "GetDashDetails (+)")

	var lDashDetails DashboardDetails
	var lDashDetailsArr []DashboardDetails

	lCoreString := `SELECT Id,role, ifnull(Description,""),ifnull(router,"")
			FROM role_management
			WHERE isActive='Y';`
	lRows, lErr := database.Gdb.Query(lCoreString)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GDD001", lErr.Error())
		return lDashDetailsArr, lErr
	}

	for lRows.Next() {
		lErr = lRows.Scan(&lDashDetails.Id, &lDashDetails.Role, &lDashDetails.Description, &lDashDetails.Router)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GDD002", lErr.Error())
			return lDashDetailsArr, lErr
		}
		lDashDetailsArr = append(lDashDetailsArr, lDashDetails)
	}

	pDebug.Log(helpers.Statement, "GetDashDetails (-)")
	return lDashDetailsArr, nil
}
