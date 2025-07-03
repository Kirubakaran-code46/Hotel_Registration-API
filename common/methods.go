package common

import (
	database "HOTEL-REGISTRY_API/Db_Setup"
	tomlread "HOTEL-REGISTRY_API/common/TomlRead"
	"HOTEL-REGISTRY_API/helpers"
	"encoding/base64"
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
// function reads the constants from the config.toml file
// --------------------------------------------------------------------

// --------------------------------------------------------------------
// READ COOKIE
// --------------------------------------------------------------------

// GetCookieValue reads a cookie by name from the request and returns its value.
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

// GET BASE64 FILE
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
		pDebug.Log(helpers.Elog, "ILD001", lErr.Error())
		return Status, lErr
	}
	for lRows.Next() {
		lErr = lRows.Scan(&lExist)
		if lErr != nil {
			pDebug.Log(helpers.Elog, "ILD002", lErr.Error())
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
