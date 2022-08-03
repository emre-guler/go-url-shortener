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
	const sheetRange string = "A:B"
	ctx := context.Background()

	pwd, _ := os.Getwd()
	credentialsJson, _ := ioutil.ReadFile(pwd + "/db/credentials.json")
	service, serviceError := sheets.NewService(ctx, option.WithCredentialsJSON(credentialsJson))

	if serviceError != nil {
		fmt.Println("Unable connect to service, please try again later!")
		return false
	}

	var names [][]interface{}
	var names2 []interface{}
	names2 = append(names2, path)
	names2 = append(names2, longUrl)
	names = append(names, names2)

	var valueRange sheets.ValueRange
	valueRange.MajorDimension = strings.ToUpper("ROWS")
	valueRange.Values = names

	_, err := service.Spreadsheets.Values.Append(spreadsheetId, sheetRange, &valueRange).ValueInputOption("USER_ENTERED").Fields("*").Do()

	if err != nil {
		fmt.Println("Something went wrong, please try again later.")
	}

	return true
}
