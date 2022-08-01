package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// TODO Use Google Sheets as DB.
// TODO Use Redis for caching.
const (
	spreadsheetId = "SHEET_ID_HERE"
)

func CheckShortPath(path string) bool {
	const readRange string = "A"
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

	for _, row := range response.Values {
		fmt.Printf("%s \n", row[0])
	}

	return false
}

func SaveShortPath(path string, shortUrl string, longUrl string) bool {
	return false
}
