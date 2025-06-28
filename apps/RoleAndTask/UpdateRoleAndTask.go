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

type RoleAndTaskStruct struct {
	Id          int    `json:"id"`
	Role        string `json:"role"`
	Task        string `json:"task"`
	Router      string `json:"router"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
	IsActive    string `json:"isActive"`
}

type UpdateReqStruct struct {
	Title       string            `json:"title"`
	UpdatedData RoleAndTaskStruct `json:"updatedData"`
}

func UpdateRoleAndTaskAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "UpdateRoleAndTaskAPI (+)")

	if r.Method == http.MethodPost {

		var lReq UpdateReqStruct

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "URTAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("URTAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "URTAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("URTAPI002", lErr.Error()))
			return
		}
		lErr = UpdateRoleTask(lDebug, lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "URTAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("URTAPI003", lErr.Error()))
			return
		}
		fmt.Fprint(w, helpers.GetMsg_String("S", "Updated Successfully"))
	}
	lDebug.Log(helpers.Statement, "UpdateRoleAndTaskAPI (-)")
}

func UpdateRoleTask(pDebug *helpers.HelperStruct, pData UpdateReqStruct) error {
	pDebug.Log(helpers.Statement, "UpdateRoleTask (+)")
	var lCoreString string
	var lErr error
	if strings.EqualFold(pData.Title, "Role Details") {
		lCoreString = `UPDATE role_master
						SET role= ?, Description= ?, UpdatedBy='AutoBot', UpdatedDate= now(), isActive= ?
						WHERE Id= ?`

		_, lErr = database.Gdb.Exec(lCoreString, pData.UpdatedData.Role, pData.UpdatedData.Description, pData.UpdatedData.IsActive, pData.UpdatedData.Id)
	} else if strings.EqualFold(pData.Title, "Task Details") {
		lCoreString = `UPDATE task_master
						SET task= ?, Description= ?,  UpdatedBy='AutoBot', UpdatedDate= now(), router= ?, isActive= ?
						WHERE Id= ?`

		_, lErr = database.Gdb.Exec(lCoreString, pData.UpdatedData.Task, pData.UpdatedData.Description, pData.UpdatedData.Router, pData.UpdatedData.IsActive, pData.UpdatedData.Id)
	} else if strings.EqualFold(pData.Title, "Role and Task Details") {
		lCoreString = `UPDATE role_task_master
						SET roleId=(select id from role_master where role=?) , taskId= (select id from task_master where task=?), UpdatedBy='AutoBot', UpdatedDate= now(), isActive= ?, description= ?
						WHERE Id= ?`

		_, lErr = database.Gdb.Exec(lCoreString, pData.UpdatedData.Task, pData.UpdatedData.IsActive, pData.UpdatedData.Description, pData.UpdatedData.Id)
	}
	if lErr != nil {

		pDebug.Log(helpers.Elog, "URT001 -> Error in -", pData.Title, lErr.Error())
		helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "UpdateRoleTask (-)")
	return nil
}
