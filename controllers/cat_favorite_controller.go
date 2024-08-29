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
	if err := json.Unmarshal(f.Ctx.Input.RequestBody, &favorite); err != nil {
		f.handleError("Invalid request", err, 400)
		return
	}

	apiKey, err := beego.AppConfig.String("apikey")
	if err != nil {
		f.handleError("Server configuration error", err, 500)
		return
	}

	baseUrl, err := beego.AppConfig.String("baseUrl")
	if err != nil {
		f.handleError("Server configuration error", err, 500)
		return
	}

	responseChan := make(chan map[string]interface{})
	errorChan := make(chan error)

	go func() {
		defer close(responseChan)
		defer close(errorChan)

		apiURL := fmt.Sprintf("%s/v1/favourites", baseUrl)
		jsonData, err := json.Marshal(favorite)
		if err != nil {
			errorChan <- fmt.Errorf("error marshalling favorite data: %w", err)
			return
		}

		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
		if err != nil {
			errorChan <- fmt.Errorf("error creating HTTP request: %w", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", apiKey)

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- fmt.Errorf("error making HTTP request: %w", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("error reading HTTP response: %w", err)
			return
		}

		var apiResponse map[string]interface{}
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			errorChan <- fmt.Errorf("error parsing JSON response: %w", err)
			return
		}

		responseChan <- apiResponse
	}()

	select {
	case err := <-errorChan:
		f.handleError("Internal server error", err, 500)
	case apiResponse := <-responseChan:
		f.Data["json"] = apiResponse
		f.ServeJSON()
	}
}

func (gf *GetFavController) GetAllFav() {
	responseChan := make(chan []ImageResponse)
	errorChan := make(chan error)

	baseUrl, err := beego.AppConfig.String("baseUrl")
	if err != nil {
		gf.handleError("Server configuration error", err, 500)
		return
	}

	apiKey, err := beego.AppConfig.String("apikey")
	if err != nil {
		gf.handleError("Server configuration error", err, 500)
		return
	}

	go func() {
		defer close(responseChan)
		defer close(errorChan)

		apiURL := fmt.Sprintf("%s/v1/favourites", baseUrl)

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			errorChan <- fmt.Errorf("error creating HTTP request: %w", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- fmt.Errorf("error making HTTP request: %w", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("error reading HTTP response: %w", err)
			return
		}

		var apiResponse []ImageResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			errorChan <- fmt.Errorf("error parsing JSON response: %w", err)
			return
		}

		responseChan <- apiResponse
	}()

	select {
	case err := <-errorChan:
		gf.handleError("Internal server error", err, 500)
	case apiResponse := <-responseChan:
		gf.Data["json"] = apiResponse
		gf.ServeJSON()
	}
}

// handleError is a helper function to set error responses
func (f *FavoriteController) handleError(message string, err error, statusCode int) {
	fmt.Println("Error:", err)
	f.Ctx.Output.SetStatus(statusCode)
	f.Data["json"] = map[string]string{"error": message}
	f.ServeJSON()
}

// handleError is a helper function to set error responses for GetFavController
func (gf *GetFavController) handleError(message string, err error, statusCode int) {
	fmt.Println("Error:", err)
	gf.Ctx.Output.SetStatus(statusCode)
	gf.Data["json"] = map[string]string{"error": message}
	gf.ServeJSON()
}
