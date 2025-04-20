package mastercontrol

import (
	"encoding/json"
	"fcs23pkg/apps/validation/apiaccess"
	"fcs23pkg/common"
	"fcs23pkg/ftdb"
	"fcs23pkg/helpers"
	"fmt"
	"log"
	"net/http"
)

type masterdataResp struct {
	IpoMasterData []ipoMaster `json:"ipoMasterData"`
	SgbMasterData []sgbMaster `json:"sgbMasterData"`
	NcbMasterData []ncbMaster `json:"ncbMasterData"`
	Status        string      `json:"status"`
	ErrMsg        string      `json:"errMsg"`
}

type ipoMaster struct {
	Id               int     `json:"id"`
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	BiddingStartDate string  `json:"biddingStartDate"`
	BiddingEndDate   string  `json:"biddingEndDate"`
	DailyStartTime   string  `json:"dailyStartTime"`
	DailyEndTime     string  `json:"dailyEndTime"`
	MaxPrice         float64 `json:"maxPrice"`
	MinPrice         float64 `json:"minPrice"`
	MinBidQuantity   int     `json:"minBidQuantity"`
	LotSize          int     `json:"lotSize"`
	Registrar        string  `json:"registrar"`
	T1ModStartDate   string  `json:"t1ModStartDate"`
	T1ModStartTime   string  `json:"t1ModStartTime"`
	T1ModEndTime     string  `json:"t1ModEndTime"`
	T1ModEndDate     string  `json:"t1ModEndDate"`
	TickSize         float64 `json:"tickSize"`
	FaceValue        float64 `json:"faceValue"`
	IssueSize        float64 `json:"issueSize"`
	CutOffPrice      float64 `json:"cutOffPrice"`
	Isin             string  `json:"isin"`
	IssueType        string  `json:"issueType"`
	SubType          string  `json:"subType"`
	Exchange         string  `json:"exchange"`
	SoftDelete       string  `json:"softDelete"`
}

type sgbMaster struct {
	Id                   int     `json:"id"`
	Symbol               string  `json:"symbol"`
	Series               string  `json:"series"`
	Name                 string  `json:"name"`
	IssueType            string  `json:"issueType"`
	LotSize              int     `json:"lotSize"`
	FaceValue            float64 `json:"faceValue"`
	MinBidQuantity       int     `json:"minBidQuantity"`
	MinPrice             float64 `json:"minPrice"`
	MaxPrice             float64 `json:"maxPrice"`
	TickSize             float64 `json:"tickSize"`
	BiddingStartDate     string  `json:"biddingStartDate"`
	BiddingEndDate       string  `json:"biddingEndDate"`
	DailyStartTime       string  `json:"dailyStartTime"`
	DailyEndTime         string  `json:"dailyEndTime"`
	T1ModStartDate       string  `json:"t1ModStartDate"`
	T1ModEndDate         string  `json:"t1ModEndDate"`
	T1ModStartTime       string  `json:"t1ModStartTime"`
	T1ModEndTime         string  `json:"t1ModEndTime"`
	Isin                 string  `json:"isin"`
	IssueSize            int     `json:"issueSize"`
	IssueValueSize       int     `json:"issueValueSize"`
	MaxQuantity          int     `json:"maxQuantity"`
	AllotmentDate        string  `json:"allotmentDate"`
	IncompleteModEndDate string  `json:"incompleteModEndDate"`
	Exchange             string  `json:"exchange"`
	Redemption           string  `json:"redemption"`
	SoftDelete           string  `json:"softDelete"`
}
type ncbMaster struct {
	Id                    int     `json:"id"`
	Symbol                string  `json:"symbol"`
	Series                string  `json:"series"`
	Name                  string  `json:"name"`
	LotSize               int     `json:"lotSize"`
	FaceValue             float64 `json:"faceValue"`
	MinBidQuantity        int     `json:"minBidQuantity"`
	MinPrice              float64 `json:"minPrice"`
	MaxPrice              float64 `json:"maxPrice"`
	TickSize              float64 `json:"tickSize"`
	CutOffPrice           float64 `json:"cutOffPrice"`
	BiddingStartDate      string  `json:"biddingStartDate"`
	BiddingEndDate        string  `json:"biddingEndDate"`
	DailyStartTime        string  `json:"dailyStartTime"`
	DailyEndTime          string  `json:"dailyEndTime"`
	T1ModStartDate        string  `json:"t1ModStartDate"`
	T1ModEndDate          string  `json:"t1ModEndDate"`
	T1ModStartTime        string  `json:"t1ModStartTime"`
	T1ModEndTime          string  `json:"t1ModEndTime"`
	Isin                  string  `json:"isin"`
	IssueSize             float64 `json:"issueSize"`
	IssueValueSize        int     `json:"issueValueSize"`
	MaxQuantity           string  `json:"maxQuantity"`
	AllotmentDate         string  `json:"allotmentDate"`
	LastDayBiddingEndTime string  `json:"lastDayBiddingEndTime"`
	Exchange              string  `json:"exchange"`
	RBIName               string  `json:"rbiName"`
	SoftDelete            string  `json:"softDelete"`
}

