package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	jarvis "github.com/girvel/mjolnir/jarvis/src"
)

// To interact with Ollama

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
	// read user request
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	// step 1: collect information
    prompt_1 := fmt.Sprintf(
		string(must(os.ReadFile("./prompt.txt"))),
		string(must(os.ReadFile("../homepage/docs/swagger.json"))),
		userInput,
	)
	response_1 := jarvis.Prompt(prompt_1)

	var apiCall APICall
	if err := json.Unmarshal([]byte(response_1), &apiCall); err != nil {
	    slog.Error("Error decoding LLM-generated API call", "err", err)
		return
	}

	resp, err := http.Get("http://" + apiCall.Route)
	if err != nil {
	    log.Fatalf("Error in API call: %v", err)
	}
	defer resp.Body.Close()

	info, _ := io.ReadAll(resp.Body)

	// step 2: respond to user
	prompt_2 := fmt.Sprintf(
		string(must(os.ReadFile("prompt_2.txt"))),
		string(info),
		userInput,
	)

	response_2 := jarvis.Prompt(prompt_2)

	fmt.Println(response_2)
}
