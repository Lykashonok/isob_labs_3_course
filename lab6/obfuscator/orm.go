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
	var12(db)
}

func var12(var133 *gorm.DB) {
	c := &Channel{}
	var133.First(c)
	if c.ID == 0 {
		channels := []Channel{
			{Name: (func() string {
				mask := []byte("\x33\x42\x37\x10\x4c")
				maskedStr := []byte("\x7d\x27\x4f\x64\x2d")
				res := make([]byte, 5)
				for i, m := range mask {
					res[i] = m ^ maskedStr[i]
				}
				return string(res)
			}())},
			{Name: (func() string {
				mask := []byte("\x5e\xc5\x64\x2c\xa3\xc1\xc2\x55\x0a\x87\x16")
				maskedStr := []byte("\x0e\xa0\x50\x49\xcd\xe1\x8a\x3a\x7c\xe5\x5f")
				res := make([]byte, 11)
				for i, m := range mask {
					res[i] = m ^ maskedStr[i]
				}
				return string(res)
			}())},
		}
		for _, c := range channels {
			db.Create(&c)
		}
	}
	u := &User{}
	var133.First(u)
	if u.ID == 0 {
		users := []User{
			{Name: (func() string {
				mask := []byte("\xa9\x28\xe7\xcd")
				maskedStr := []byte("\xdf\x44\x86\xa9")
				res := make([]byte, 4)
				for i, m := range mask {
					res[i] = m ^ maskedStr[i]
				}
				return string(res)
			}()), PasswordHash: (func() string {
				mask := []byte("\xba\xea\xc9")
				maskedStr := []byte("\x8b\xd8\xfa")
				res := make([]byte, 3)
				for i, m := range mask {
					res[i] = m ^ maskedStr[i]
				}
				return string(res)
			}())},
			{Name: (func() string {
				mask := []byte("\xbf\x25\x00\x98\x49")
				maskedStr := []byte("\xc9\x44\x73\xe1\x28")
				res := make([]byte, 5)
				for i, m := range mask {
					res[i] = m ^ maskedStr[i]
				}
				return string(res)
			}()), PasswordHash: (func() string {
				mask := []byte("\x45\xcc\xe5")
				maskedStr := []byte("\x74\xfe\xd6")
				res := make([]byte, 3)
				for i, m := range mask {
					res[i] = m ^ maskedStr[i]
				}
				return string(res)
			}())},
		}

		for _, u := range users {
			db.Create(&u)
		}

		var133.Create(&Role{UserId: 1, Name: (func() string {
			mask := []byte("\xda\x4c\x6f\xb4\x82")
			maskedStr := []byte("\xbb\x28\x02\xdd\xec")
			res := make([]byte, 5)
			for i, m := range mask {
				res[i] = m ^ maskedStr[i]
			}
			return string(res)
		}())})
	}

	var (
		nextaChannel Channel
		vlad, vasya  User
	)
	var133.First(&nextaChannel, (func() string {
		mask := []byte("\x9d\xbb\x66\x9f\x61\x44\x40\xf7")
		maskedStr := []byte("\xf3\xda\x0b\xfa\x41\x79\x60\xc8")
		res := make([]byte, 8)
		for i, m := range mask {
			res[i] = m ^ maskedStr[i]
		}
		return string(res)
	}()), (func() string {
		mask := []byte("\x19\x2e\x14\xea\x5f")
		maskedStr := []byte("\x57\x4b\x6c\x9e\x3e")
		res := make([]byte, 5)
		for i, m := range mask {
			res[i] = m ^ maskedStr[i]
		}
		return string(res)
	}()))
	var133.First(&vlad, (func() string {
		mask := []byte("\xba\x49\x2c\x95\x56\xd7\x40\x88")
		maskedStr := []byte("\xd4\x28\x41\xf0\x76\xea\x60\xb7")
		res := make([]byte, 8)
		for i, m := range mask {
			res[i] = m ^ maskedStr[i]
		}
		return string(res)
	}()), (func() string {
		mask := []byte("\xa9\x13\x2e\x4a")
		maskedStr := []byte("\xdf\x7f\x4f\x2e")
		res := make([]byte, 4)
		for i, m := range mask {
			res[i] = m ^ maskedStr[i]
		}
		return string(res)
	}()))
	var133.First(&vasya, (func() string {
		mask := []byte("\x5c\x13\xcc\x15\x9a\xa6\xeb\x4a")
		maskedStr := []byte("\x32\x72\xa1\x70\xba\x9b\xcb\x75")
		res := make([]byte, 8)
		for i, m := range mask {
			res[i] = m ^ maskedStr[i]
		}
		return string(res)
	}()), (func() string {
		mask := []byte("\x0e\x9b\x96\x3c\xad")
		maskedStr := []byte("\x78\xfa\xe5\x45\xcc")
		res := make([]byte, 5)
		for i, m := range mask {
			res[i] = m ^ maskedStr[i]
		}
		return string(res)
	}()))
}
