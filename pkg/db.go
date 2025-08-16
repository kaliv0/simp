package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ClipboardItem struct {
	ID        uint // TODO: works by default anyway `gorm:"primaryKey"`
	ClipText  string
	TextHash  string //`gorm:"index"`
	TimeStamp time.Time
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(dbPath string, shouldMigrate bool) *Repository {
	// TODO: refactor constructor
	// create simp.db if not exists
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		// TODO
	}

	if shouldMigrate {
		err = db.AutoMigrate(&ClipboardItem{})
		if err != nil {
			// TODO
		}
	}
	return &Repository{db}
}

func (r *Repository) Write(item []byte) {
	hasher := sha256.New()
	hasher.Write(item)
	// TODO: rename
	textHash := hex.EncodeToString(hasher.Sum(nil))

	var existingItem = ClipboardItem{}
	result := r.db.Where(&ClipboardItem{TextHash: textHash}).First(&existingItem)
	if result.RowsAffected > 0 {
		//if existingItem.ID != 0 {
		existingItem.TimeStamp = time.Now()
		r.db.Save(&existingItem)
	} else {
		r.db.Create(&ClipboardItem{
			ClipText:  string(item),
			TextHash:  textHash,
			TimeStamp: time.Now(),
		})
	}
}

func (r *Repository) Reset() {
	// result := r.db.Exec("Delete from clipboards")
	result := r.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&ClipboardItem{})
	if result.Error != nil {
		// TODO
	}
	//return true
}
