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
type APICall struct {
    Route string `json:"route"`
    Args map[string]any `json:"args"`
}

func must[T any](result T, err error) T {
    if err != nil {
        slog.Error("initialization failed", "error", err)
        panic("must() failed")
    }
    return result
}

func main() {
    // 1. Construct the prompt for the LLM
    prompt := string(must(os.ReadFile("./prompt.txt")))
	apiSpec := string(must(os.ReadFile("../homepage/docs/swagger.json")))
    userInput := "What is thor1?"

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
		return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    var ollamaResp OllamaResponse
    if err := json.Unmarshal(body, &ollamaResp); err != nil {
        log.Fatalf("Error decoding Ollama response: %v", err)
		return
    }

    // Clean up the LLM's output to get just the JSON
    llmOutput := strings.TrimSpace(ollamaResp.Response)
    fmt.Printf("LLM Output: `%s`\n", llmOutput)

	var apiCall APICall
	if err := json.Unmarshal([]byte(llmOutput), &apiCall); err != nil {
	    log.Fatalf("Error decoding LLM-generated API call: %v", err)
		return
	}

	resp, err = http.Get("http://" + apiCall.Route)
	if err != nil {
	    log.Fatalf("Error in API call: %v", err)
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)

	prompt_2 := fmt.Sprintf(
		string(must(os.ReadFile("prompt_2.txt"))),
		string(body),
		userInput,
	)

	fmt.Println("PROMPT #2:", prompt_2);

	ollamaReq = OllamaRequest{
	    Model: "llama3",
		Prompt: prompt_2,
		Stream: false,
	}
    reqBody, _ = json.Marshal(ollamaReq)

    resp, err = http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(reqBody))
    if err != nil {
        log.Fatalf("Error calling Ollama: %v", err)
		return
    }
    defer resp.Body.Close()

    body, _ = io.ReadAll(resp.Body)
    if err := json.Unmarshal(body, &ollamaResp); err != nil {
        log.Fatalf("Error decoding Ollama response: %v", err)
		return
    }

    // Clean up the LLM's output to get just the JSON
    llmOutput = strings.TrimSpace(ollamaResp.Response)
    fmt.Printf("LLM Output 2: `%s`\n", llmOutput)
}
