package Curl

import (
	"io/ioutil"
	"net/http"
)

func GetGoodsList(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body
}
