package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
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

func main() {
	// Example user input
    userInput := "Turn the living room light on and set it to red"

    // 1. Construct the prompt for the LLM
    prompt := fmt.Sprintf(`
        You are an AI assistant for a smart home. Your task is to translate human language into a JSON object representing an API call.
        The available actions are "turn_on", "turn_off", and "set_color".
        The only available device is "living_room_light".
        Based on the user's request: "%s", create a JSON object with the action, device, and color if specified.
        Only output the JSON object, nothing else.
    `, userInput)

    // 2. Send the request to Ollama
    ollamaReq := OllamaRequest{
        Model:  "tinydolphin",
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
