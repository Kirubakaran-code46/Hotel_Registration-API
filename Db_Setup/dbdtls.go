package dbsetup

import (
	tomlread "HOTEL-REGISTRY_API/common/TomlRead"
	"fmt"
	"strconv"
)

const (
	REGISTRYDB = "REGISTRYDB"
	ADMINDB    = "ADMINDB"
)

// Initializing DB Details
func (d *AllUsedDatabases) Init() {
	dbconfig := tomlread.ReadTomlConfig("./toml/Dbconfig.toml")

	//setting hotel registry connection details
	d.REGISTRYDB.Server = fmt.Sprintf("%v", dbconfig.(map[string]any)["REGHOST"])
	d.REGISTRYDB.Port, _ = strconv.Atoi(fmt.Sprintf("%v", dbconfig.(map[string]any)["REGPORT"]))
	d.REGISTRYDB.User = fmt.Sprintf("%v", dbconfig.(map[string]any)["REGUSER"])
	d.REGISTRYDB.Password = fmt.Sprintf("%v", dbconfig.(map[string]any)["REGPASSWORD"])
	d.REGISTRYDB.Database = fmt.Sprintf("%v", dbconfig.(map[string]any)["REGDBNAME"])
	d.REGISTRYDB.DBType = fmt.Sprintf("%v", dbconfig.(map[string]any)["REGDBTYPE"])
	d.REGISTRYDB.DB = REGISTRYDB

	//setting admin process connection details
	d.ADMINDB.Server = fmt.Sprintf("%v", dbconfig.(map[string]any)["ADMHOST"])
	d.ADMINDB.Port, _ = strconv.Atoi(fmt.Sprintf("%v", dbconfig.(map[string]any)["ADMPORT"]))
	d.ADMINDB.User = fmt.Sprintf("%v", dbconfig.(map[string]any)["ADMUSER"])
	d.ADMINDB.Password = fmt.Sprintf("%v", dbconfig.(map[string]any)["ADMPASSWORD"])
	d.ADMINDB.Database = fmt.Sprintf("%v", dbconfig.(map[string]any)["ADMDBNAME"])
	d.ADMINDB.DBType = fmt.Sprintf("%v", dbconfig.(map[string]any)["ADMDBTYPE"])
	d.ADMINDB.DB = ADMINDB

}
