package models

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Key string `gorm:"unique;not null"`
	UserID int
	User User
	Title string `gorm:"type:varchar(200)"`
	Summary string`gorm:"type:varchar(800)"`
	Content string`gorm:"type:text"`
	Visit int `gorm:"default:0"`
	Praise int `gorm:"default:0"`
}

func QueryNoteByKey(key string) (note Note, err error) {
	return note, db.Where("key = ?",key).Take(&note).Error
}

func SaveNote(note *Note) error {
	return db.Save(note).Error
}