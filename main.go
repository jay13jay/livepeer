package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Payload struct {
	Name     string    `json:"name"`
	Profiles []Profile `json:"profiles"`
}
type Profile struct {
	Name    string `json:"name"`
	Bitrate int    `json:"bitrate"`
	Fps     int    `json:"fps"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

func main() {
	getEnv()
	bearer := string("Bearer " + os.Getenv("APIKEY"))
	apiEndpoint := string(os.Getenv("APIENDPOINT"))
	fmt.Printf("apiEndpoiont:\t%s\n", apiEndpoint)
	data := Payload{
		Name: "test_stream",
		Profiles: []Profile{
			{
				Name:    "720p",
				Bitrate: 2000000,
				Fps:     30,
				Width:   1280,
				Height:  720,
			},
			{
				Name:    "480p",
				Bitrate: 1000000,
				Fps:     30,
				Width:   854,
				Height:  480,
			},
			{
				Name:    "360p",
				Bitrate: 500000,
				Fps:     30,
				Width:   640,
				Height:  360,
			},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", apiEndpoint, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)
	fmt.Println("Printing debug information...")
	fmt.Printf("Headers:\n%s\n", req.Header)
	fmt.Printf("JSON:\n%s\n\n", data)
	fmt.Printf("Payload:\n%s\n\n\n", payloadBytes)

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	// handle err
	// }
	// defer resp.Body.Close()
}

func getEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
