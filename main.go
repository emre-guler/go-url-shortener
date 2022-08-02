package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/emre-guler/url-shortener/db"
	"github.com/emre-guler/url-shortener/validation"
)

const shortUrl string = "https://www.emreguler.dev/"

func main() {
	fmt.Println("Welcome to url-shortener!")
	fmt.Println("---------------------")
	fmt.Println("Enter the link you want to shorten: ")

	readerObject := bufio.NewReader(os.Stdin)
	guestLongUrl, _ := readerObject.ReadString('\n')
	guestLongUrl = strings.Replace(guestLongUrl, "\n", "", -1)
	guestLongUrl = strings.Replace(guestLongUrl, "\r", "", -1)
	if validation.IsValidUrl(guestLongUrl) {
		fmt.Println("Enter the path you want: ")
		guestShortPath, _ := readerObject.ReadString('\n')
		guestShortPath = strings.Replace(guestShortPath, "\n", "", -1)
		guestShortPath = strings.Replace(guestShortPath, "\r", "", -1)
		var currentUrl = shortUrl + guestShortPath
		if validation.IsValidUrl(currentUrl) {
			fmt.Println("Checking db for availability...")
			if db.CheckShortPath(guestShortPath) {
				if db.SaveShortPath(guestShortPath, currentUrl, guestLongUrl) {

				} else {
					fmt.Println("Something went wrong! Try again later.")
				}
			} else {
				fmt.Println("This abbreviation is in use, please try another abbreviation.")
			}
		} else {
			fmt.Println("Enter a valid URL.")
		}
	} else {
		fmt.Println("Enter a valid URL.")
	}
}
