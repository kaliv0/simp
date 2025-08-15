package pkg

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Clipboard struct {
	gorm.Model
	ClipText string
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Write(b []byte) {
	r.db.Create(&Clipboard{ClipText: string(b)})
	// TODO: why not pass just string(bytes)?
}

func NewRepository(dbPath string, shouldMigrate bool) *Repository {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		// TODO
	}

	if shouldMigrate {
		err = db.AutoMigrate(&Clipboard{})
		if err != nil {
			// TODO
		}
	}
	return &Repository{db}
}
