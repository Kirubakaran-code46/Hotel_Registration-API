package apientry

import (
	"HOTEL-REGISTRY_API/helpers"
	"net/http"
)

type APIRequestStruct struct {
	RealIP      string
	ForwardedIP string
	Method      string
	Path        string
	Host        string
	RemoteAddr  string
	Header      string
	Body        string
	EndPoint    string
	RequestType string
}

func GetAPIRequestDetail(pDebug *helpers.HelperStruct, r *http.Request) APIRequestStruct {
	pDebug.Log(helpers.Statement, "GetAPIRequestDetail (+)")

	var lReqDetails APIRequestStruct
	lReqDetails.RealIP = r.Header.Get("Referer")
	lReqDetails.ForwardedIP = r.Header.Get("X-Forwarded-For")
	lReqDetails.Method = r.Method
	lReqDetails.Path = r.URL.Path + "?" + r.URL.RawQuery
	lReqDetails.Host = r.Host
	lReqDetails.RemoteAddr = r.RemoteAddr
	lReqDetails.EndPoint = r.URL.Path
	lReqDetails.RequestType = r.Header.Get("Content-Type")

	lReqDetails.Header = GetHeaderDetails(pDebug, r)
	pDebug.Log(helpers.Statement, "GetAPIRequestDetail (-)")

	return lReqDetails
}

func GetHeaderDetails(pDebug *helpers.HelperStruct, r *http.Request) string {
	pDebug.Log(helpers.Statement, "GetHeaderDetails+")
	lValue := ""
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			lValue = lValue + " " + name + "-" + value
		}
	}
	pDebug.Log(helpers.Statement, "GetHeaderDetails-")
	return lValue
}
