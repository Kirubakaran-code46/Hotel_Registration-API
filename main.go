package main

import (
	apientry "HOTEL-REGISTRY_API/ApiEntry"
	GDB "HOTEL-REGISTRY_API/Db_Setup"
	availabilityinfo "HOTEL-REGISTRY_API/apps/AvailabilityInfo"
	basicinfo "HOTEL-REGISTRY_API/apps/BasicInfo"
	description "HOTEL-REGISTRY_API/apps/Description"
	docsupload "HOTEL-REGISTRY_API/apps/DocsUpload"
	getuserdetails "HOTEL-REGISTRY_API/apps/GetUserDetails"
	locationinfo "HOTEL-REGISTRY_API/apps/LocationInfo"
	mealsinfo "HOTEL-REGISTRY_API/apps/MealsInfo"
	policyinfo "HOTEL-REGISTRY_API/apps/PolicyInfo"
	propertydetails "HOTEL-REGISTRY_API/apps/PropertyDetails"
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

	GDB.InitDb()

	defer GDB.Gdb.Close()
	defer GDB.AdminGdb.Close()

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
	// GET USER INFO
	router.HandleFunc("/getUserInfo", getuserdetails.GetUserDetailsAPI).Methods(http.MethodPost)

	// Basic Info
	router.HandleFunc("/getPropertyTypes", basicinfo.GetBasicInfoDropdown).Methods(http.MethodGet)
	router.HandleFunc("/insertBasicInfo", basicinfo.InsertBasicDetailsAPI).Methods(http.MethodPost)

	// Location Info
	router.HandleFunc("/getStates", locationinfo.GetStateDropdown).Methods(http.MethodGet)
	router.HandleFunc("/insertLocationInfo", locationinfo.InsertLocationDetailsAPI).Methods(http.MethodPost)

	// Room Info
	router.HandleFunc("/getRoomDropdown", roomtypes.GetRoomTypesDropdown).Methods(http.MethodGet)
	router.HandleFunc("/inserRoomDetails", roomtypes.InsertRoomDetailsAPI).Methods(http.MethodPost)

	// Meals Info
	router.HandleFunc("/insertMealsInfo", mealsinfo.InsertMealsInfoAPI).Methods(http.MethodPost)

	// Availability Info
	router.HandleFunc("/insertAvailability", availabilityinfo.InsertAvailabilityDetailsAPI).Methods(http.MethodPost)

	// Policies
	router.HandleFunc("/getPoliciesDropdown", policyinfo.GetPoliciesInfoDropdown).Methods(http.MethodGet)
	router.HandleFunc("/insertPropertyPolicies", policyinfo.InsertPolicyInfoAPI).Methods(http.MethodPost)

	// DocsInfo
	router.HandleFunc("/insertDocsInfo", docsupload.InsertDocsInfoAPI).Methods(http.MethodPost)
	router.HandleFunc("/getIFSCdetails", docsupload.GetIFSCdetailsAPI).Methods(http.MethodPost)

	// PropertyInfo
	router.HandleFunc("/insertPropertyInfo", propertydetails.InsertPropertyInfoAPI).Methods(http.MethodPost)

	// DescriptionInfo
	router.HandleFunc("/insertDescInfo", description.InsertDescInfoAPI).Methods(http.MethodPost)

	// Clear Session
	router.HandleFunc("/clearSession", getuserdetails.ClearCookieAPI).Methods(http.MethodGet)

	return router
}
