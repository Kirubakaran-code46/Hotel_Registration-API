package roleandtask

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status          string            `json:"status"`
	ErrMsg          string            `json:"errMsg"`
	RoleDetails     []RoleDetails     `json:"roleDetails"`
	TaskDetails     []TaskDetails     `json:"taskDetails"`
	RoleTaskDetails []RoleTaskDetails `json:"roletaskDetails"`
}

type RoleDetails struct {
	Id          int    `json:"id"`
	Role        string `json:"role"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
	IsActive    string `json:"isActive"`
}

type TaskDetails struct {
	Id          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
	Router      string `json:"router"`
	IsActive    string `json:"isActive"`
}

type RoleTaskDetails struct {
	Id          int    `json:"id"`
	Role        string `json:"role"`
	Task        string `json:"task"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
	IsActive    string `json:"isActive"`
}

func GetRoleAndTaskAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetRoleAndTaskAPI (+)")

	if r.Method == http.MethodGet {
		var lResponse Response

		lResponse.Status = common.SUCCESSCODE

		lRoleDetails, lErr := GetRoleDetails(lDebug)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRATAPI001", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}
		lResponse.RoleDetails = lRoleDetails

		lTaskDetails, lErr := GetTaskDetails(lDebug)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRATAPI002", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}
		lResponse.TaskDetails = lTaskDetails

		lRoleTaskDetails, lErr := GetRoleTaskDetails(lDebug)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GRATAPI003", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}
		lResponse.RoleTaskDetails = lRoleTaskDetails

		lResp, lErr := json.Marshal(&lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GDDAPI004", lErr.Error())
			lResponse.Status = common.ERRORCODE
			lResponse.ErrMsg = lErr.Error()
		}
		fmt.Fprint(w, string(lResp))
	}
	lDebug.Log(helpers.Statement, "GetRoleAndTaskAPI (-)")
}

func GetRoleDetails(pDebug *helpers.HelperStruct) ([]RoleDetails, error) {
	pDebug.Log(helpers.Statement, "GetRoleDetails (+)")

	var lRoleDetails RoleDetails
	var lRoleDetailsArr []RoleDetails

	lCoreString := `SELECT Id, role, Description, CreatedBy, createdDate, UpdatedBy, UpdatedDate,isActive
					FROM role_master;`
	lRows, lErr := database.Gdb.Query(lCoreString)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GRD001", lErr.Error())
		return lRoleDetailsArr, lErr
	}

	for lRows.Next() {
		lErr = lRows.Scan(&lRoleDetails.Id, &lRoleDetails.Role, &lRoleDetails.Description, &lRoleDetails.CreatedBy, &lRoleDetails.CreatedDate, &lRoleDetails.UpdatedBy, &lRoleDetails.UpdatedDate, &lRoleDetails.IsActive)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GRD002", lErr.Error())
			return lRoleDetailsArr, lErr
		}
		lRoleDetailsArr = append(lRoleDetailsArr, lRoleDetails)
	}

	pDebug.Log(helpers.Statement, "GetRoleDetails (-)")
	return lRoleDetailsArr, nil
}

func GetTaskDetails(pDebug *helpers.HelperStruct) ([]TaskDetails, error) {
	pDebug.Log(helpers.Statement, "GetTaskDetails (+)")

	var lTaskDetails TaskDetails
	var lTaskDetailsArr []TaskDetails

	lCoreString := `SELECT Id, task, Description, CreatedBy, createdDate, UpdatedBy, UpdatedDate, router, isActive
					FROM task_master`
	lRows, lErr := database.Gdb.Query(lCoreString)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GRD001", lErr.Error())
		return lTaskDetailsArr, lErr
	}

	for lRows.Next() {
		lErr = lRows.Scan(&lTaskDetails.Id, &lTaskDetails.Task, &lTaskDetails.Description, &lTaskDetails.CreatedBy, &lTaskDetails.CreatedDate, &lTaskDetails.UpdatedBy, &lTaskDetails.UpdatedDate, &lTaskDetails.Router, &lTaskDetails.IsActive)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GRD002", lErr.Error())
			return lTaskDetailsArr, lErr
		}
		lTaskDetailsArr = append(lTaskDetailsArr, lTaskDetails)
	}

	pDebug.Log(helpers.Statement, "GetTaskDetails (-)")
	return lTaskDetailsArr, nil
}

func GetRoleTaskDetails(pDebug *helpers.HelperStruct) ([]RoleTaskDetails, error) {
	pDebug.Log(helpers.Statement, "GetRoleTaskDetails (+)")

	var lRoleTaskDetails RoleTaskDetails
	var lRoleTaskDetailsArr []RoleTaskDetails

	lCoreString := `SELECT rtm.Id, rm.role, tm.task,ifnull(rtm.description,""), rtm.CreatedBy, rtm.createdDate, rtm.UpdatedBy, rtm.UpdatedDate, rtm.isActive FROM role_task_master rtm join role_master rm on rtm.roleId = rm.Id join task_master tm on rtm.taskId =tm.id`
	lRows, lErr := database.Gdb.Query(lCoreString)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "GRTD001", lErr.Error())
		return lRoleTaskDetailsArr, lErr
	}

	for lRows.Next() {
		lErr = lRows.Scan(&lRoleTaskDetails.Id, &lRoleTaskDetails.Role, &lRoleTaskDetails.Task, &lRoleTaskDetails.Description, &lRoleTaskDetails.CreatedBy, &lRoleTaskDetails.CreatedDate, &lRoleTaskDetails.UpdatedBy, &lRoleTaskDetails.UpdatedDate, &lRoleTaskDetails.IsActive)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "GRTD002", lErr.Error())
			return lRoleTaskDetailsArr, lErr
		}
		lRoleTaskDetailsArr = append(lRoleTaskDetailsArr, lRoleTaskDetails)
	}

	pDebug.Log(helpers.Statement, "GetRoleTaskDetails (-)")
	return lRoleTaskDetailsArr, nil
}
