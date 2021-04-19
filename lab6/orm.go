package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"varchar(50)"`
	PasswordHash string `gorm:"varchar(64)"`
}

type Channel struct {
	gorm.Model
	Name string `gorm:"varchar(50)"`
}

type Role struct {
	gorm.Model
	UserId uint
	User   User
	Name   string `gorm:"varchar(50)"`
}

type Message struct {
	gorm.Model
	Content   string `gorm:"varchar(255)"`
	UserId    uint
	User      User
	ChannelId uint
	Channel   Channel
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Channel{}, &Message{}, &Role{})
	seed(db)
}

func seed(db *gorm.DB) {
	c := &Channel{}
	db.First(c)
	if c.ID == 0 {
		channels := []Channel{
			{Name: "Nexta"},
			{Name: "Pe4en HovbI"},
		}
		for _, c := range channels {
			db.Create(&c)
		}
	}
	u := &User{}
	db.First(u)
	if u.ID == 0 {
		users := []User{
			{Name: "vlad", PasswordHash: "123"},
			{Name: "vasya", PasswordHash: "123"},
		}

		for _, u := range users {
			db.Create(&u)
		}

		db.Create(&Role{UserId: 1, Name: "admin"})
	}

	var (
		nextaChannel Channel
		vlad, vasya  User
	)
	db.First(&nextaChannel, "name = ?", "Nexta")
	db.First(&vlad, "name = ?", "vlad")
	db.First(&vasya, "name = ?", "vasya")
}
