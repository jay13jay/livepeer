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
	Name     string     `json:"name"`
	Profiles []Profiles `json:"profiles"`
}
type Profiles struct {
	Name    string `json:"name"`
	Bitrate int    `json:"bitrate"`
	Fps     int    `json:"fps"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

func main() {
	getEnv()
	bearer := "Bearer " + os.Getenv("APIKEY")
	data := Payload{
		// fill struct
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)
	fmt.Printf("Headers:\n%s\n\n", req.Header)
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
