package jarvis

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type ollamaRequest struct {
    Model  string `json:"model"`
    Prompt string `json:"prompt"`
    Stream bool   `json:"stream"`
}

type ollamaResponse struct {
    Response string `json:"response"`
}

func Prompt(prompt string) string {
	slog.Info("LLM prompted", "prompt", prompt)

    ollamaReq := ollamaRequest{
        Model:  "llama3",
        Prompt: prompt,
        Stream: false,
    }
    reqBody, _ := json.Marshal(ollamaReq)

    resp, err := http.Post(
		"http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(reqBody),
	)

    if err != nil {
        slog.Error("Error calling Ollama", "err", err)
		panic(err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    var ollamaResp ollamaResponse
    if err := json.Unmarshal(body, &ollamaResp); err != nil {
        slog.Error("Error decoding Ollama response", "err", err)
		panic(err)
    }

    llmOutput := strings.TrimSpace(ollamaResp.Response)
	slog.Info("LLM responded", "output", llmOutput)
	return llmOutput
}
