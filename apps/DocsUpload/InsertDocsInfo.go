package docsupload

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/common"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func InsertDocsInfoAPI(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	lDebug := new(helpers.HelperStruct)

	lDebug.SetUid(r)
	lDebug.Log(helpers.Statement, "InsertDocsInfoAPI (+)")

	if r.Method == http.MethodPost {
		lErr := r.ParseMultipartForm(32 << 20) // max memory 32MB
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IDIAPI001", lErr.Error())
			fmt.Fprint(w, helpers.GetError_String("IDIAPI001", lErr.Error()))
			return
		}

		var lReq common.DocsUpload

		jsonString := r.FormValue("data")
		lErr = json.Unmarshal([]byte(jsonString), &lReq)

		if lErr != nil {
			lDebug.Log(helpers.Elog, "IDIAPI002", lErr)
			fmt.Fprint(w, helpers.GetError_String("IDIAPI002", lErr.Error()))
			return
		}

		// gstFile, _, lErr := r.FormFile("GST_Docid")
		// if lErr != nil {
		// 	lDebug.Log(helpers.Elog, "IPIAPI001", lErr)
		// 	fmt.Fprint(w, helpers.GetError_String("IPIAPI001", lErr.Error()))
		// 	return
		// }
		// defer gstFile.Close()

		// chequeFile, _, lErr := r.FormFile("cancelledChequeDocid")
		// if lErr != nil {
		// 	lDebug.Log(helpers.Elog, "IPIAPI001", lErr)
		// 	fmt.Fprint(w, helpers.GetError_String("IPIAPI001", lErr.Error()))
		// 	return
		// }
		// defer chequeFile.Close()

		// for _, util := range lReq.Utilities {
		// 	fileKey := util.BillDocid
		// 	if fileKey == "" {
		// 		continue
		// 	}

		// 	file, _, lErr := r.FormFile(fileKey)
		// 	if lErr != nil {
		// 		lDebug.Log(helpers.Elog, "IPIAPI001", lErr)
		// 		fmt.Fprint(w, helpers.GetError_String("IPIAPI001", lErr.Error()))
		// 		return
		// 	}

		// 	fmt.Println("✅ File received for:", fileKey, " (BillType:", util.BillType, ")")
		// 	defer file.Close()
		// }

		// GET UID FROM COOKIE
		lReq.Uid, lErr = common.GetCookieValue(r, common.UIDCOOKIENAME)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IDIAPI003", lErr.Error())
		}

		lErr = InsertDocumentInfo(lDebug, r, lReq)
		if lErr != nil {
			lDebug.Log(helpers.Elog, "IPIAPI004", lErr)
			fmt.Fprint(w, helpers.GetError_String("IPIAPI004", lErr.Error()))
			return
		}
		fmt.Fprint(w, helpers.GetMsg_String("S", "Inserted Successfully"))
	}
	lDebug.Log(helpers.Statement, "InsertDocsInfoAPI (-)")
}

func InsertDocumentInfo(pDebug *helpers.HelperStruct, r *http.Request, pReq common.DocsUpload) error {
	pDebug.Log(helpers.Statement, "InsertDocumentInfo (+)")

	// lExist, lErr := common.CheckUidInTable(pDebug, "document_upload", pReq.Uid)
	// if lErr != nil {
	// 	pDebug.Log(helpers.Elog, "IDI001", lErr.Error())
	// 	return helpers.ErrReturn(lErr)
	// }

	// UPDATE LOCATION INFO
	// if !lExist {
	// lErr = UpdateLocationDetails(pDebug, pReq)
	// if lErr != nil {
	// 	pDebug.Log(helpers.Elog, "ILD002", lErr.Error())
	// 	return helpers.ErrReturn(lErr)
	// }
	// 	fmt.Println("kjhk")
	// } else {

	var gstDocID string
	var chequeDocID string
	var lErr error
	// insert GST File
	if strings.EqualFold(pReq.GST_Docid, "") {
		gstDocID, lErr = common.FilesUpload(pDebug, r, "GST_Docid")
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IDI001", lErr)
			return helpers.ErrReturn(lErr)
		}

	} else {
		gstDocID = pReq.GST_Docid
	}

	// insert Cancellation Cheque File
	// if pReq.CancelledChequeDocid != "" {
	if strings.EqualFold(pReq.CancelledChequeDocid, "") {
		chequeDocID, lErr = common.FilesUpload(pDebug, r, "cancelledChequeDocid")
		if lErr != nil {

			pDebug.Log(helpers.Elog, "IDI002", lErr)
			return helpers.ErrReturn(lErr)
		}
	} else {
		chequeDocID = pReq.CancelledChequeDocid
	}
	lExist, lErr := common.CheckUidInTable(pDebug, "document_upload", pReq.Uid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IDI003", lErr.Error())
		return helpers.ErrReturn(lErr)
	}

	if lExist {
		// SOFT DELETE OLD VALUES
		lQueryString := `update document_upload set isActive='N' where Uid=?`

		_, lErr = database.Gdb.Exec(lQueryString, pReq.Uid)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IDI004", lErr.Error())
			return helpers.ErrReturn(lErr)
		}

		// SOFT DELETE OLD VALUES
		lQueryString = `update utility_types set isActive='N' where Uid=?`

		_, lErr = database.Gdb.Exec(lQueryString, pReq.Uid)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IDI005", lErr.Error())
			return helpers.ErrReturn(lErr)
		}
	}

	// INSERT DOCUMENT_UPLOAD TABLE

	lQueryString := `INSERT INTO document_upload
(Uid, Bank_Name, Account_Number, Acc_HolderName, IFSC_code, Branch, GST_Number, GST_docId, cancelledCheque_docId,propertyOwnership,Start_Date,End_Date, isActive, CreatedBy, createdDate, UpdatedBy, UpdatedDate)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?,'Y', 'AutoBot', now(),'AutoBot', now());`

	_, lErr = database.Gdb.Exec(lQueryString, pReq.Uid, pReq.BankName, pReq.AccountNumber, pReq.AccHolderName, pReq.IFSC_Code, pReq.Branch, pReq.GST_Number, gstDocID, chequeDocID, pReq.PropertyOwnership, pReq.StartDate, pReq.EndDate)

	if lErr != nil {
		pDebug.Log(helpers.Elog, "IDI006", lErr)
		return helpers.ErrReturn(lErr)
	}

	// INSERT UTILITY TYPES TABLE

	for index, utility := range pReq.Utilities {

		var utilityDocId string

		// check if the utility contains new file
		if strings.HasPrefix(utility.BillDocid, "billDocid_") {

			fileField := fmt.Sprintf("billDocid_%d", index)
			// Attempt to upload file
			utilityDocId, lErr = common.FilesUpload(pDebug, r, fileField)
			if lErr != nil {
				pDebug.Log(helpers.Elog, "IDI007", lErr)
				return helpers.ErrReturn(lErr)
			}
		} else {
			utilityDocId = utility.BillDocid
		}

		lQueryString = `INSERT INTO utility_types
						(Uid, Bill_Type, Bill_docId, isActive, CreatedBy, createdDate, UpdatedBy, UpdatedDate)
						VALUES( ?, ?, ?, 'Y', 'AutoBot', now(),'AutoBot', now());`

		_, lErr = database.Gdb.Exec(lQueryString, pReq.Uid, utility.BillType, utilityDocId)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "IDI008", lErr)
			return helpers.ErrReturn(lErr)
		}
	}

	pDebug.Log(helpers.Statement, "InsertDocumentInfo (-)")
	return nil
}
