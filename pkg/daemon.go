package pkg

import (
	"context"
	"os"

	"github.com/shirou/gopsutil/process"
	"golang.design/x/clipboard"
)

// TODO: move to handler.go?

func TrackClipboard(dbPath string) {
	//create/get new gorm repo handler, migrate if any changes in clipboard text
	db := NewRepository(dbPath, true)
	//  init clipboard handler,
	err := clipboard.Init()
	if err != nil {
		// TODO
	}
	// open new Watch chanel -> NB: tracks text but no images
	changes := clipboard.Watch(context.Background(), clipboard.FmtText)
	// loop through && write to db any changes in clipboard
	for item := range changes {
		db.Write(item)
	}
}

func StopAllInstances() {
	processes, err := process.Processes()
	if err != nil {
		// TODO
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			// TODO
		}
		if n == "simp" && int32(os.Getpid()) != p.Pid {
			err := p.Kill()
			if err != nil {
				//TODO
			}
		}
	}
}
