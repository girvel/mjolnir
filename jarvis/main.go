package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	//"github.com/go-audio/audio"
	"github.com/go-audio/wav"

	jarvis "github.com/girvel/mjolnir/jarvis/src"
)

type APICall struct {
    Route *string `json:"route"`
    Args map[string]any `json:"args"`
}

func should(err error) {
    if err != nil {
        slog.Error("initialization failed", "error", err)
        panic("must() failed")
    }
}

func must[T any](result T, err error) T {
	should(err)
    return result
}

func main() {
	// read user request
	var userInput string; {
		audioPath := "/home/nikita/sample.wav"
		f := must(os.Open(audioPath))
		defer f.Close()

		decoder := wav.NewDecoder(f)
		buf := must(decoder.FullPCMBuffer())
		if buf.Format.NumChannels != 1 {
		    panic("Audio must be mono")
		}

		source := buf.AsFloat32Buffer().Data
		floats := make([]float32, len(source))
		for i, s := range source {
		    floats[i] = s / 32768
		}

		modelPath := "/home/nikita/workshop/whisper.cpp/models/ggml-base.en.bin"
		model := must(whisper.New(modelPath))
		defer model.Close()

		context := must(model.NewContext())

		should(context.Process(floats, nil, nil, nil))
		for {
		    segment, err := context.NextSegment()
			if err != nil {
			    break
			}
			fmt.Print(segment.Text)
		}
		fmt.Println()
	}
	panic("");

	// step 1: collect information
    prompt_1 := fmt.Sprintf(
		string(must(os.ReadFile("./prompts/step_1.txt"))),
		string(must(os.ReadFile("../homepage/docs/swagger.json"))),
		string(must(os.ReadFile("../api/docs/swagger.json"))),
		userInput,
	)
	response_1 := jarvis.Prompt(prompt_1)

	var apiCall APICall
	if err := json.Unmarshal([]byte(response_1), &apiCall); err != nil {
	    slog.Error("Error decoding LLM-generated API call", "err", err)
		return
	}

	var info string
	if apiCall.Route != nil {
		resp, err := http.Get("http://" + *apiCall.Route)
		if err != nil {
			log.Fatalf("Error in API call: %v", err)
		}
		defer resp.Body.Close()

		content, _ := io.ReadAll(resp.Body)
		info = string(content)
	} else {
	    info = "(No context needed)"
	}

	// step 2: respond to user
	prompt_2 := fmt.Sprintf(
		string(must(os.ReadFile("./prompts/step_2.txt"))),
		info,
		userInput,
	)

	response_2 := jarvis.Prompt(prompt_2)

	fmt.Println(response_2)
}
