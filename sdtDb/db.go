package sdtdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Structure to hold database connection details
type DatabaseType struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
	DBType   string
	DB       string
}

// structure to hold all db connection details used in this program
type AllUsedDatabases struct {
	SDTDB DatabaseType
}

var Gdb *sql.DB

// ---------------------------------------------------------------------------------
// function opens the db connection and return connection variable
// ---------------------------------------------------------------------------------
func InitDb() {
	DbDetails := new(AllUsedDatabases)
	DbDetails.Init()

	var lErr error
	var dataBaseConnection DatabaseType
	// get connection details
	if DbDetails.SDTDB.DB == "SDTDB" {
		dataBaseConnection = DbDetails.SDTDB
	}
	// log.Println("localDBtype", localDBtype)
	// log.Println("dataBaseConnection", dataBaseConnection)

	// Prepare connection string

	prepareConnString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dataBaseConnection.User, dataBaseConnection.Password, dataBaseConnection.Server, dataBaseConnection.Port, dataBaseConnection.Database)

	Gdb, lErr = sql.Open("mysql", prepareConnString)
	if lErr != nil {
		log.Fatalf("Error opening database: %v", lErr)
	}

	// Test connection
	if lErr := Gdb.Ping(); lErr != nil {
		log.Fatalf("Error connecting to database: %v", lErr)
	}
}

// --------------------------------------------------------------------
//
//	execute bulk inserts
//
// --------------------------------------------------------------------
func ExecuteBulkStatement(db *sql.DB, sqlStringValues string, sqlString string) error {
	log.Println("ExecuteBulkStatement+")
	//trim the last ,
	sqlStringValues = sqlStringValues[0 : len(sqlStringValues)-1]
	_, err := db.Exec(sqlString + sqlStringValues)
	if err != nil {
		log.Println(err)
		log.Println("ExecuteBulkStatement-")
		return err
	} else {
		log.Println("inserted Sucessfully")
	}
	log.Println("ExecuteBulkStatement-")
	return nil
}
