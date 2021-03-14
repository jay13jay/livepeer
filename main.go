package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	// Set up the payload to create the streams
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

	// Print some debug information to the terminal
	fmt.Println("Printing debug information...")

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, payloadBytes, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}

	log.Printf("Payload:\n%s\n\n", string(prettyJSON.Bytes()))
	// fmt.Printf("Payload:\n%s\n\n\n", payloadBytes)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	fmt.Printf("HTTP Response:\t%d\n", int(resp.StatusCode))
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bytes))
}

func getEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
