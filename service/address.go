package service

import (
	"domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func FindAddress(latlong domain.AddressRequest) (interface{}, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.vworld.kr/req/address", nil)
	if err != nil {
		log.Print(err)
		return "", err
	}

	longitudeStr := fmt.Sprintf("%f", latlong.Longitude)
	latitudeStr := fmt.Sprintf("%f", latlong.Latitude)

	log.Println(latlong, longitudeStr, latitudeStr)

	q := req.URL.Query()
	q.Add("service", "address")
	q.Add("request", "getAddress")
	q.Add("key", os.Getenv("ADDRESS_KEY"))
	q.Add("point", longitudeStr+","+latitudeStr)
	q.Add("type", "PARCEL")
	q.Add("crs", "epsg:4326")

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	var res interface{}
	json.Unmarshal(resp_body, &res)

	return res, nil
}
