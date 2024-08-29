package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type VoteController struct {
	beego.Controller
}

type Vote struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
	Value   int    `json:"value"`
}

func (v *VoteController) PostVote() {

	// Initialize an instance of the Vote struct
	var vote Vote

	// Parse the JSON request body into the Vote struct
	err := json.Unmarshal(v.Ctx.Input.RequestBody, &vote)
	if err != nil {
		fmt.Println("Error parsing request:", err)
		v.Ctx.Output.SetStatus(400)
		v.Data["json"] = map[string]string{"error": "Invalid request"}
		v.ServeJSON()
		return
	}

	// Get API key from config
	apiKey, err := beego.AppConfig.String("apikey")
	if err != nil {
		fmt.Println("Error getting API key:", err)
		v.Ctx.Output.SetStatus(500)
		v.Data["json"] = map[string]string{"error": "Server configuration error"}
		v.ServeJSON()
		return
	}


	baseUrl, err := beego.AppConfig.String("baseUrl")
	if err != nil {
		fmt.Println("Error getting baseUrl :", err)
		v.Ctx.Output.SetStatus(500)
		v.Data["json"] = map[string]string{"error": "Server configuration error"}
		v.ServeJSON()
		return
	}

	// Prepare channel to receive response
	responseChan := make(chan map[string]interface{})

	// Make the API call in a separate goroutine
	go func() {
		apiURL := fmt.Sprintf("%s/v1/votes", baseUrl)
		jsonData, err := json.Marshal(vote)
		if err != nil {
			fmt.Println("Error marshalling vote data:", err)
			responseChan <- map[string]interface{}{"error": "Internal server error"}
			return
		}

		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error creating HTTP request:", err)
			responseChan <- map[string]interface{}{"error": "Internal server error"}
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", apiKey)

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making HTTP request:", err)
			responseChan <- map[string]interface{}{"error": "Failed to connect to external API"}
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading HTTP response:", err)
			responseChan <- map[string]interface{}{"error": "Failed to read external API response"}
			return
		}
		// Convert byte array to string

		var apiResponse map[string]interface{}
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			fmt.Println("Error parsing JSON response:", err)
			responseChan <- map[string]interface{}{"error": "Failed to parse external API response"}
			return
		}

		responseChan <- apiResponse
	}()

	// Receive the response from the channel
	apiResponse := <-responseChan

	// Send the API response back to the client
	v.Data["json"] = apiResponse
	v.ServeJSON()
}
