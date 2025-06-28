package dbsetup

import (
	"HOTEL-REGISTRY_API/common"
	"fmt"
	"strconv"
)

const (
	SDTDB = "SDTDB"
)

// Initializing DB Details
func (d *AllUsedDatabases) Init() {
	dbconfig := common.ReadTomlConfig("./toml/Dbconfig.toml")

	//setting IPO db connection details
	d.SDTDB.Server = fmt.Sprintf("%v", dbconfig.(map[string]any)["HOST"])
	d.SDTDB.Port, _ = strconv.Atoi(fmt.Sprintf("%v", dbconfig.(map[string]any)["PORT"]))
	d.SDTDB.User = fmt.Sprintf("%v", dbconfig.(map[string]any)["USER"])
	d.SDTDB.Password = fmt.Sprintf("%v", dbconfig.(map[string]any)["PASSWORD"])
	d.SDTDB.Database = fmt.Sprintf("%v", dbconfig.(map[string]any)["DBNAME"])
	d.SDTDB.DBType = fmt.Sprintf("%v", dbconfig.(map[string]any)["DBTYPE"])
	d.SDTDB.DB = SDTDB
}
