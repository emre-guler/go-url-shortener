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

// TODO 1- > Use Google Sheets as DB.
// * Finished ✔️
// TODO 2- > Use Environment Variables for SpreadsheetId.
// * Finished ✔️
// TODO 3 -> Dockerize.
// * Finished ✔️
// TODO 3 -> Use Redis for caching.

var spreadsheetId string = os.Getenv("URL_SHORTENER_PROJECT_ SPREADSHEET_ID")

func authentication() *sheets.Service {
	ctx := context.Background()
	pwd, _ := os.Getwd()
	credentialsJson, _ := ioutil.ReadFile(pwd + "/db/credentials.json")
	service, _ := sheets.NewService(ctx, option.WithCredentialsJSON(credentialsJson))
	return service
}

func CheckShortPath(path string, service *sheets.Service) bool {
	response, responseError := service.Spreadsheets.Values.Get(spreadsheetId, "A:A").Do()
	if responseError != nil {
		return false
	}
	var result bool = true
	for _, row := range response.Values {
		if row[0] == path {
			result = false
			break
		}
	}
	return result
}

func SaveShortPath(path string, shortUrl string, longUrl string) bool {
	service := authentication()
	if service != nil {
		if CheckShortPath(path, service) {
			var sheetValues [][]interface{}
			var values []interface{}
			values = append(values, path, longUrl)
			sheetValues = append(sheetValues, values)

			_, err := service.Spreadsheets.Values.Append(spreadsheetId, "A:B", &sheets.ValueRange{
				MajorDimension: strings.ToUpper("ROWS"),
				Values:         sheetValues,
			}).ValueInputOption("USER_ENTERED").Fields("*").Do()

			if err != nil {
				fmt.Println(err)
				return false
			}
			fmt.Println("Your new link is: ", longUrl)
			return true
		}
	}
	return false
}
