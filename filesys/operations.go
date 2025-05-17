package filesys

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/aniketkharel/rssreader/models"
)

/*RWX PERM*/
const DEFAULT_PERMISSION int = 0755
const DEFAULT_CONFIG string = `{"urls":[]}`

func Read() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	setup_files_folders(home)
	if err != nil {
		log.Fatal(err)
	}
}

func setup_files_folders(HOME string) {
	path := filepath.Join(HOME, ".config", "lazyrss")
	error := os.MkdirAll(path, os.FileMode(DEFAULT_PERMISSION))
	if error != nil {
		log.Fatal(error)
	}

	filePath := filepath.Join(path, "config.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		werr := os.WriteFile(filePath, []byte(DEFAULT_CONFIG), os.FileMode(DEFAULT_PERMISSION))
		if werr != nil {
			log.Fatal(werr)
		}
	}
}

func Load_config() *models.Config {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	file_path := filepath.Join(home, ".config", "lazyrss", "config.json")
	data, rerr := os.ReadFile(file_path)
	if rerr != nil {
		log.Fatal(rerr)
	}
	var config models.Config
	uerr := json.Unmarshal(data, &config)
	if uerr != nil {
		log.Fatal(uerr)
	}
	return &config
}
