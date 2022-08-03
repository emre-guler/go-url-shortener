package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// TODO Use Google Sheets as DB.
// TODO Use Redis for caching.
// TODO Use Environment Variables for credentials and SpreadsheetId
const (
	spreadsheetId = "1P7Me-PLTskt4v-LLDxnCu92X30Z0p7M4DhzWLS6hDF4"
)

func CheckShortPath(path string) bool {
	const readRange string = "A:A" // Select's short paths
	ctx := context.Background()

	pwd, _ := os.Getwd()
	credentialsJson, _ := ioutil.ReadFile(pwd + "/db/credentials.json")
	service, serviceError := sheets.NewService(ctx, option.WithCredentialsJSON(credentialsJson))

	if serviceError != nil {
		fmt.Println("Unable connect to service, please try again later!")
		return false
	}

	response, responseError := service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	if responseError != nil {
		fmt.Println("Unable access to db, please try again later!")
		fmt.Println(responseError)
		return false
	}

	var result bool = true
	println("Db record => ", len(response.Values))
	for _, row := range response.Values {
		if row[0] == path {
			result = false
			break
		}
	}
	return result
}

func SaveShortPath(path string, shortUrl string, longUrl string) bool {
	const sheetRange string = "A1"
	ctx := context.Background()

	pwd, _ := os.Getwd()
	credentialsJson, _ := ioutil.ReadFile(pwd + "/db/credentials.json")
	service, serviceError := sheets.NewService(ctx, option.WithCredentialsJSON(credentialsJson))

	if serviceError != nil {
		fmt.Println("Unable connect to service, please try again later!")
		return false
	}

	var names []interface{}
	names = append(names, path)

	var valueRange sheets.ValueRange
	valueRange.MajorDimension = strings.ToUpper("ROWS")
	// todo valueRange.Values = names

	response, err := service.Spreadsheets.Values.Append(spreadsheetId, sheetRange, &valueRange).ValueInputOption("USER_ENTERED").Fields("*").Do()

	fmt.Println(response)
	fmt.Println(err)

	return true
}
