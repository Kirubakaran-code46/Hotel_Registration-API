package roleandtask

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"net/http"
)

type DropdownResponse struct {
	Status string   `json:"status"`
	ErrMsg string   `json:"errMsg"`
	Role   []string `json:"role"`
	Task   []string `json:"task"`
}

func GetRoleAndTaskDropdownAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetRoleAndTaskDropdownAPI (+)")

	if r.Method == http.MethodGet {
		var lResponse DropdownResponse
		lResponse.Status = common.SUCCESSCODE

		var lErr error

		lResponse.Role, lErr = GetRoleDropDwon(lDebug)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRATDD001", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}

		lResponse.Task, lErr = GetTaskDropDwon(lDebug)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRATDD002", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}

		lResp, lErr := json.Marshal(&lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRATDD003", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}
		fmt.Fprint(w, string(lResp))

	}
	lDebug.Log(helpers.Statement, "GetRoleAndTaskDropdownAPI (-)")
}

func GetRoleDropDwon(pDebug *helpers.HelperStruct) ([]string, error) {
	pDebug.Log(helpers.Statement, "GetDropDwon (+)")
	var lRole string
	var lRoleArr []string
	lCoreString := `select role from role_master rm where rm.isActive='Y'`

	lRows, lErr := database.Gdb.Query(lCoreString)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GRDD001", lErr.Error())
		return lRoleArr, lErr
	}
	for lRows.Next() {
		lErr = lRows.Scan(&lRole)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GRDD002", lErr.Error())
			return lRoleArr, lErr
		}
		lRoleArr = append(lRoleArr, lRole)
	}

	pDebug.Log(helpers.Statement, "GetDropDwon (-)")
	return lRoleArr, nil
}

func GetTaskDropDwon(pDebug *helpers.HelperStruct) ([]string, error) {
	pDebug.Log(helpers.Statement, "GetDropDwon (+)")
	var lTask string
	var lTaskArr []string
	lCoreString := `select task from task_master tm where tm.isActive='Y'`

	lRows, lErr := database.Gdb.Query(lCoreString)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GRDD001", lErr.Error())
		return lTaskArr, lErr
	}
	for lRows.Next() {
		lErr = lRows.Scan(&lTask)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GRDD002", lErr.Error())
			return lTaskArr, lErr
		}
		lTaskArr = append(lTaskArr, lTask)
	}

	pDebug.Log(helpers.Statement, "GetDropDwon (-)")
	return lTaskArr, nil
}
