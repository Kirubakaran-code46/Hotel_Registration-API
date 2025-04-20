package ftdb

import (
	common "command-line-argumentsD:\\BusinessProject\\SDT_ADMIN_API\\common\\methods.go"
	"fmt"
	"strconv"
)

const (
	PMDB       = "PMDB"
	ADMINDB    = "ADMINDB"
	HOPDB      = "HOPDB"
	SSODB      = "SSODB"
	MariaFTPRD = "MARIAFTPRD"

	BrosePrefix     = "brose.dbo."
	TechExcelPrefix = "TECHEXCELPROD.CAPSFO.dbo."
)

// Initializing DB Details
func (d *AllUsedDatabases) Init() {
	dbconfig := common.ReadTomlConfig("../commonconfig/dbconfig.toml")

	//setting IPO db connection details
	d.IPODB.Server = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["IPODBServer"])
	d.IPODB.Port, _ = strconv.Atoi(fmt.Sprintf("%v", dbconfig.(map[string]interface{})["IPODBPort"]))
	d.IPODB.User = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["IPODBUser"])
	d.IPODB.Password = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["IPODBPassword"])
	d.IPODB.Database = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["IPODBDatabase"])
	d.IPODB.DBType = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["IPODBDBType"])
	d.IPODB.DB = IPODB
	//setting Maria db connection details
	d.MariaDB.Server = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBServer"])
	d.MariaDB.Port, _ = strconv.Atoi(fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBPort"]))
	d.MariaDB.User = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBUser"])
	d.MariaDB.Password = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBPassword"])
	d.MariaDB.Database = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBDatabase"])
	d.MariaDB.DBType = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBDBType"])
	d.MariaDB.DB = MariaFTPRD
}
