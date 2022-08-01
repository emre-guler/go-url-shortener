package validation

import "regexp"

const regExpression string = "(\\b(https?|ftp|file)://)?[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]\n"

func IsValidUrl(url string) bool {
	isValid, _ := regexp.MatchString(regExpression, url)
	return isValid
}
