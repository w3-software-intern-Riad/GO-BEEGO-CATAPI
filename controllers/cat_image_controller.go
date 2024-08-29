package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type CatController struct {
	beego.Controller
}

type GetCatImagesController struct {
	beego.Controller
}

// CatImage represents the structure of the response from the Cat API
type CatImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// GetCatImage is the handler function for the API endpoint using Go channels
func (c *CatController) GetCatImage() {
	// Channel to receive the result of the API call
	resultChan := make(chan []CatImage)
	errorChan := make(chan error)
	baseUrl, err := beego.AppConfig.String("baseUrl")
	if err != nil {
		fmt.Println("Error getting baseUrl :", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Server configuration error"}
		c.ServeJSON()
		return
	}
	// Start a goroutine to fetch the cat image
	go func() {
		apiURL := fmt.Sprintf("%s/v1/images/search?size=med&mime_types=jpg&format=json&has_breeds=false&order=RANDOM&page=0&limit=1",baseUrl)
		apiKey, err := beego.AppConfig.String("apikey")
		if err != nil {
			errorChan <- fmt.Errorf("failed to load API key")
			return
		}

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			errorChan <- fmt.Errorf("failed to create request")
			return
		}

		req.Header.Set("x-api-key", apiKey)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- fmt.Errorf("failed to fetch cat image")
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("failed to read response body")
			return
		}

		var catImages []CatImage
		if err := json.Unmarshal(body, &catImages); err != nil {
			errorChan <- fmt.Errorf("failed to parse response")
			return
		}

		if len(catImages) == 0 {
			errorChan <- fmt.Errorf("no cat image found")
			return
		}

		// Send the result back through the channel
		resultChan <- catImages
	}()

	// Wait for the result or error from the channels
	select {
	case catImages := <-resultChan:
		// Successfully received the cat image
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.Ctx.Output.Header("Content-Type", "application/json")
		c.Data["json"] = catImages[0]
	case err := <-errorChan:
		// An error occurred
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
	}

	// Send the JSON response
	c.ServeJSON()
}

// GetCatImages fetches cat image URLs for a specific breed using a channel.
func (gci *GetCatImagesController) GetCatImages() {
	// Define channels for responses and errors
	responseChan := make(chan []string)
	errorChan := make(chan error)

	// Get breed_id from query parameters
	breedID := gci.Ctx.Input.Param(":breed_id")

	fmt.Println("breed_id : ",breedID)
	if breedID == "" {
		gci.Data["json"] = map[string]string{"error": "breed_id is required"}
		gci.ServeJSON()
		return
	}
	baseUrl, err := beego.AppConfig.String("baseUrl")
	if err != nil {
		fmt.Println("Error getting baseUrl :", err)
		gci.Ctx.Output.SetStatus(500)
		gci.Data["json"] = map[string]string{"error": "Server configuration error"}
		gci.ServeJSON()
		return
	}
	// Fetch cat image URLs in a goroutine
	go func() {
		apiURL := fmt.Sprintf("%s/v1/images/search?limit=5&breed_ids=%s",baseUrl, breedID)

		resp, err := http.Get(apiURL)
		if err != nil {
			errorChan <- fmt.Errorf("failed to fetch cat images: %w", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("failed to read response body: %w", err)
			return
		}

		var images []map[string]interface{}
		if err := json.Unmarshal(body, &images); err != nil {
			errorChan <- fmt.Errorf("failed to parse JSON response: %w", err)
			return
		}

		// Extract image URLs
		var urls []string
		for _, img := range images {
			if url, ok := img["url"].(string); ok {
				urls = append(urls, url)
			}
		}

		responseChan <- urls
	}()

	// Select to wait for either the response or an error
	select {
	case err := <-errorChan:
		gci.Data["json"] = map[string]string{"error": err.Error()}
	case urls := <-responseChan:
		gci.Data["json"] = urls
	}

	gci.ServeJSON()
}
