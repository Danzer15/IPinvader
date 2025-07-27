// API Untuk Lokasi
package apiLoc

import (
	"encoding/json"
	"net/http"
)

type Location struct {
	Scountry string `json:"country"`
	Scity    string `json:"city"`
	Sisp     string `json:"isp"`
}

func GetLocation(ip string) Location {
	url := "http://ip-api.com/json/" + ip
	resp, err := http.Get(url)
	if err != nil {
		return Location{}
	}
	defer resp.Body.Close()

	var loc Location
	err = json.NewDecoder(resp.Body).Decode(&loc)
	if err != nil {
		return Location{}
	}

	return loc
}
