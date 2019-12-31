package model

import (
	"testgame/database"
	"time"
)

func init() {
	database.Db.AutoMigrate(
		&User{},
		&Log{},
	)
}

type User struct {
	ID            uint `gorm:"primary_key"json:"-"`
	UserCode      string
	GameOneStatus int        // 1：测试 2：正式1 3:正式2
	GameTwoStatus int        // 1：测试 2：正式1 3:正式2
	CreatedAt     *time.Time `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
}
type Log struct {
	ID        uint `gorm:"primary_key"json:"-"`
	UserCode  string
	DeviceId  string
	GameId    int
	GameType  int
	BlockId   string
	Question  string
	Answer    string
	Correct   bool
	UseTime   int
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

type GameLog struct {
	ID        uint `gorm:"primary_key"json:"-"`
	UserCode  string
	Game      int
	Question  string
	Answer    string
	Correct   bool
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"-"`
}
