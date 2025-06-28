package roleandtask

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

type ReqStruct struct {
	Title           string          `json:"title"`
	RoleDetails     RoleDetails     `json:"roleDetails"`
	TaskDetails     TaskDetails     `json:"taskDetails"`
	RoleTaskDetails RoleTaskDetails `json:"roleTaskDetails"`
}

func AddRoleAndTaskAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "AddRoleAndTaskAPI (+)")

	if r.Method == http.MethodPost {

		var lReq ReqStruct

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "ARTAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("ARTAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "ARTAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("ARTAPI002", lErr.Error()))
			return
		}

		if strings.EqualFold(lReq.Title, "Role Details") {

			lErr = AddRoleDetails(lDebug, lReq.RoleDetails)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "ARTAPI003", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("ARTAPI003", lErr.Error()))
				return
			}
		} else if strings.EqualFold(lReq.Title, "Task Details") {

			lErr = AddTaskDetails(lDebug, lReq.TaskDetails)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "ARTAPI004", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("ARTAPI004", lErr.Error()))
				return
			}
		} else if strings.EqualFold(lReq.Title, "Role and Task Details") {

			lErr = AddRoleAndTaskDetails(lDebug, lReq.RoleTaskDetails)
			if lErr != nil {
				lDebug.Log(helpers.Elog, "ARTAPI005", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("ARTAPI005", lErr.Error()))
				return
			}
		}
		fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	}
	lDebug.Log(helpers.Statement, "AddRoleAndTaskAPI (-)")
}

func AddRoleDetails(pDebug *helpers.HelperStruct, pRoleData RoleDetails) error {
	pDebug.Log(helpers.Statement, "AddRoleDetails (+)")
	lCoreString := `INSERT INTO role_master
					(role, Description, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
					VALUES(?,?,"AutoBot",now(),"AutoBot",now(),?);`

	_, lErr := database.Gdb.Exec(lCoreString, pRoleData.Role, pRoleData.Description, pRoleData.IsActive)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "ARD001", lErr.Error())
		helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "AddRoleDetails (-)")
	return nil
}

func AddTaskDetails(pDebug *helpers.HelperStruct, pTaskData TaskDetails) error {
	pDebug.Log(helpers.Statement, "AddTaskDetails (+)")
	lCoreString := `INSERT INTO task_master
					(task, Description,router, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
					VALUES(?,?,?,"AutoBot",now(),"AutoBot",now(),?);`

	_, lErr := database.Gdb.Exec(lCoreString, pTaskData.Task, pTaskData.Description, pTaskData.Router, pTaskData.IsActive)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "ATD001", lErr.Error())
		helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "AddTaskDetails (-)")
	return nil
}

func AddRoleAndTaskDetails(pDebug *helpers.HelperStruct, pRoleTaskData RoleTaskDetails) error {
	pDebug.Log(helpers.Statement, "AddRoleAndTaskDetails (+)")
	lCoreString := `INSERT INTO role_task_master
					(roleId, taskId, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive, description)
					VALUES((select id from role_master where role=?), (select id from task_master where task=?),'AutoBot',now(),'AutoBot',now(), ?, ?);`

	_, lErr := database.Gdb.Exec(lCoreString, pRoleTaskData.Role, pRoleTaskData.Task, pRoleTaskData.IsActive, pRoleTaskData.Description)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "ARD001", lErr.Error())
		helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "AddRoleAndTaskDetails (-)")
	return nil
}
