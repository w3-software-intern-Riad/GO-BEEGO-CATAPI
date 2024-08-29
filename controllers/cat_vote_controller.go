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

// PostVote handles the voting process
func (v *VoteController) PostVote() {
	var vote Vote

	// Parse JSON request body
	if err := json.Unmarshal(v.Ctx.Input.RequestBody, &vote); err != nil {
		v.handleError("Invalid request", err, 400)
		return
	}

	apiKey, err := beego.AppConfig.String("apikey")
	if err != nil {
		v.handleError("Server configuration error", err, 500)
		return
	}

	baseUrl, err := beego.AppConfig.String("baseUrl")
	if err != nil {
		v.handleError("Server configuration error", err, 500)
		return
	}

	responseChan := make(chan map[string]interface{})
	defer close(responseChan)

	// Make API call in a goroutine
	go func() {
		defer close(responseChan)
		
		apiURL := fmt.Sprintf("%s/v1/votes", baseUrl)
		jsonData, err := json.Marshal(vote)
		if err != nil {
			responseChan <- map[string]interface{}{"error": "Internal server error"}
			return
		}

		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
		if err != nil {
			responseChan <- map[string]interface{}{"error": "Internal server error"}
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", apiKey)

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			responseChan <- map[string]interface{}{"error": "Failed to connect to external API"}
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			responseChan <- map[string]interface{}{"error": "Failed to read external API response"}
			return
		}

		var apiResponse map[string]interface{}
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			responseChan <- map[string]interface{}{"error": "Failed to parse external API response"}
			return
		}

		responseChan <- apiResponse
	}()

	// Handle API response
	apiResponse := <-responseChan
	if errorResp, ok := apiResponse["error"].(string); ok {
		v.handleError(errorResp, nil, 500)
		return
	}

	v.Data["json"] = apiResponse
	v.ServeJSON()
}

// handleError is a helper function to send error responses
func (v *VoteController) handleError(message string, err error, statusCode int) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	v.Ctx.Output.SetStatus(statusCode)
	v.Data["json"] = map[string]string{"error": message}
	v.ServeJSON()
}
