package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/emre-guler/url-shortener/db"
)

const shortUrl string = "https://www.emreguler.dev/"

func main() {
	fmt.Println("Welcome to url-shortener!")
	fmt.Println("---------------------")
	fmt.Println("Enter the link you want to shorten: ")

	var redirectUrl string
	fmt.Scanln(&redirectUrl)

	if govalidator.IsURL(redirectUrl) {
		fmt.Println("Enter the path you want: ")
		var shortPath string
		fmt.Scanln(&shortPath)
		var currentUrl string = shortUrl + shortPath
		if govalidator.IsURL(currentUrl) {
			if db.SaveShortPath(shortPath, shortUrl, currentUrl) {

			}
		}
	}
}