func GetAllMaster(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	for _, allowedOrigin := range common.ABHIAllowOrigin {
		if allowedOrigin == origin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			log.Println(origin)
			break
		}
	}
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	if r.Method == "GET" {
		log.Println("GetAllMasterData (+)")
		var lGetMasterData masterdataResp
		lGetMasterData.Status = common.SuccessCode

		//-----------START TO GETTING CLIENT AND STAFF DETAILS--------------
		lClientId := ""
		var lErr1 error
		lClientId, lErr1 = apiaccess.VerifyApiAccess(r, common.ABHIAppName, common.ABHICookieName, "/mastercontrol")
		if lErr1 != nil {
			log.Println("MGAM01", lErr1)
			lGetMasterData.Status = common.ErrorCode
			lGetMasterData.ErrMsg = "MGAM01" + lErr1.Error()
			fmt.Fprintf(w, helpers.GetErrorString("MGAM01", "UserDetails not Found"))
			return
		} else {
			if lClientId == "" {
				fmt.Fprintf(w, helpers.GetErrorString("MGAM02", "UserDetails not Found"))
				return
			}
		}
		//-----------END OF GETTING CLIENT AND STAFF DETAILS----------------
		lIpoMasterArr, lErr3 := getIpoMaster()
		if lErr3 != nil {
			log.Println("MGAM03", lErr3)
			lGetMasterData.Status = common.ErrorCode
			lGetMasterData.ErrMsg = "MGAM03" + lErr3.Error()
			fmt.Fprintf(w, helpers.GetErrorString("MGAM03", "Unable to Get an IPO Master Details"))
			return
		} else {
			lSgbMasterArr, lErr4 := getSgbMaster()
			if lErr4 != nil {
				log.Println("MGAM04", lErr4)
				lGetMasterData.Status = common.ErrorCode
				lGetMasterData.ErrMsg = "MGAM04" + lErr4.Error()
				fmt.Fprintf(w, helpers.GetErrorString("MGAM04", "Unable to Get an SGB Master Details"))
				return
			} else {
				lNcbMasterArr, lErr5 := getNcbMaster()
				if lErr5 != nil {
					log.Println("MGAM05", lErr5)
					lGetMasterData.Status = common.ErrorCode
					lGetMasterData.ErrMsg = "MGAM05" + lErr5.Error()
					fmt.Fprintf(w, helpers.GetErrorString("MGAM05", "Unable to Get an NCB Master Details"))
					return
				} else {
					lGetMasterData.IpoMasterData = lIpoMasterArr
					lGetMasterData.SgbMasterData = lSgbMasterArr
					lGetMasterData.NcbMasterData = lNcbMasterArr

				}
			}
		}

		lData, lErr4 := json.Marshal(lGetMasterData)
		if lErr4 != nil {
			log.Println("VGAV05", lErr4)
			fmt.Fprintf(w, helpers.GetErrorString("VGAV05", "Error Occur in Marshalling Datas.."))
			return
		} else {
			fmt.Fprintf(w, string(lData))
		}
		log.Println("GetAllMasterData (-)")
	}
}

