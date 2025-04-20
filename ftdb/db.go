package ftdb

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
	PM    DatabaseType
	ADMIN DatabaseType
	HOP   DatabaseType
}

// ---------------------------------------------------------------------------------
// function opens the db connection and return connection variable
// ---------------------------------------------------------------------------------
func LocalDbConnect(DBtype string) (*sql.DB, error) {
	DbDetails := new(AllUsedDatabases)
	DbDetails.Init()

	var db *sql.DB
	var err error
	var dataBaseConnection DatabaseType
	// get connection details
	if DBtype == DbDetails.PM.DB {
		dataBaseConnection = DbDetails.PM
	} else if DBtype == DbDetails.HOP.DB {
		dataBaseConnection = DbDetails.HOP
	} else if DBtype == DbDetails.ADMIN.DB {
		dataBaseConnection = DbDetails.ADMIN
	}
	// log.Println("localDBtype", localDBtype)
	// log.Println("dataBaseConnection", dataBaseConnection)

	// Prepare connection string

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dataBaseConnection.User, dataBaseConnection.Password, dataBaseConnection.Server, dataBaseConnection.Port, dataBaseConnection.Database)

	return db, err
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
