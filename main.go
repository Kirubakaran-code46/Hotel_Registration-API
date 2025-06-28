package main

import (
	apientry "HOTEL-REGISTRY_API/ApiEntry"
	sdtdb "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/apps"
	availabilityinfo "HOTEL-REGISTRY_API/apps/AvailabilityInfo"
	basicinfo "HOTEL-REGISTRY_API/apps/BasicInfo"
	getuserdetails "HOTEL-REGISTRY_API/apps/GetUserDetails"
	locationinfo "HOTEL-REGISTRY_API/apps/LocationInfo"
	mealsinfo "HOTEL-REGISTRY_API/apps/MealsInfo"
	policyinfo "HOTEL-REGISTRY_API/apps/PolicyInfo"
	roomtypes "HOTEL-REGISTRY_API/apps/RoomTypes"
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
	router.HandleFunc("/getPropertyTypes", basicinfo.GetBasicInfoDropdown).Methods(http.MethodGet)
	router.HandleFunc("/getStates", locationinfo.GetStateDropdown).Methods(http.MethodGet)
	router.HandleFunc("/getRoomDropdown", roomtypes.GetRoomTypesDropdown).Methods(http.MethodGet)
	router.HandleFunc("/insertBasicInfo", basicinfo.InsertBasicDetailsAPI).Methods(http.MethodPost)
	router.HandleFunc("/insertLocationInfo", locationinfo.InsertLocationDetailsAPI).Methods(http.MethodPost)
	router.HandleFunc("/getUserInfo", getuserdetails.GetUserDetailsAPI).Methods(http.MethodPost)
	router.HandleFunc("/inserRoomDetails", roomtypes.InsertRoomDetailsAPI).Methods(http.MethodPost)
	router.HandleFunc("/insertMealsInfo", mealsinfo.InsertMealsInfoAPI).Methods(http.MethodPost)
	router.HandleFunc("/insertAvailability", availabilityinfo.InsertAvailabilityDetailsAPI).Methods(http.MethodPost)
	// Policies
	router.HandleFunc("/getPoliciesDropdown", policyinfo.GetPoliciesInfoDropdown).Methods(http.MethodGet)
	router.HandleFunc("/insertPropertyPolicies", policyinfo.InsertPolicyInfoAPI).Methods(http.MethodPost)

	return router
}
