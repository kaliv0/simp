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

func (r *Repository) Write(item []byte) {
	// todo:why not pass just string(bytes)?
	r.db.Create(&Clipboard{ClipText: string(item)})

	// TODO: use upsert -> if item already in db -> change update_timestamp
}

func NewRepository(dbPath string, shouldMigrate bool) *Repository {
	// create simp.db if not exists
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
