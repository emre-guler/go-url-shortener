package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
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

const shortUrl string = "https://www.emreguler.dev/"

func authentication() *sheets.Service {
	ctx := context.Background()
	pwd, _ := os.Getwd()
	credentialsJson, _ := ioutil.ReadFile(pwd + "/db/credentials.json")
	service, _ := sheets.NewService(ctx, option.WithCredentialsJSON(credentialsJson))
	return service
}

func connectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func CheckShortPath(path string, service *sheets.Service) bool {
	rdb := connectRedis()
	cachedValue, _ := rdb.Get(context.Background(), ("path_key_" + path)).Result()

	if cachedValue == "" {
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
	return false
}

func SaveShortPath(path string, redirectUrl string) bool {
	service := authentication()
	if service != nil {
		if CheckShortPath(path, service) {
			var sheetValues [][]interface{}
			var values []interface{}
			values = append(values, path, redirectUrl)
			sheetValues = append(sheetValues, values)

			_, err := service.Spreadsheets.Values.Append(spreadsheetId, "A:B", &sheets.ValueRange{
				MajorDimension: strings.ToUpper("ROWS"),
				Values:         sheetValues,
			}).ValueInputOption("USER_ENTERED").Fields("*").Do()

			if err != nil {
				fmt.Println(err)
				return false
			}
			rdb := connectRedis()
			rdb.Set(context.Background(), ("path_key_" + path), redirectUrl, 3*time.Hour)
			return true
		}
	}
	return false
}
