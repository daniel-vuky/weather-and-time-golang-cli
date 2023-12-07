package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	URI    string `json:"uri"`
	APIKey string `json:"api_key"`
}

var settings Settings

func AddSettings(newUri, newApiKey string) {
	newSettings := &Settings{
		URI:    newUri,
		APIKey: newApiKey,
	}

	// Marshal the struct directly to JSON
	newSettingsAsJson, newSettingsAsJsonErr := json.Marshal(newSettings)
	if newSettingsAsJsonErr != nil {
		panic(newSettingsAsJsonErr)
	}
	writer, writerErr := os.Create(getSettingsFileName())
	if writerErr != nil {
		panic(writerErr)
	}

	_, writeError := writer.Write(newSettingsAsJson)
	if writeError != nil {
		panic(writeError)
	}

	fmt.Println("Write new settings success.")
}

func GetUri() string {
	readSettingsFile()
	return settings.URI
}

func GetApiKey() string {
	readSettingsFile()
	return settings.APIKey
}

func readSettingsFile() Settings {
	if settings.APIKey == "" || settings.URI == "" {
		fileContent, err := os.ReadFile(getSettingsFileName())
		if err != nil {
			panic("Can not reading settings file.")
		}

		// Unmarshal the JSON data into the struct
		err = json.Unmarshal(fileContent, &settings)
		if err != nil {
			panic(fmt.Sprintf("Error unmarshalling JSON: %s", err))
		}
	}

	return settings
}

func getSettingsFileName() string {
	currentFolder, err := os.Getwd()
	if err != nil {
		panic("Can not find the config folder.")
	}
	return fmt.Sprintf("%s/%s", currentFolder, "settings.json")
}
