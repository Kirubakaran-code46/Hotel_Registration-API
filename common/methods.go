package common

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

// --------------------------------------------------------------------
// function reads the constants from the config.toml file
// --------------------------------------------------------------------
func ReadTomlConfig(filename string) interface{} {
	var f interface{}
	if _, err := toml.DecodeFile(filename, &f); err != nil {
		log.Println(err)
	}
	return f
}

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

//----------------------------------------------------------------
// Creating custom error
// ----------------------------------------------------------------

func CustomError(pErrorMsg string) error {
	err := errors.New(pErrorMsg)
	return err
}
