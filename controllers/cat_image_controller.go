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
type CatImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (c *CatController) GetCatImage() {
	resultChan := make(chan []CatImage)
	errorChan := make(chan error)
	defer close(resultChan)
	defer close(errorChan)

	baseUrl, err := beego.AppConfig.String("baseUrl")
	if err != nil {
		c.handleError("Server configuration error", err, 500)
		return
	}

	go func() {
		apiURL := fmt.Sprintf("%s/v1/images/search?size=med&mime_types=jpg&format=json&has_breeds=false&order=RANDOM&page=0&limit=1", baseUrl)
		apiKey, err := beego.AppConfig.String("apikey")
		if err != nil {
			errorChan <- fmt.Errorf("failed to load API key: %w", err)
			return
		}

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			errorChan <- fmt.Errorf("failed to create request: %w", err)
			return
		}
		req.Header.Set("x-api-key", apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- fmt.Errorf("failed to fetch cat image: %w", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("failed to read response body: %w", err)
			return
		}

		var catImages []CatImage
		if err := json.Unmarshal(body, &catImages); err != nil {
			errorChan <- fmt.Errorf("failed to parse response: %w", err)
			return
		}

		if len(catImages) == 0 {
			errorChan <- fmt.Errorf("no cat image found")
			return
		}

		resultChan <- catImages
	}()

	select {
	case catImages := <-resultChan:
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.Ctx.Output.Header("Content-Type", "application/json")
		c.Data["json"] = catImages[0]
	case err := <-errorChan:
		c.handleError(err.Error(), err, http.StatusInternalServerError)
	}
	c.ServeJSON()
}

func (gci *GetCatImagesController) GetCatImages() {
	responseChan := make(chan []string)
	errorChan := make(chan error)
	defer close(responseChan)
	defer close(errorChan)

	breedID := gci.Ctx.Input.Param(":breed_id")
	if breedID == "" {
		gci.Data["json"] = map[string]string{"error": "breed_id is required"}
		gci.ServeJSON()
		return
	}

	baseUrl, err := beego.AppConfig.String("baseUrl")
	if err != nil {
		gci.handleError("Server configuration error", err, 500)
		return
	}

	go func() {
		apiURL := fmt.Sprintf("%s/v1/images/search?limit=5&breed_ids=%s", baseUrl, breedID)

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

		var urls []string
		for _, img := range images {
			if url, ok := img["url"].(string); ok {
				urls = append(urls, url)
			}
		}

		responseChan <- urls
	}()

	select {
	case err := <-errorChan:
		gci.handleError(err.Error(), err, http.StatusInternalServerError)
	case urls := <-responseChan:
		gci.Data["json"] = urls
	}

	gci.ServeJSON()
}

func (c *CatController) handleError(message string, err error, statusCode int) {
	if err != nil {
		fmt.Println("Error:", err)  // Use your preferred logging mechanism
	}
	c.Ctx.Output.SetStatus(statusCode)
	c.Data["json"] = map[string]string{"error": message}
	c.ServeJSON()
}





func (gci *GetCatImagesController) handleError(message string, err error, statusCode int) {
	if err != nil {
		fmt.Println("Error:", err)  // Use your preferred logging mechanism
	}
	gci.Ctx.Output.SetStatus(statusCode)
	gci.Data["json"] = map[string]string{"error": message}
	gci.ServeJSON()
}