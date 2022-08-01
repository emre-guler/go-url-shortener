package db

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"os"
)

// TODO Use Google Sheets as DB.
// TODO Use Redis for caching.
const (
	spreadsheetId = "SHEET_ID_HERE"
	readRange     = "A:B"
)

func CheckShortPath(path string) bool {
	ctx := context.Background()
	pwd, _ := os.Getwd()
	jsonFile, _ := ioutil.ReadFile(pwd + "/db/credentials.json")
	service, serviceError := sheets.NewService(ctx, option.WithCredentialsJSON(jsonFile))

	if serviceError != nil {
		fmt.Println("Unable to connect service, please try again later!1")
		return false
	}

	var deneme = service.Spreadsheets.Values.Get(spreadsheetId, readRange)
	response, responseError := deneme.Do()

	if responseError != nil {
		fmt.Println("Unable to connect service, please try again later!2")
		fmt.Println(responseError)
		return false
	}

	if len(response.Values) == 0 {
		fmt.Print("No data found.")
		return false
	} else {
		for _, row := range response.Values {
			fmt.Printf("%s, %s\n", row[0], row[1])
		}
	}

	return false
}

func SaveShortPath(path string, shortUrl string, longUrl string) bool {
	return false
}
