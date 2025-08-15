package pkg

import (
	"context"
	"fmt"
	"os"

	"github.com/shirou/gopsutil/process"
	"golang.design/x/clipboard"
)

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
	for b := range changes {
		db.Write(b)
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
			pid := os.Getpid()
			fmt.Printf("%d\n", pid)
			fmt.Printf("%d\n", p.Pid)

			err := p.Terminate()
			if err != nil {
				//TODO
			}
		}
	}
	//return fmt.Errorf("process not found")
}
