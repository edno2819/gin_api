package models

type Video struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null;unique" json:"title" `
	Description string `json:"description"`
	URL         string `gorm:"not null;unique" json:"url"`
	CreatedAt   int64  `gorm:"autoCreateTime"` // Use unix seconds as creating time

}