func getIpoMaster() ([]ipoMaster, error) {
	log.Println("GetIpoMaster (+)")
	var lIpoRec ipoMaster
	var lIpoMasterArr []ipoMaster
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("MGIM01", lErr1)
		return lIpoMasterArr, lErr1
	} else {
		defer lDb.Close()
		lCoreString := `SELECT nvl(Id,0), nvl(Symbol,''),nvl(Name,'') ,nvl(BiddingStartDate,'') ,nvl(BiddingEndDate,'')  
		,nvl(DailyStartTime,'')  ,nvl(DailyEndTime,'')  , nvl(MaxPrice,0) , nvl(MinPrice,0) ,nvl(MinBidQuantity,0)  , nvl(LotSize,0) 
		,nvl(Registrar,'')  ,nvl(T1ModStartDate,'')  , nvl(T1ModStartTime,'') , nvl(T1ModEndTime,'') , nvl(T1ModEndDate,'') ,
		nvl(TickSize,0) , nvl(FaceValue,0)  ,  nvl(IssueSize,0) ,  nvl(CutOffPrice,0) , 
		nvl(Isin,'') ,  nvl(IssueType,'') , nvl(SubType,'')  ,  nvl(Exchange,''),nvl(SoftDelete,'N') 
			FROM a_ipo_master order by Id desc`
		lRows, lErr2 := lDb.Query(lCoreString)
		if lErr2 != nil {
			log.Println("MGIM02", lErr2)
			return lIpoMasterArr, lErr2
		} else {
			for lRows.Next() {
				lErr3 := lRows.Scan(&lIpoRec.Id, &lIpoRec.Symbol, &lIpoRec.Name, &lIpoRec.BiddingStartDate, &lIpoRec.BiddingEndDate, &lIpoRec.DailyStartTime, &lIpoRec.DailyEndTime, &lIpoRec.MaxPrice, &lIpoRec.MinPrice, &lIpoRec.MinBidQuantity, &lIpoRec.LotSize, &lIpoRec.Registrar, &lIpoRec.T1ModStartDate, &lIpoRec.T1ModStartTime, &lIpoRec.T1ModEndTime, &lIpoRec.T1ModEndDate, &lIpoRec.TickSize, &lIpoRec.FaceValue, &lIpoRec.IssueSize, &lIpoRec.CutOffPrice, &lIpoRec.Isin, &lIpoRec.IssueType, &lIpoRec.SubType, &lIpoRec.Exchange, &lIpoRec.SoftDelete)
				if lErr3 != nil {
					log.Println("MGIM03", lErr3)
					return lIpoMasterArr, lErr3
				} else {
					lIpoMasterArr = append(lIpoMasterArr, lIpoRec)
				}
			}
		}
	}
	log.Println("GetIpoMaster(-) ")
	return lIpoMasterArr, nil

}

func getSgbMaster() ([]sgbMaster, error) {
	log.Println("getSgbMaster (+)")
	var lSgbRec sgbMaster
	var lSgbMasterArr []sgbMaster
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("MGSM01", lErr1)
		return lSgbMasterArr, lErr1
	} else {
		defer lDb.Close()
		lCoreString := `SELECT nvl(id,0),nvl(Symbol,''),nvl(Series,''),
		nvl(Name,''),nvl(IssueType,'') ,nvl(Lotsize,0) ,
		nvl(FaceValue,0),nvl(MinBidQuantity,0) ,
		nvl(MinPrice,0), nvl(FaceValue,0),nvl(TickSize,0) , 
	   nvl(BiddingStartDate,'') ,nvl(BiddingEndDate,'') ,nvl(DailyStartTime,'') ,nvl(DailyEndTime,'') ,
		nvl(T1ModStartDate,''),nvl(T1ModEndDate,'') ,nvl(T1ModStartTime,'') ,nvl(T1ModEndTime,'') , 
		nvl(Isin,'') ,nvl(IssueSize,0) ,nvl(IssueValueSize,0) , 
	   nvl(AllotmentDate,'') ,nvl(IncompleteModEndDate,'') ,nvl(Exchange,'') ,nvl(Redemption,''),nvl(SoftDelete,'N') 
	   FROM a_sgb_master order by id desc`
		lRows, lErr2 := lDb.Query(lCoreString)
		if lErr2 != nil {
			log.Println("MGSM02", lErr2)
			return lSgbMasterArr, lErr2
		} else {
			for lRows.Next() {
				lErr3 := lRows.Scan(&lSgbRec.Id, &lSgbRec.Symbol, &lSgbRec.Series, &lSgbRec.Name, &lSgbRec.IssueType, &lSgbRec.LotSize, &lSgbRec.FaceValue, &lSgbRec.MinBidQuantity, &lSgbRec.MinPrice, &lSgbRec.FaceValue, &lSgbRec.TickSize, &lSgbRec.BiddingStartDate, &lSgbRec.BiddingEndDate, &lSgbRec.DailyStartTime, &lSgbRec.DailyEndTime, &lSgbRec.T1ModStartDate, &lSgbRec.T1ModEndDate, &lSgbRec.T1ModStartTime, &lSgbRec.T1ModEndTime, &lSgbRec.Isin, &lSgbRec.IssueSize, &lSgbRec.IssueValueSize, &lSgbRec.AllotmentDate, &lSgbRec.IncompleteModEndDate, &lSgbRec.Exchange, &lSgbRec.Redemption, &lSgbRec.SoftDelete)
				if lErr3 != nil {
					log.Println("MGSM03", lErr3)
					return lSgbMasterArr, lErr3
				} else {
					lSgbMasterArr = append(lSgbMasterArr, lSgbRec)
				}
			}
		}
	}
	log.Println("getSgbMaster(-) ")
	return lSgbMasterArr, nil

}

