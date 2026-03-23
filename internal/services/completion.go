package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"prompt-control-go/internal/models/prompt"
)

const (
	GenerateUrl = "http://localhost:8081/api/prompts/generate"
	RefineUrl = "http://localhost:8081/api/prompts/refine"
	EnrichUrl = "http://localhost:8081/api/prompts/enrich"
)

type GeneratePromptRequest struct {
	Prompt string `json:"prompt"`
}

type RefinePromptRequest struct {
	UserQuery string `json:"user_query"`
}

func SendGenerateNewPrompt(query string) (string, error) {
	body, _ := json.Marshal(GeneratePromptRequest{query})

	resp, err := http.Post(GenerateUrl, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
  	if err != nil {
  		return "", err
  	}

  	if resp.StatusCode >= 400 {
  		return "", fmt.Errorf("prompt service returned %d: %s", resp.StatusCode, string(data))
  	}

  	return string(data), nil
}

func SendRefinePrompt(query string) (string) {
	resp := sendRequest(RefinePromptRequest{query}, RefineUrl)

	defer resp.Body.Close()

	var result map[string]interface{}
	json.Unmarshal([]byte(readData(resp)), &result)
	compiledPrompt := result["compiled_prompt"].(string)
	return compiledPrompt
}

func SendEnrichPrompt(rq models.EnrichRequest) models.EnrichResponse {
	resp := sendRequest(rq, EnrichUrl)

	defer resp.Body.Close()

	var result models.EnrichResponse
	json.Unmarshal([]byte(readData(resp)), &result)

	return result
}

func sendRequest(request any, url string) *http.Response {
	body, _ := json.Marshal(request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Printf("failed send request: \n%s", request)
		return nil
	}

	log.Println("response: ", resp)

	return resp
}

func readData(resp *http.Response) string {
	data, err := io.ReadAll(resp.Body)
  	if err != nil {
  		return ""
  	}

	return string(data)
}