package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

func logic() error {
    bytes, err := os.ReadFile("sample.toml")
	if err != nil {
	    return err
	}

    var schedule map[string]map[string]string
	err = toml.Unmarshal(bytes, &schedule)
	if err != nil {
	    return err
	}

	fmt.Println(schedule)
	return nil
}

func main() {
	err := logic()
	if err != nil {
	    panic(err)
	}
}
