package pkg

import (
	"os"
	"path/filepath"
)

// TODO: move to db.go

func GetDbPath() string {
	var subDirsList []string
	xdfConfig := os.Getenv("XDG_CONFIG_HOME")
	if xdfConfig != "" {
		subDirsList = append(subDirsList, xdfConfig)
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			// TODO
		}
		subDirsList = append(subDirsList, homeDir, ".config")
	}
	subDirsList = append(subDirsList, "simp")
	configDir := filepath.Join(subDirsList...)

	err := os.MkdirAll(configDir, 0744) // TODO: modify perm, skip calling mkdir if exists?
	if err != nil {
		// TODO
	}
	return filepath.Join(configDir, "simp.db")
}
