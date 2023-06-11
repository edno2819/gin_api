package models

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null;unique"`
	Password  string `gorm:"not null;unique"`
	UpdatedAt int64  `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	CreatedAt int64  `gorm:"autoCreateTime"`      // Use unix seconds as creating time
}
