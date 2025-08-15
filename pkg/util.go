package pkg

import (
	"os"
	"path/filepath"
)

// TODO: move to db.go

func GetDbPath() string {
	// resolve config dir
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

	// find/create .config/simp
	err := os.MkdirAll(configDir, 0744) // TODO: modify perm
	if err != nil {
		// TODO
	}
	// create simp.db if not exists
	dbPath := filepath.Join(configDir, "simp.db")
	_, err = os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		// TODO
	}
	return dbPath
}
