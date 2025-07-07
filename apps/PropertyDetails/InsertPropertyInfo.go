package propertydetails

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	s3filehandler "HOTEL-REGISTRY_API/common/S3FileHandler"

	"HOTEL-REGISTRY_API/helpers"
	"fmt"
	"net/http"
)

func InsertPropertyInfoAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "InsertPropertyInfoAPI (+)")

	if r.Method == http.MethodPost {
		lErr := r.ParseMultipartForm(32 << 20) // max memory 32MB
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IPIAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IPIAPI001", lErr.Error()))
			return
		}

		lErr = InsertPropertyDetails(lDebug, r)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IPIAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IPIAPI001", lErr.Error()))
			return
		}
		fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	}
	lDebug.Log(helpers.Statement, "InsertPropertyInfoAPI (-)")
}

func InsertPropertyDetails(pDebug *helpers.HelperStruct, pReq *http.Request) error {
	pDebug.Log(helpers.Statement, "InsertPropertyDetails (+)")

	// GET UID FROM COOKIE
	lUid, lErr := common.GetCookieValue(pReq, common.UIDCOOKIENAME)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IPD001", lErr)
		return helpers.ErrReturn(lErr)
	}

	lFileKeys := []string{"Facade_docId", "Parking_docId", "Lobby_docId", "Reception_docId", "Corridors_docId", "Lift_docId", "Bathroom_docId", "OtherArea_docId", "PropertyImg_docId"}

	for _, lKey := range lFileKeys {

		// Separate check for file key presence
		fileHeaders, filePresent := pReq.MultipartForm.File[lKey]

		if filePresent {
			if len(fileHeaders) > 0 {
				// File Upload
				// lFileDocId, lErr := common.FilesUpload(pDebug, pReq, lKey)
				lFileDocId, lErr := s3filehandler.S3FileUpload(pDebug, pReq, lKey)

				if lErr != nil {
					pDebug.Log(helpers.Elog, "IDI001", lErr)
					return helpers.ErrReturn(lErr)
				}
				lErr = InsertDocIdInTable(pDebug, lFileDocId, lKey, lUid)
				if lErr != nil {
					pDebug.Log(helpers.Elog, "IDI001", lErr)
					return helpers.ErrReturn(lErr)
				}
			}
		}

	}

	pDebug.Log(helpers.Statement, "InsertPropertyDetails (-)")
	return nil
}

func InsertDocIdInTable(pDebug *helpers.HelperStruct, pFileDocId, pColumnKey, pUid string) error {
	pDebug.Log(helpers.Statement, "InsertDocIdInTable (+)")
	// CHECK UID ALREADY PRESENT OR NOT
	lExist, lErr := common.CheckUidInTable(pDebug, "property_details", pUid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IDI003", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	if lExist {
		// UPDATE DATA
		lQueryString := fmt.Sprintf(`UPDATE property_details
		SET %s=?, UpdatedBy='AutoBot', UpdatedDate= now()
		WHERE Uid = ? and isActive='Y'`, pColumnKey)

		_, lErr = database.Gdb.Exec(lQueryString, pFileDocId, pUid)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IDI003", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	} else {
		// INSERT DATA

		lQueryString := fmt.Sprintf(`INSERT INTO property_details
		(Uid, %s, isActive, CreatedBy, createdDate, UpdatedBy, UpdatedDate)
		VALUES(?,?, 'Y', 'AutoBot', now(), 'AutoBot', now());`, pColumnKey)

		_, lErr = database.Gdb.Exec(lQueryString, pUid, pFileDocId)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IDI003", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	}
	pDebug.Log(helpers.Statement, "InsertDocIdInTable (-)")
	return nil
}
