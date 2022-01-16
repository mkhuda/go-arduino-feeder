package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/mkhuda/go-arduino-feeder/configs"
	"github.com/mkhuda/go-arduino-feeder/models"
)

func Api_Get() {
	// Create a resty client
	baseUrl := configs.GetConfig().ApiHost

	client := resty.New()

	resp, err := client.R().Get(baseUrl)

	fmt.Printf("\nError: %v %v", os.Getenv("API_HOST"), err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
}

func Api_Post_Temperature(temperature int, humidity int, heat int) {
	baseUrl := configs.GetConfig().ApiHost

	client := resty.New()

	payload := models.Temperature{
		DeviceSerial: "AABBBCC112233",
		Temperature:  int(temperature),
		Humidity:     int(humidity),
		Heat:         int(heat),
		CreatedOn:    time.Now().Format(time.RFC3339),
	}

	//Convert User to byte using Json.Marshal
	//Ignoring error.
	// body, _ := json.Marshal(payload)
	resp1, err1 := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(baseUrl + "/temperature")
	fmt.Println(baseUrl+"/temperature", resp1, err1)
}
