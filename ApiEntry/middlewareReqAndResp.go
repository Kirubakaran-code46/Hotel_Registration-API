package apientry

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	"HOTEL-REGISTRY_API/helpers"
	"net/http"
	"strings"
)

// MIDDLEWARE REQUEST CAPTURE

func LogRequest(pDebug *helpers.HelperStruct, pToken string, pReqDtl APIRequestStruct, pRequestID string) {
	pDebug.Log(helpers.Statement, "LogRequest (+)")

	//insert token
	if strings.EqualFold(pReqDtl.RequestType, "multipart/form-data") {
		pReqDtl.Body = "file"
	}
	lQueryString := "insert into middleware_req_log(request_id,token,requesteddate,realip,forwardedip,method,path,host,remoteaddr,header,body,endpoint) values (?,?,NOW(),?,?,?,?,?,?,?,?,?)"
	_, lErr := database.Gdb.Exec(lQueryString, pRequestID, pToken, pReqDtl.RealIP, pReqDtl.ForwardedIP, pReqDtl.Method, pReqDtl.Path, pReqDtl.Host, pReqDtl.RemoteAddr, pReqDtl.Header, helpers.ReplaceBase64String(pReqDtl.Body, 0), pReqDtl.EndPoint)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "MIDDLEWARE REQUEST CAPTURE ERROR", lErr.Error())
	}

	pDebug.Log(helpers.Statement, "LogRequest (-)")

}

// MIDDLEWARE RESPONSE CAPTURE

func LogResponse(pDebug *helpers.HelperStruct, pReq *http.Request, pRespStatus int, pRespData []byte, pRequestID string) {
	pDebug.Log(helpers.Statement, "LogResponse (+)")

	lReqDtl := GetAPIRequestDetail(pDebug, pReq)
	//insert token
	lQueryString := "insert into middleware_resp_log(request_id,response,responseStatus,requesteddate,realip,forwardedip,method,path,host,remoteaddr,header,body,endpoint) values (?,?,?,NOW(),?,?,?,?,?,?,?,?,?)"
	_, lErr := database.Gdb.Exec(lQueryString, pRequestID, helpers.ReplaceBase64String(string(pRespData), 0), pRespStatus, lReqDtl.RealIP, lReqDtl.ForwardedIP, lReqDtl.Method, lReqDtl.Path, lReqDtl.Host, lReqDtl.RemoteAddr, lReqDtl.Header, helpers.ReplaceBase64String(lReqDtl.Body, 0), lReqDtl.EndPoint)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "MIDDLEWARE RESPONSE CAPTURE ERROR", lErr.Error())
	}

	pDebug.Log(helpers.Statement, "LogResponse (-)")

}
