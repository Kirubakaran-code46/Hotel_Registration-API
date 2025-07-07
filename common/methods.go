package common

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	tomlread "HOTEL-REGISTRY_API/common/TomlRead"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// --------------------------------------------------------------------
// READ COOKIE
// --------------------------------------------------------------------

func GetCookieValue(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return "", errors.New("cookie not found")
		}
		return "", err
	}
	return cookie.Value, nil
}

// --------------------------------------------------------------------
// function convert the time and date format to customized format
// --------------------------------------------------------------------
func ChangeTimeFormat(pCustomizeLayout string, pInput string) (string, error) {
	log.Println("ChangeTimeFormat (+)")
	var lFormattedValue string

	Layout := ""
	length := len(pInput)
	if length == 19 {
		Layout = "02-01-2006 15:04:05"
	} else if length == 5 {
		Layout = "15:04"
	} else if length == 8 {
		Layout = "15:04:05"
	} else {
		Layout = "02-01-2006 15:04"
	}
	lTimevalue, lErr1 := time.Parse(Layout, pInput)
	if lErr1 != nil {
		log.Println("Error in Parse Timing:", lErr1)
		return lFormattedValue, lErr1
	} else {
		lFormattedValue = lTimevalue.Format(pCustomizeLayout)
	}

	log.Println("ChangeTimeFormat (-)")
	return lFormattedValue, nil
}

func RemoveDuplicateStrings(arr []string) []string {
	uniqueMap := make(map[string]bool)
	result := []string{}

	for _, item := range arr {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			result = append(result, item)
		}
	}

	return result
}

// ----------------------------------------------------------------
// Function to CapitalizeText capitalizes the first letter of each word in a string.
// ----------------------------------------------------------------
func CapitalizeText(input string) string {
	words := strings.Fields(input) // Split the input into words
	var capitalizedWords []string

	for _, word := range words {
		// Capitalize the first letter of the word
		capitalizedWord := strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		capitalizedWords = append(capitalizedWords, capitalizedWord)
	}

	// Join the capitalized words back into a string
	return strings.Join(capitalizedWords, " ")
}

// GENERATE DOCID (Timestamp + Random)
func GenerateDocID() string {
	timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS
	random := rand.Intn(10000)
	return fmt.Sprintf("DOC%s_%04d", timestamp, random)
}

// --------------------------------------------------------------------
// UPLOAD FILE TO THE SPECFIC LOCATION
// --------------------------------------------------------------------

func FilesUpload(pDebug *helpers.HelperStruct, pReq *http.Request, pFileKey string) (string, error) {
	pDebug.Log(helpers.Statement, "FilesUpload (+)")

	if pFileKey == "" {
		return "", nil
	}

	filename := ""
	// Make Directory
	config := tomlread.ReadTomlConfig("toml/ApiCredentials.toml")
	lFilePath := fmt.Sprintf("%v", config.(map[string]interface{})["FilesUpload_Path"])

	if _, lErr := os.Stat(lFilePath); os.IsNotExist(lErr) {
		lErr := os.MkdirAll(lFilePath, os.ModePerm)
		if lErr != nil {

			pDebug.Log(helpers.Elog, "FU001", lErr)
			return filename, lErr
		}
	}
	// Read FormFile
	lFile, handler, lErr := pReq.FormFile(pFileKey)
	if lErr == http.ErrMissingFile {

		// File is optional, skip silently or log
		pDebug.Log(helpers.Elog, "FU002", lErr)
		return filename, nil

	} else if lErr != nil {
		pDebug.Log(helpers.Elog, "FU003", lErr)
		return filename, lErr
	}

	defer lFile.Close()
	//Generate DocId (or) FileName
	lDocId := GenerateDocID()
	ext := filepath.Ext(handler.Filename)
	filename = fmt.Sprintf("%s%s", lDocId, ext)

	filePath := filepath.Join(lFilePath, filename)
	lErr = saveFile(pDebug, lFile, filePath)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "FU004", lErr)
		return filename, lErr
	}

	pDebug.Log(helpers.Statement, "FilesUpload (-)")
	return filename, nil
}

// --------------------------------------------------------------------
// GET FILE WITH DOCID
// --------------------------------------------------------------------

func GetFileBase64(pDebug *helpers.HelperStruct, docID string) (string, error) {
	pDebug.Log(helpers.Statement, "GetFileBase64 (+)")

	if docID == "" {
		pDebug.Log(helpers.Elog, "GFB64", fmt.Errorf("docid empty"))
		return "", fmt.Errorf("docid empty")
	}
	config := tomlread.ReadTomlConfig("toml/ApiCredentials.toml")
	uploadedPath := fmt.Sprintf("%v", config.(map[string]interface{})["FilesUpload_Path"])
	filePath := filepath.Join(uploadedPath, docID) // Adjust this path
	data, lErr := os.ReadFile(filePath)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "GFB64", lErr)
		return "", lErr
	}

	// Detect MIME type from file extension
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream" // fallback
	}

	base64Data := base64.StdEncoding.EncodeToString(data)
	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Data)

	// Add appropriate MIME type if you want, e.g., application/pdf, image/png
	// mimePrefix := "data:application/octet-stream;base64,"
	// pDebug.Log(helpers.Statement, "GetFileBase64 (-)")

	return dataURL, nil
}