func getNcbMaster() ([]ncbMaster, error) {
	log.Println("getNcbMaster (+)")
	var lNcbRec ncbMaster
	var lNcbMasterArr []ncbMaster
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("MGNM01", lErr1)
		return lNcbMasterArr, lErr1
	} else {
		defer lDb.Close()
		lCoreString := `	SELECT nvl(id,0),nvl(Symbol,'') ,nvl(Series,'') , 
		nvl(Name,''),nvl(Lotsize,0) ,nvl(FaceValue,0) ,
		 nvl(MinBidQuantity,0) ,nvl(MinPrice,0) ,nvl(MaxPrice,0) ,
		nvl(TickSize,0) ,nvl(CutoffPrice,0) ,nvl(BiddingStartDate,'') ,
		nvl(BiddingEndDate,'') ,nvl(DailyStartTime,'') ,nvl(DailyEndTime,'') ,
		nvl(T1ModStartDate,'') ,nvl(T1ModEndDate,'') ,nvl(T1ModStartTime,'') ,
		nvl(T1ModEndTime,'') ,nvl(Isin,'') ,nvl(IssueSize,0) ,nvl(IssueValueSize,0) , 
		nvl(MaxQuantity,0),nvl(AllotmentDate,'') ,nvl( LastDayBiddingEndTime,''), 
		nvl(Exchange,''),nvl(RbiName,''),nvl(SoftDelete,'N')
	FROM a_ncb_master
	order by id desc`
		lRows, lErr2 := lDb.Query(lCoreString)
		if lErr2 != nil {
			log.Println("MGSM02", lErr2)
			return lNcbMasterArr, lErr2
		} else {
			for lRows.Next() {
				lErr3 := lRows.Scan(&lNcbRec.Id, &lNcbRec.Symbol, &lNcbRec.Series, &lNcbRec.Name, &lNcbRec.LotSize, &lNcbRec.FaceValue, &lNcbRec.MinBidQuantity, &lNcbRec.MinPrice, &lNcbRec.MaxPrice, &lNcbRec.TickSize, &lNcbRec.CutOffPrice, &lNcbRec.BiddingStartDate, &lNcbRec.BiddingEndDate, &lNcbRec.DailyStartTime, &lNcbRec.DailyEndTime, &lNcbRec.T1ModStartDate, &lNcbRec.T1ModEndDate, &lNcbRec.T1ModStartTime, &lNcbRec.T1ModEndTime, &lNcbRec.Isin, &lNcbRec.IssueSize, &lNcbRec.IssueValueSize, &lNcbRec.MaxQuantity, &lNcbRec.AllotmentDate, &lNcbRec.LastDayBiddingEndTime, &lNcbRec.Exchange, &lNcbRec.RBIName, &lNcbRec.SoftDelete)
				if lErr3 != nil {
					log.Println("MGNM03", lErr3)
					return lNcbMasterArr, lErr3
				} else {
					lNcbMasterArr = append(lNcbMasterArr, lNcbRec)
				}
			}
		}
	}
	log.Println("getNcbMaster(-) ")
	return lNcbMasterArr, nil

}
