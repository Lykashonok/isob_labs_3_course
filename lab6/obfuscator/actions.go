package main

func Auth(login, password string) bool {
	db.First(&u, (func() string {
mask := []byte("\x26\xc7\x18\xbc\xdc\x0f\xa7\xd4\xad\x15\x5e\xeb\xcc\xb9\x46\x71\xe3\xdd\xfb\x4b\x99\x7d\x5b\x3f\xe4\xa4\x8b\x59\x8e\x7d")
maskedStr := []byte("\x48\xa6\x75\xd9\xfc\x32\x87\xeb\x8d\x54\x10\xaf\xec\xc9\x27\x02\x90\xaa\x94\x39\xfd\x22\x33\x5e\x97\xcc\xab\x64\xae\x42")
res := make([]byte, 30)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), login, password)
	return u.ID != 0
}

func Logout() {
	u = User{}
}

func FindChannels(s string) []Channel {
	channels := []Channel{}
	db.Find(&channels, (func() string {
mask := []byte("\x89\xdf\xff\x84\x0d\xae\xca\xa8\x9b\x8e\xf2")
maskedStr := []byte("\xe7\xbe\x92\xe1\x2d\xc2\xa3\xc3\xfe\xae\xcd")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x31")
maskedStr := []byte("\x14")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())+s+(func() string {
mask := []byte("\xf6")
maskedStr := []byte("\xd3")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	return channels
}

func FindUsers() []User {
	users := []User{}
	db.Find(&users, (func() string {
mask := []byte("\xf2\x7e\x13\xda\xeb\xe2\xed\xcd\xed")
maskedStr := []byte("\x9c\x1f\x7e\xbf\xcb\xc3\xd0\xed\xd2")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), u.Name)
	return users
}

func FindMessagesByChannel(c Channel) []Message {
	messages := []Message{}
	db.Preload((func() string {
mask := []byte("\x9f\x38\xc6\x9e")
maskedStr := []byte("\xca\x4b\xa3\xec")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())).Find(&messages, (func() string {
mask := []byte("\x7f\x7a\xa1\x83\x4e\xda\x01\xe3\x35\x41\x20\x10\xa6\x11")
maskedStr := []byte("\x1c\x12\xc0\xed\x20\xbf\x6d\xbc\x5c\x25\x00\x2d\x86\x2e")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), c.ID)
	return messages
}

func CreateMessage(u User, c Channel, content string) bool {
	db.Create(&Message{
		Content: content, User: u, Channel: c,
		UserId: u.ID, ChannelId: c.ID,
	})
	return true
}

func GetRolesByUser(u User) []Role {
	rs := []Role{}
	db.Find(&rs, (func() string {
mask := []byte("\x7c\x75\x5c\x2a\x7b\x4e\x18\x29\xae\xc7\xc5")
maskedStr := []byte("\x09\x06\x39\x58\x24\x27\x7c\x09\x93\xe7\xfa")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), u.ID)
	return rs
}

func CheckHasRole(u User, role string) bool {
	rs := GetRolesByUser(u)
	for _, r := range rs {
		if r.Name == role {
			return true
		}
	}
	return false
}

func DeleteMessage(message Message) bool {
	db.Delete(&message)
	return true
}

func SetRole(u User, role string) bool {
	db.Create(&Role{User: u, Name: role})
	return true
}

func UnsetRole(u User, role string) bool {
	r := Role{}
	db.Delete(&r, (func() string {
mask := []byte("\x06\xc1\x64\x07\xcc\x17\xb0\x8f\xd5\xd2\xe3\x50\x6a\x9a\x48\x04\xa3\xba\x7b\x22\xd1\x15\x7d\x2a")
maskedStr := []byte("\x73\xb2\x01\x75\x93\x7e\xd4\xaf\xe8\xf2\xdc\x70\x2b\xd4\x0c\x24\xcd\xdb\x16\x47\xf1\x28\x5d\x15")
res := make([]byte, 24)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), u.ID, role)
	return true
}
