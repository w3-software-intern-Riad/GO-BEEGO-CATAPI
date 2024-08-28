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
type GetBreedController struct {
	beego.Controller
}

func (gb *GetBreedsController) GetAllBreeds() {
	// Define a channel to handle responses
	responseChan := make(chan []map[string]interface{})
	errorChan := make(chan error)

	go func() {
		apiURL := "https://api.thecatapi.com/v1/breeds"

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			errorChan <- fmt.Errorf("failed to create request")
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- fmt.Errorf("failed to fetch data")
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("failed to read response body")
			return
		}

		var breeds []map[string]interface{}
		if err := json.Unmarshal(body, &breeds); err != nil {
			errorChan <- fmt.Errorf("failed to parse JSON response")
			return
		}

		// Extract required fields
		var processedBreeds []map[string]interface{}
		for _, breed := range breeds {
			processedBreed := map[string]interface{}{"id": breed["id"], "name": breed["name"], "description": breed["description"], "wikipedia_url": breed["wikipedia_url"], "origin": breed["origin"]}
			processedBreeds = append(processedBreeds, processedBreed)
		}

		responseChan <- processedBreeds
	}()

	// Check for errors or receive the response
	select {
	case err := <-errorChan:
		fmt.Println("Error:", err)
		gb.Data["json"] = map[string]interface{}{"error": err.Error()}
	case apiResponse := <-responseChan:
		gb.Data["json"] = apiResponse
	}

	gb.ServeJSON()
}
