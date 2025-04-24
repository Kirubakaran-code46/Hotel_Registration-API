package apps

import (
	"SDT_ADMIN_API/helpers"
	database "SDT_ADMIN_API/sdtDb"
	"encoding/json"

	"fmt"
	"net/http"
)

type Response struct {
	Status     string       `json:"status"`
	ErrMsg     string       `json:"errMsg"`
	EmpDetails []EmpDetails `json:"empDetails"`
}

type EmpDetails struct {
	Id          int    `json:"id"`
	EmpId       string `json:"empid"`
	EmpName     string `json:"empname"`
	Designation string `json:"designation"`
}

func GetAllEmpDetailsAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "GetAllEmpDetailsAPI (+)")

	if r.Method == http.MethodGet {
		var lResponse Response
		var lEmpDetails EmpDetails
		var lEmpDetailsArr []EmpDetails

		lResponse.Status = "S"

		lCoreString := `SELECT id, emp_id, emp_name, designation
                       FROM emp_details where isActive='Y' order by id desc`

		lRows, lErr := database.Gdb.Query(lCoreString)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "GAE001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GAE001", lErr.Error()))
			return
		}

		for lRows.Next() {
			lErr = lRows.Scan(&lEmpDetails.Id, &lEmpDetails.EmpId, &lEmpDetails.EmpName, &lEmpDetails.Designation)

			if lErr != nil {
				lDebug.Log(helpers.Elog, "GAE002", lErr.Error())
				fmt.Fprint(w, helpers.GetError_String("GAE002", lErr.Error()))
				return
			}
			lEmpDetailsArr = append(lEmpDetailsArr, lEmpDetails)
		}
		lResponse.EmpDetails = lEmpDetailsArr

		lData, lErr := json.Marshal(lResponse)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "GAE003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("GAE003", lErr.Error()))
			return
		}
		fmt.Fprint(w, string(lData))

	}
	lDebug.Log(helpers.Statement, "GetAllEmpDetailsAPI (-)")
}
