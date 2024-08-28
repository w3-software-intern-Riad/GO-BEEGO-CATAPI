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

type FavoriteController struct {
	beego.Controller
}

type GetFavController struct {
	beego.Controller
}

type Favorite struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
}

type ImageResponse struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	ImageID   string `json:"image_id"`
	SubID     string `json:"sub_id"`
	CreatedAt string `json:"created_at"`
	Image     struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"image"`
}

func (f *FavoriteController) PostFavorite() {
	var favorite Favorite
	err := json.Unmarshal(f.Ctx.Input.RequestBody, &favorite)
	if err != nil {
		fmt.Println("Error parsing request:", err)
		f.Ctx.Output.SetStatus(400)
		f.Data["json"] = map[string]string{"error": "Invalid request"}
		f.ServeJSON()
		return
	}
	// Get API key from config
	apiKey, err := beego.AppConfig.String("apikey")
	if err != nil {
		fmt.Println("Error getting API key:", err)
		f.Ctx.Output.SetStatus(500)
		f.Data["json"] = map[string]string{"error": "Server configuration error"}
		f.ServeJSON()
		return
	}
	responseChan := make(chan map[string]interface{})
	go func() {
		apiURL := "https://api.thecatapi.com/v1/favourites"
		jsonData, err := json.Marshal(favorite)
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
	f.Data["json"] = apiResponse
	f.ServeJSON()
}

func (gf *GetFavController) GetAllFav() {
	responseChan := make(chan []ImageResponse)
	errorChan := make(chan error)
	go func() {
		apiURL := "https://api.thecatapi.com/v1/favourites"
		apiKey, err := beego.AppConfig.String("apikey")
		if err != nil {
			errorChan <- fmt.Errorf("failed to load API key: %v", err)
			return
		}

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			errorChan <- fmt.Errorf("failed to create request: %v", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", apiKey)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- fmt.Errorf("failed to fetch cat images: %v", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("failed to read response body: %v", err)
			return
		}

		var apiResponse []ImageResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			errorChan <- fmt.Errorf("failed to parse response body: %v", err)
			return
		}

		responseChan <- apiResponse
	}()

	select {
	case apiResponse := <-responseChan:
		gf.Data["json"] = apiResponse
		gf.ServeJSON()
	case err := <-errorChan:
		gf.Data["json"] = map[string]string{"error": err.Error()}
		gf.ServeJSON()
	}
}
