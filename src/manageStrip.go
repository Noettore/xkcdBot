package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type xkcdStrip struct {
	ID         int    `json:"num"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Alternate  string `json:"alt"`
	Transcript string `json:"transcript"`
	ImgLink    string `json:"img"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Link       string `json:"link"`
	News       string `json:"news"`
}

const (
	xkcdURL           string = "https://xkcd.com/"
	stripJSONFileName string = "info.0.json"
	lastStripJSONURL  string = "https://xkcd.com/info.0.json"
)

func stripScraper(stripURL string) (xkcdStrip, error) {
	var strip xkcdStrip

	res, err := http.Get(stripURL)
	if err != nil {
		log.Printf("Error getting last strip json: %v", err)
		return strip, errors.Wrap(err, "getting last strip failed")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading page body: %v", err)
		return strip, errors.Wrap(err, "reading page body failed")
	}

	err = json.Unmarshal(body, &strip)
	if err != nil {
		log.Printf("Error unmarshalling last strip json: %v", err)
		return strip, errors.Wrap(err, "unmarshalling last strip failed")
	}
	return strip, nil
}

func getStripInfoFromDB(stripID int) (xkcdStrip, error) {

}

func getLastStripInfoFromDB(stripID int) (xkcdStrip, error) {

}

func getStripInfoFromWeb(stripID int) (xkcdStrip, error) {
	return stripScraper(xkcdURL + strconv.Itoa(stripID) + "/" + stripJSONFileName)
}

func getLastStripInfoFromWeb() (xkcdStrip, error) {
	return stripScraper(lastStripJSONURL)
}

func updateDB() {
	strip, err := getLastStripInfoFromWeb()
	if err != nil {
		log.Printf("Error getting last strip info from web: %v", err)
		return
	}

}
