package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

// To interact with Ollama
type OllamaRequest struct {
    Model  string `json:"model"`
    Prompt string `json:"prompt"`
    Stream bool   `json:"stream"`
}

type OllamaResponse struct {
    Response string `json:"response"`
}

// The structure your LLM should output
type API_Call struct {
    Action string `json:"action"` // e.g., "turn_on", "turn_off", "set_color"
    Device string `json:"device"`
    Color  string `json:"color,omitempty"`
}

func must[T any](result T, err error) T {
    if err != nil {
        slog.Error("initialization failed", "error", err)
        panic("must() failed")
    }
    return result
}

func main() {
	// Example user input

    // 1. Construct the prompt for the LLM
    prompt := string(must(os.ReadFile("./prompt.txt")))
	apiSpec := string(must(os.ReadFile("../homepage/docs/swagger.json")))
    userInput := "Tell me the address of Grafana"

	prompt = fmt.Sprintf(prompt, apiSpec, userInput)

	fmt.Println("PROMPT:", prompt);

    // 2. Send the request to Ollama
    ollamaReq := OllamaRequest{
        Model:  "llama3",
        Prompt: prompt,
        Stream: false,
    }
    reqBody, _ := json.Marshal(ollamaReq)

    resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(reqBody))
    if err != nil {
        log.Fatalf("Error calling Ollama: %v", err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    var ollamaResp OllamaResponse
    if err := json.Unmarshal(body, &ollamaResp); err != nil {
        log.Fatalf("Error decoding Ollama response: %v", err)
    }

    // Clean up the LLM's output to get just the JSON
    llmOutput := strings.TrimSpace(ollamaResp.Response)
    fmt.Printf("LLM Output: %s\n", llmOutput)
}
