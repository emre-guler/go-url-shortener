package db

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// TODO Use Google Sheets as DB.
// TODO Use Redis for caching.
const (
	spreadsheetId = "SHEET_ID_HERE"
	readRange     = "A:B"
)

func CheckShortPath(path string) bool {
	ctx := context.Background()
	service, serviceError := sheets.NewService(ctx, option.WithAPIKey("YOUR_API_KEY_HERE"))

	if serviceError != nil {
		fmt.Println("Unable to connect service, please try again later!1")
		return false
	}

	response, responsError := service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	if responsError != nil {
		fmt.Println("Unable to connect service, please try again later!2")
		fmt.Println(responsError)
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
