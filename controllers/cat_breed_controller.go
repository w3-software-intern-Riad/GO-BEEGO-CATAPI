package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type GetBreedsController struct {
	beego.Controller
}

func (gb *GetBreedsController) GetAllBreeds() {
	// Define channels to handle responses and errors
	responseChan := make(chan []map[string]interface{})
	errorChan := make(chan error)

	// Close channels when done
	defer close(responseChan)
	defer close(errorChan)

	// Fetch baseUrl from configuration
	baseUrl, err := beego.AppConfig.String("baseUrl")
	if err != nil {
		gb.handleError("Server configuration error", err)
		return
	}

	go func() {
		defer close(responseChan)
		defer close(errorChan)

		apiURL := fmt.Sprintf("%s/v1/breeds", baseUrl)
		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			errorChan <- fmt.Errorf("failed to create request: %w", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- fmt.Errorf("failed to fetch data: %w", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("failed to read response body: %w", err)
			return
		}

		var breeds []map[string]interface{}
		if err := json.Unmarshal(body, &breeds); err != nil {
			errorChan <- fmt.Errorf("failed to parse JSON response: %w", err)
			return
		}

		// Process and send response
		var processedBreeds []map[string]interface{}
		for _, breed := range breeds {
			processedBreed := map[string]interface{}{
				"id":             breed["id"],
				"name":           breed["name"],
				"description":    breed["description"],
				"wikipedia_url":  breed["wikipedia_url"],
				"origin":         breed["origin"],
			}
			processedBreeds = append(processedBreeds, processedBreed)
		}

		responseChan <- processedBreeds
	}()

	// Handle response or error
	select {
	case err := <-errorChan:
		gb.handleError(err.Error(), err)
	case apiResponse := <-responseChan:
		gb.Data["json"] = apiResponse
	}

	gb.ServeJSON()
}

// handleError is a helper function to set error responses
func (gb *GetBreedsController) handleError(message string, err error) {
	fmt.Println("Error:", err)
	gb.Ctx.Output.SetStatus(500)
	gb.Data["json"] = map[string]interface{}{"error": message}
	gb.ServeJSON()
}
