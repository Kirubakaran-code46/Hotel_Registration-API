package coresettings

import (
	"SDT_ADMIN_API/common"
	database "SDT_ADMIN_API/sdtdb"
	"log"
)

// --------------------------------------------------------------------
// function to get value from core setting for a given Key
// --------------------------------------------------------------------
func GetCoreSettingValue(dbName string, key string) string {
	method := "coresettings.GetCoreSettingValue"
	var value string
	sqlString := "select valuev from CoreSettings where keyv ='" + key + "'"
	rows, err := database.Gdb.Query(sqlString)
	if err != nil {
		log.Println(common.NoPanic, method, err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&value)
		if err != nil {
			log.Println(common.NoPanic, method, err.Error())
		}
	}
	return value

}
