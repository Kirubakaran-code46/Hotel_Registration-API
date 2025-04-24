package main

import (
	apientry "SDT_ADMIN_API/ApiEntry"
	"SDT_ADMIN_API/apps"
	dashboard "SDT_ADMIN_API/apps/Dashboard"
	roleandtask "SDT_ADMIN_API/apps/RoleAndTask"
	sdtdb "SDT_ADMIN_API/sdtDb"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	lLog, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("error opening file: %v", lErr)
	}
	defer lLog.Close()
	log.SetOutput(lLog)

	sdtdb.InitDb()

	defer sdtdb.Gdb.Close()

	router := CreateRouter()

	handler := apientry.APIMiddleware(router)
	PortNo := ":8081"

	//handler=router.
	srv := &http.Server{
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
		Addr:         PortNo,
	}

	log.Println("Server is Started In Port " + PortNo + "...")
	fmt.Println("Server is Started In Port " + PortNo + "...")

	log.Fatal(srv.ListenAndServe())

}

func CreateRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/getEmpDetails", apps.GetAllEmpDetailsAPI).Methods(http.MethodGet)
	router.HandleFunc("/getDashboardDetails", dashboard.GetDashBoardDetailsAPI).Methods(http.MethodGet)
	router.HandleFunc("/getRoleTaskDetails", roleandtask.GetRoleAndTaskAPI).Methods(http.MethodGet)

	return router
}
