package mealsinfo

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	locationinfo "HOTEL-REGISTRY_API/apps/LocationInfo"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func InsertMealsInfoAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "InsertMealsInfoAPI (+)")

	if r.Method == http.MethodPost {

		var lReq common.MealsInfo

		lBody, lErr := io.ReadAll(r.Body)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IMIAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IMIAPI001", lErr.Error()))
			return
		}

		lErr = json.Unmarshal(lBody, &lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IMIAPI002", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IMIAPI002", lErr.Error()))
			return
		}

		// GET UID FROM COOKIE
		lReq.Uid, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IMIAPI003", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IMIAPI003", lErr.Error()))
			return
		}

		lErr = InsertMealsDetails(lDebug, lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IMIAPI004", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IMIAPI004", lErr.Error()))
			return
		}

	}
	fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	lDebug.Log(helpers.Statement, "InsertMealsInfoAPI (-)")
}

func InsertMealsDetails(pDebug *helpers.HelperStruct, pReq common.MealsInfo) error {
	pDebug.Log(helpers.Statement, "InsertMealsDetails (+)")

	lExist, lErr := locationinfo.CheckUidInTable(pDebug, "meals_info", pReq.Uid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IMD001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	// UPDATE MEALS INFO
	if lExist {
		lErr = UpdateMealsDetails(pDebug, pReq)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IMD002", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	} else {
		var lCoreString string
		lTypesOfMealsArr := strings.Join(pReq.TypesOfMeals, ",")

		// INSERT MEALS INFO

		if pReq.IsOperationalRestaurant == "No" {
			lCoreString = `INSERT INTO meals_info
							(Uid, isOptionalRestaurant,CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
							VALUES(?, ?,'AutoBot', now(), 'AutoBot', now(), 'Y');`
			_, lErr = database.Gdb.Exec(lCoreString, pReq.Uid, pReq.IsOperationalRestaurant)
		} else {
			lCoreString = `INSERT INTO meals_info
	(Uid, isOptionalRestaurant, Meal_Package, Types_Of_Meals, Meal_Rack_Price, CreatedBy, createdDate, UpdatedBy, UpdatedDate, isActive)
	VALUES(?, ?, ?, ?, ?,'AutoBot', now(), 'AutoBot', now(), 'Y');`
			_, lErr = database.Gdb.Exec(lCoreString, pReq.Uid, pReq.IsOperationalRestaurant, pReq.MealPackage, lTypesOfMealsArr, pReq.MealRackPrice)
		}
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IMD003", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "InsertMealsDetails (-)")
	return nil

}

func UpdateMealsDetails(pDebug *helpers.HelperStruct, pReq common.MealsInfo) error {
	pDebug.Log(helpers.Statement, "UpdateMealsDetails (+)")
	lTypesOfMealsArr := strings.Join(pReq.TypesOfMeals, ",")
	var lCoreString string
	var lErr error

	if pReq.IsOperationalRestaurant == "No" {
		lCoreString = `UPDATE meals_info
						SET isOptionalRestaurant= ?, Meal_Package='', Types_Of_Meals='', Meal_Rack_Price='', UpdatedBy='AutoBot', UpdatedDate=now() WHERE Uid = ?;`

		_, lErr = database.Gdb.Exec(lCoreString, pReq.IsOperationalRestaurant, pReq.Uid)
	} else {
		lCoreString = `UPDATE meals_info
	SET  isOptionalRestaurant= ?, Meal_Package= ?, Types_Of_Meals= ?, Meal_Rack_Price= ?, UpdatedBy='AutoBot', UpdatedDate= now()
	WHERE Uid = ?;`
		_, lErr = database.Gdb.Exec(lCoreString, pReq.IsOperationalRestaurant, pReq.MealPackage, lTypesOfMealsArr, pReq.MealRackPrice, pReq.Uid)
	}

	if lErr != nil {
		pDebug.Log(helpers.Elog, "UMD001", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	pDebug.Log(helpers.Statement, "UpdateMealsDetails (-)")
	return nil

}