// Save File
func saveFile(pDebug *helpers.HelperStruct, file io.Reader, path string) error {
	out, lErr := os.Create(path)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "saveFile01", lErr)
		return lErr
	}
	defer out.Close()
	_, lErr = io.Copy(out, file)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "saveFile02", lErr)
		return lErr
	}
	return nil
}

// --------------------------------------------------------------------
// CHECK UID PREST IN TABLE
// --------------------------------------------------------------------
func CheckUidInTable(pDebug *helpers.HelperStruct, pTableName, pUid string) (bool, error) {
	pDebug.Log(helpers.Statement, "CheckUidInTable (+)")

	var lExist string

	var Status bool

	lQueryString := fmt.Sprintf(`
	SELECT CASE 
		WHEN EXISTS (
			SELECT 1 
			FROM %s 
			WHERE Uid = ? and isActive='Y'
		) 
		THEN 'YES' 
		ELSE 'NO' 
	END AS result;`, pTableName)

	lRows, lErr := database.Gdb.Query(lQueryString, pUid)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "CUID001", lErr.Error())
		return Status, lErr
	}
	for lRows.Next() {
		lErr = lRows.Scan(&lExist)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "CUID002", lErr.Error())
			return Status, lErr
		}
	}
	if strings.EqualFold(lExist, "YES") {
		Status = true
	} else if strings.EqualFold(lExist, "NO") {
		Status = false
	}
	pDebug.Log(helpers.Statement, "CheckUidInTable (+)")
	return Status, nil
}

// --------------------------------------------------------------------
// GET IFSC DETAILS FROM RAZORPAY
// --------------------------------------------------------------------

type IFSCResponse struct {
	MICR     string `json:"MICR"`
	Bank     string `json:"BANK"`
	IFSC     string `json:"IFSC"`
	Branch   string `json:"BRANCH"`
	Address  string `json:"ADDRESS"`
	Contact  string `json:"CONTACT"`
	City     string `json:"CITY"`
	District string `json:"DISTRICT"`
	State    string `json:"STATE"`
	BankCode string `json:"BANKCODE"`
}

func GetIFSCDetails(pDebug *helpers.HelperStruct, pIFSC string, pUid string) (IFSCResponse, error) {
	pDebug.Log(helpers.Statement, "GetIFSCDetails (+)")
	// READ RUL FROM TOML
	config := tomlread.ReadTomlConfig("toml/ApiCredentials.toml")
	lUrl := fmt.Sprintf("%v", config.(map[string]interface{})["RazorPayURL"])

	lFullUrl := lUrl + pIFSC

	var lIfscResp IFSCResponse

	// MAKE APICALL
	lResp, lErr := http.Get(lFullUrl)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "GIFSCD001", fmt.Errorf("HTTP request failed: %w", lErr))
		return lIfscResp, lErr
	}
	defer lResp.Body.Close()

	body, _ := io.ReadAll(lResp.Body)
	lBodyStr := string(body)

	// if lResp.StatusCode != 200 {
	// 	pDebug.Log(helpers.Elog, "GIFSCD002", fmt.Errorf("invalid ifsc or error: %s", lResp.Status))
	// 	return lIfscResp, fmt.Errorf("invalid ifsc or error: %s", lResp.Status)
	// }

	if lResp.StatusCode == http.StatusNotFound {
		pDebug.Log(helpers.Elog, "GIFSCD002", fmt.Errorf("invalid IFSC code"))
		return lIfscResp, fmt.Errorf("invalid IFSC code")
	} else if lResp.StatusCode != http.StatusOK {
		pDebug.Log(helpers.Elog, "GIFSCD003", fmt.Errorf("API returned error: %s", lResp.Status))
		return lIfscResp, fmt.Errorf("API returned error: %s", lResp.Status)
	}

	// UNMARSHAL IFSC RESP STRUCT
	if lErr := json.Unmarshal(body, &lIfscResp); lErr != nil {
		pDebug.Log(helpers.Elog, "GIFSCD004", fmt.Errorf("JSON unmarshal failed: %s", lResp.Status))
		return lIfscResp, fmt.Errorf("JSON unmarshal failed: %v", lErr)
	}

	// INSERT RESPONSE IN TABLE
	lErr = InsertRazorpayResp(pDebug, lResp.StatusCode, lBodyStr, pUid, pIFSC)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "GIFSCD005", lErr.Error())
		return lIfscResp, lErr
	}

	pDebug.Log(helpers.Statement, "GetIFSCDetails (-)")
	return lIfscResp, nil
}

func InsertRazorpayResp(pDebug *helpers.HelperStruct, pStatusCode int, pRespBody, pUid, pIFSC string) error {
	pDebug.Log(helpers.Statement, "InsertRazorpayResp (+)")

	lQueryString := `INSERT INTO razorpay_resplog
					(Uid, ifsc, status_code, respBody, CreatedBy, createdDate)
					VALUES( ?,  ?,  ?,  ?, 'AutoBot', now());`

	_, lErr := database.Gdb.Exec(lQueryString, pUid, pIFSC, pStatusCode, pRespBody)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "IRPAY001", lErr.Error())
		return lErr
	}

	pDebug.Log(helpers.Statement, "InsertRazorpayResp (-)")
	return nil
}
